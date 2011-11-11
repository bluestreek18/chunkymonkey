package chunkymonkey

import (
	"fmt"
	"log"
	"os"
	"net"

	. "chunkymonkey/entity"
	"chunkymonkey/gamerules"
	"chunkymonkey/player"
	"chunkymonkey/proto"
	"chunkymonkey/server_auth"
	"chunkymonkey/shardserver"
	"chunkymonkey/worldstore"
	"nbt"
)

// TODO Refactor this more simply after a good re-working of the chunkymonkey/proto package.

const (
	connTypeUnknown = iota
	connTypeLogin
	connTypeServerQuery
)

var (
	clientErrGeneral      = os.NewError("Server error.")
	clientErrUsername     = os.NewError("Bad username.")
	clientErrLoginDenied  = os.NewError("You do not have access to this server.")
	clientErrHandshake    = os.NewError("Handshake error.")
	clientErrLoginGeneral = os.NewError("Login error.")
	clientErrAuthFailed   = os.NewError("Minecraft authentication failed.")
	clientErrUserData     = os.NewError("Error reading user data. Please contact the server administrator.")

	loginErrorConnType    = os.NewError("unknown/bad connection type")
	loginErrorMaintenance = os.NewError("server under maintenance")
	loginErrorServerList  = os.NewError("server list poll")
)

type GameInfo struct {
	game           *Game
	maxPlayerCount int
	serverDesc     string
	maintenanceMsg string
	serverId       string
	shardManager   *shardserver.LocalShardManager
	entityManager  *EntityManager
	worldStore     *worldstore.WorldStore
	authserver     server_auth.IAuthenticator
}

// Handles connections for a game on the given socket.
type ConnHandler struct {
	// UpdateGameInfo is used to reconfigure a running ConnHandler. A Game must
	// pass something in before this ConnHandler will accept connections.
	UpdateGameInfo chan *GameInfo

	listener net.Listener
	gameInfo *GameInfo
}

// NewConnHandler creates and starts a ConnHandler.
func NewConnHandler(listener net.Listener, gameInfo *GameInfo) *ConnHandler {
	ch := &ConnHandler{
		UpdateGameInfo: make(chan *GameInfo),
		listener:       listener,
		gameInfo:       gameInfo,
	}

	go ch.run()

	return ch
}

// Stop stops the connection handler from accepting any further connections.
func (ch *ConnHandler) Stop() {
	close(ch.UpdateGameInfo)
	ch.listener.Close()
}

func (ch *ConnHandler) run() {
	defer ch.listener.Close()
	var ok bool

	for {
		conn, err := ch.listener.Accept()
		if err != nil {
			log.Print("Accept: ", err)
			return
		}

		// Check for updated game info.
		select {
		case ch.gameInfo, ok = <-ch.UpdateGameInfo:
			if !ok {
				log.Print("Connection handler shut down.")
				return
			}
		default:
		}

		newLogin := &pktHandler{
			gameInfo: ch.gameInfo,
			conn:     conn,
		}
		go newLogin.handle()
	}
}

type pktHandler struct {
	gameInfo *GameInfo
	conn     net.Conn
	ps       proto.PacketSerializer
}

func (l *pktHandler) handle() {
	var err, clientErr os.Error

	defer func() {
		if err != nil {
			log.Print("Connection closed ", err.String())
			if clientErr == nil {
				clientErr = clientErrGeneral
			}
			l.ps.WritePacket(l.conn, &proto.PacketDisconnect{
				Reason: clientErr.String(),
			})
			l.conn.Close()
		}
	}()

	pkt, err := l.ps.ReadPacketExpect(l.conn, true, 0x02, 0xfe)
	if err != nil {
		clientErr = clientErrLoginGeneral
		return
	}

	switch p := pkt.(type) {
	case *proto.PacketHandshake:
		err, clientErr = l.handleLogin(p)
	case *proto.PacketServerListPing:
		err, clientErr = l.handleServerQuery()
	default:
		err = loginErrorConnType
	}
}

func (l *pktHandler) handleLogin(pktHandshake *proto.PacketHandshake) (err, clientErr os.Error) {
	username := pktHandshake.UsernameOrHash
	if !validPlayerUsername.MatchString(username) {
		err = clientErrUsername
		clientErr = err
		return
	}

	log.Print("Client ", l.conn.RemoteAddr(), " connected as ", username)

	// TODO Allow admins to connect.
	if l.gameInfo.maintenanceMsg != "" {
		err = loginErrorMaintenance
		clientErr = os.NewError(l.gameInfo.maintenanceMsg)
		return
	}

	// Load player permissions.
	permissions := gamerules.Permissions.UserPermissions(username)
	if !permissions.Has("login") {
		err = fmt.Errorf("Player %q does not have login permission", username)
		clientErr = clientErrLoginDenied
		return
	}

	if err = l.ps.WritePacket(l.conn, &proto.PacketHandshake{l.gameInfo.serverId}); err != nil {
		clientErr = clientErrHandshake
		return
	}

	if l.gameInfo.serverId != "-" {
		var authenticated bool
		authenticated, err = l.gameInfo.authserver.Authenticate(l.gameInfo.serverId, username)
		if !authenticated || err != nil {
			var reason string
			if err != nil {
				reason = "Authentication check failed: " + err.String()
			} else {
				reason = "Failed authentication"
			}
			err = fmt.Errorf("Client %v: %s", l.conn.RemoteAddr(), reason)
			clientErr = clientErrAuthFailed
			return
		}
		log.Print("Client ", l.conn.RemoteAddr(), " passed minecraft.net authentication")
	}

	if _, err = l.ps.ReadPacketExpect(l.conn, true, 0x01); err != nil {
		clientErr = clientErrLoginGeneral
		return
	}

	entityId := l.gameInfo.entityManager.NewEntity()

	var playerData *nbt.Compound
	if playerData, err = l.gameInfo.game.worldStore.PlayerData(username); err != nil {
		clientErr = clientErrUserData
		return
	}

	player := player.NewPlayer(entityId, l.gameInfo.shardManager, l.conn, username, l.gameInfo.worldStore.SpawnPosition, l.gameInfo.game.playerDisconnect, l.gameInfo.game)
	if playerData != nil {
		if err = player.UnmarshalNbt(playerData); err != nil {
			// Don't let the player log in, as they will only have default inventory
			// etc., which could lose items from them. Better for an administrator to
			// sort this out.
			err = fmt.Errorf("Error parsing player data for %q: %v", username, err)
			clientErr = clientErrUserData
			return
		}
	}

	l.gameInfo.game.playerConnect <- player
	player.Run()

	return
}

func (l *pktHandler) handleServerQuery() (err, clientErr os.Error) {
	err = loginErrorServerList
	clientErr = fmt.Errorf(
		"%s§%d§%d",
		l.gameInfo.serverDesc,
		l.gameInfo.game.PlayerCount(), l.gameInfo.maxPlayerCount)
	return
}