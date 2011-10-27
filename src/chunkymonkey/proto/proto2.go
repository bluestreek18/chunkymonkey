package proto

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"log"
	"os"
	"reflect"

	. "chunkymonkey/types"
)

type IPacket interface {
	// IsPacket doesn't do anything, it's present purely for type-checking
	// packets.
	IsPacket()
}

// Packet definitions.

type PacketKeepAlive struct {
	Id int32
}

func (*PacketKeepAlive) IsPacket() {}

type PacketLogin struct {
	VersionOrEntityId int32
	Username          string
	MapSeed           RandomSeed
	GameMode          int32
	Dimension         DimensionId
	Difficulty        GameDifficulty
	WorldHeight       byte
	MaxPlayers        byte
}

func (*PacketLogin) IsPacket() {}

type PacketHandshake struct {
	UsernameOrHash string
}

func (*PacketHandshake) IsPacket() {}

type PacketChatMessage struct {
	Message string
}

func (*PacketChatMessage) IsPacket() {}

type PacketTimeUpdate struct {
	Time Ticks
}

func (*PacketTimeUpdate) IsPacket() {}

type PacketEntityEquipment struct {
	EntityId   EntityId
	Slot       SlotId
	ItemTypeId ItemTypeId
	Data       ItemData
}

func (*PacketEntityEquipment) IsPacket() {}

type PacketSpawnPosition struct {
	X BlockCoord
	Y int32
	Z BlockCoord
}

func (*PacketSpawnPosition) IsPacket() {}

type PacketUseEntity struct {
	User      EntityId
	Target    EntityId
	LeftClick bool
}

func (*PacketUseEntity) IsPacket() {}

type PacketUpdateHealth struct {
	Health         Health
	Food           FoodUnits
	FoodSaturation float32
}

func (*PacketUpdateHealth) IsPacket() {}

type PacketRespawn struct {
	Dimension   DimensionId
	Difficulty  GameDifficulty
	GameType    GameType
	WorldHeight int16
	MapSeed     RandomSeed
}

func (*PacketRespawn) IsPacket() {}

type PacketPlayer struct {
	OnGround bool
}

func (*PacketPlayer) IsPacket() {}

type PacketPlayerPosition struct {
	X, Y, Stance, Z AbsCoord
	OnGround        bool
}

func (*PacketPlayerPosition) IsPacket() {}

func (pkt *PacketPlayerPosition) Position() AbsXyz {
	return AbsXyz{pkt.X, pkt.Y, pkt.Z}
}

type PacketPlayerLook struct {
	Look     LookDegrees
	OnGround bool
}

func (*PacketPlayerLook) IsPacket() {}

type PacketPlayerPositionLook struct {
	X, Y1, Y2, Z AbsCoord
	Look         LookDegrees
	OnGround     bool
}

func (*PacketPlayerPositionLook) IsPacket() {}

func (pkt *PacketPlayerPositionLook) Stance(fromClient bool) AbsCoord {
	if fromClient {
		return pkt.Y2
	}
	return pkt.Y1
}

func (pkt *PacketPlayerPositionLook) Position(fromClient bool) AbsXyz {
	if fromClient {
		return AbsXyz{pkt.X, pkt.Y1, pkt.Z}
	}
	return AbsXyz{pkt.X, pkt.Y2, pkt.Z}
}

type PacketPlayerBlockHit struct {
	Status DigStatus
	Block  BlockXyz
	Face   Face
}

func (*PacketPlayerBlockHit) IsPacket() {}

type PacketPlayerBlockInteract struct {
	Block BlockXyz
	Face  Face
	Tool  ItemSlot
}

func (*PacketPlayerBlockInteract) IsPacket() {}

type PacketPlayerHoldingChange struct {
	SlotId SlotId
}

func (*PacketPlayerHoldingChange) IsPacket() {}

type PacketPlayerUseBed struct {
	EntityId EntityId
	Flag     byte
	Block    BlockXyz
}

func (*PacketPlayerUseBed) IsPacket() {}

type PacketEntityAnimation struct {
	EntityId  EntityId
	Animation EntityAnimation
}

func (*PacketEntityAnimation) IsPacket() {}

type PacketEntityAction struct {
	EntityId EntityId
	Action   EntityAction
}

func (*PacketEntityAction) IsPacket() {}

type PacketNamedEntitySpawn struct {
	EntityId    EntityId
	Username    string
	Position    AbsIntXyz
	Rotation    LookBytes
	CurrentItem ItemTypeId
}

func (*PacketNamedEntitySpawn) IsPacket() {}

type PacketItemSpawn struct {
	EntityId    EntityId
	ItemTypeId  ItemTypeId
	Count       ItemCount
	Data        ItemData
	Position    AbsIntXyz
	Orientation OrientationBytes
}

func (*PacketItemSpawn) IsPacket() {}

type PacketItemCollect struct {
	CollectedItem EntityId
	Collector     EntityId
}

func (*PacketItemCollect) IsPacket() {}

type PacketObjectSpawn struct {
	EntityId EntityId
	ObjType  ObjTypeId
	Position AbsIntXyz
}

func (*PacketObjectSpawn) IsPacket() {}

type PacketMobSpawn struct {
	EntityId EntityId
	MobType  EntityMobType
	Position AbsIntXyz
	Look     LookBytes
}

func (*PacketMobSpawn) IsPacket() {}

type PacketPaintingSpawn struct {
	EntityId EntityId
	Title    string
	Position AbsIntXyz
	SideFace SideFace
}

func (*PacketPaintingSpawn) IsPacket() {}

type PacketExperienceOrb struct {
	EntityId EntityId
	Position AbsIntXyz
	Count    int16
}

func (*PacketExperienceOrb) IsPacket() {}

type PacketEntityVelocity struct {
	EntityId EntityId
	Velocity Velocity
}

func (*PacketEntityVelocity) IsPacket() {}

type PacketEntityDestroy struct {
	EntityId EntityId
}

func (*PacketEntityDestroy) IsPacket() {}

type PacketEntity struct {
	EntityId EntityId
}

func (*PacketEntity) IsPacket() {}

type PacketEntityRelMove struct {
	EntityId EntityId
	Move     RelMove
}

func (*PacketEntityRelMove) IsPacket() {}

type PacketEntityLook struct {
	EntityId EntityId
	Look     LookBytes
}

func (*PacketEntityLook) IsPacket() {}

type PacketEntityLookAndRelMove struct {
	EntityId EntityId
	Move     RelMove
	Look     LookBytes
}

func (*PacketEntityLookAndRelMove) IsPacket() {}

type PacketEntityTeleport struct {
	EntityId EntityId
	Position AbsIntXyz
	Look     LookBytes
}

func (*PacketEntityTeleport) IsPacket() {}

type PacketEntityStatus struct {
	EntityId EntityId
	Status   EntityStatus
}

func (*PacketEntityStatus) IsPacket() {}

type PacketEntityAttach struct {
	EntityId  EntityId
	VehicleId EntityId
}

func (*PacketEntityAttach) IsPacket() {}

type PacketEntityMetadata struct {
	EntityId EntityId
	Metadata EntityMetadataTable
}

func (*PacketEntityMetadata) IsPacket() {}

type PacketEntityEffect struct {
	EntityId EntityId
	Effect   EntityEffect
	Value    int8
	Duration int16
}

func (*PacketEntityEffect) IsPacket() {}

type PacketEntityRemoveEffect struct {
	EntityId EntityId
	Effect   EntityEffect
}

func (*PacketEntityRemoveEffect) IsPacket() {}

type PacketPlayerExperience struct {
	Experience      int8
	Level           int8
	TotalExperience int16
}

func (*PacketPlayerExperience) IsPacket() {}

type PacketPreChunk struct {
	ChunkLoc ChunkXz
	Mode     ChunkLoadMode
}

func (*PacketPreChunk) IsPacket() {}

type PacketMapChunk struct {
	Corner BlockXyz
	Data   ChunkData
}

func (*PacketMapChunk) IsPacket() {}

type PacketMultiBlockChange struct {
	ChunkLoc ChunkXz
	Changes  MultiBlockChanges
}

func (*PacketMultiBlockChange) IsPacket() {}

type PacketBlockChange struct {
	Position  BlockXyz
	TypeId    byte
	BlockData byte
}

func (*PacketBlockChange) IsPacket() {}

type PacketBlockAction struct {
	// TODO Hopefully other packets referencing block locations (BlockXyz) will
	// become consistent and use the same type as this for Y.
	X              int32
	Y              int16
	Z              int32
	Value1, Value2 byte
}

func (*PacketBlockAction) IsPacket() {}

type PacketExplosion struct {
	Center AbsXyz
	Radius float32
	Blocks BlocksDxyz
}

func (*PacketExplosion) IsPacket() {}

type PacketSoundEffect struct {
	Effect SoundEffect
	Block  BlockXyz
	Data   int32
}

func (*PacketSoundEffect) IsPacket() {}

type PacketState struct {
	Reason   byte
	GameType GameType
}

func (*PacketState) IsPacket() {}

type PacketThunderbolt struct {
	EntityId EntityId
	Flag     bool
	Position AbsIntXyz
}

func (*PacketThunderbolt) IsPacket() {}

type PacketWindowOpen struct {
	WindowId  WindowId
	Inventory InvTypeId
	Title     string
	NumSlots  byte
}

func (*PacketWindowOpen) IsPacket() {}

type PacketWindowClose struct {
	WindowId WindowId
}

func (*PacketWindowClose) IsPacket() {}

type PacketWindowClick struct {
	WindowId     WindowId
	Slot         SlotId
	RightClick   bool
	TxId         TxId
	Shift        bool
	ExpectedSlot ItemSlot
}

func (*PacketWindowClick) IsPacket() {}

type PacketWindowSetSlot struct {
	WindowId WindowId
	Slot     SlotId
	NewSlot  ItemSlot
}

func (*PacketWindowSetSlot) IsPacket() {}

type PacketWindowItems struct {
	WindowId WindowId
	Slots    ItemSlotSlice
}

func (*PacketWindowItems) IsPacket() {}

type PacketWindowProgressBar struct {
	WindowId WindowId
	PrgBarId PrgBarId
	Value    PrgBarValue
}

func (*PacketWindowProgressBar) IsPacket() {}

type PacketWindowTransaction struct {
	WindowId WindowId
	TxId     TxId
	Accepted bool
}

func (*PacketWindowTransaction) IsPacket() {}

type PacketCreativeInventoryAction struct {
	Slot       SlotId
	ItemTypeId ItemTypeId
	// Note that unlike other packets, the Count and Data are always present.
	Count int16
	Data  ItemData
}

func (*PacketCreativeInventoryAction) IsPacket() {}

type PacketSignUpdate struct {
	X     int32
	Y     int16
	Z     int32
	Text1 string
	Text2 string
	Text3 string
	Text4 string
}

func (*PacketSignUpdate) IsPacket() {}

type PacketItemData struct {
	ItemTypeId ItemTypeId
	MapId      ItemData
	MapData    MapData
}

func (*PacketItemData) IsPacket() {}

type PacketIncrementStatistic struct {
	StatisticId StatisticId
	Amount      byte
}

func (*PacketIncrementStatistic) IsPacket() {}

type PacketPlayerListItem struct {
	Username string
	Online   bool
	Ping     int16
}

func (*PacketPlayerListItem) IsPacket() {}

type PacketServerListPing struct{}

func (*PacketServerListPing) IsPacket() {}

type PacketDisconnect struct {
	Reason string
}

func (*PacketDisconnect) IsPacket() {}

// Special packet field types.

// EntityMetadataTable implements IMarshaler.
type EntityMetadataTable struct {
	Items []EntityMetadata
}

func (emt *EntityMetadataTable) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	emt.Items, err = readEntityMetadataField(reader)
	return
}

func (emt *EntityMetadataTable) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	return writeEntityMetadataField(writer, emt.Items)
}

// ItemSlot implements IMarshaler.
type ItemSlot struct {
	ItemTypeId ItemTypeId
	Count      ItemCount
	Data       ItemData
}

func (is *ItemSlot) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	typeIdUint16, err := ps.readUint16(reader)
	if err != nil {
		return
	}
	is.ItemTypeId = ItemTypeId(typeIdUint16)

	if is.ItemTypeId == -1 {
		is.Count = 0
		is.Data = 0
	} else {
		countUint8, err := ps.readUint8(reader)
		if err != nil {
			return err
		}
		dataUint16, err := ps.readUint16(reader)
		if err != nil {
			return err
		}

		is.Count = ItemCount(countUint8)
		is.Data = ItemData(dataUint16)
	}
	return
}

func (is *ItemSlot) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	if err = ps.writeUint16(writer, uint16(is.ItemTypeId)); err != nil {
		return
	}

	if is.ItemTypeId != -1 {
		if err = ps.writeUint8(writer, uint8(is.Count)); err != nil {
			return
		}
		if err = ps.writeUint16(writer, uint16(is.Data)); err != nil {
			return
		}
	}

	return
}

// ItemSlotSlice implements IMarshaler.
type ItemSlotSlice []ItemSlot

func (slots *ItemSlotSlice) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	var numSlots int16
	if err = binary.Read(reader, binary.BigEndian, &numSlots); err != nil {
		return
	} else if numSlots < 0 {
		return ErrorLengthNegative
	}

	*slots = make(ItemSlotSlice, numSlots)

	for i := range *slots {
		if err = (*slots)[i].MinecraftUnmarshal(reader, ps); err != nil {
			return
		}
	}

	return
}

func (slots *ItemSlotSlice) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	numSlots := int16(len(*slots))
	if err = binary.Write(writer, binary.BigEndian, numSlots); err != nil {
		return
	}

	for i := range *slots {
		if err = (*slots)[i].MinecraftMarshal(writer, ps); err != nil {
			return
		}
	}

	return
}

// ChunkData implements IMarshaler.
type ChunkData struct {
	Size ChunkDataSize
	Data []byte
}

// ChunkDataSize contains the dimensions of the data represented inside ChunkData.
type ChunkDataSize struct {
	X, Y, Z byte
}

func (cd *ChunkData) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	if err = ps.readData(reader, reflect.Indirect(reflect.ValueOf(&cd.Size))); err != nil {
		return
	}

	lengthUint32, err := ps.readUint32(reader)
	if err != nil {
		return
	}

	length := int32(lengthUint32)
	if length < 0 {
		return ErrorLengthNegative
	}

	zReader, err := zlib.NewReader(&io.LimitedReader{reader, int64(length)})
	if err != nil {
		return
	}
	defer zReader.Close()

	numBlocks := (int(cd.Size.X) + 1) * (int(cd.Size.Y) + 1) * (int(cd.Size.Z) + 1)
	expectedNumDataBytes := numBlocks + 3*(numBlocks>>1)
	cd.Data = make([]byte, expectedNumDataBytes)
	if _, err = io.ReadFull(zReader, cd.Data); err != nil {
		return
	}

	// Check that we're at the end of the compressed data to be sure of being in
	// sync with packet stream..
	n, err := io.ReadFull(zReader, dump[:])
	if err == os.EOF {
		err = nil
		if n > 0 {
			log.Printf("Unexpected extra chunk data byte of %d bytes", n)
		}
	} else if err == nil {
		log.Printf("Unexpected extra chunk data byte of at least %d bytes - assuming bad packet stream", n)
		return ErrorBadPacketData
	} else {
		// Other error.
		return err
	}

	return nil
}

func (cd *ChunkData) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	if err = ps.writeData(writer, reflect.Indirect(reflect.ValueOf(&cd.Size))); err != nil {
		return
	}

	numBlocks := (int(cd.Size.X) + 1) * (int(cd.Size.Y) + 1) * (int(cd.Size.Z) + 1)
	expectedNumDataBytes := numBlocks + 3*(numBlocks>>1)
	if len(cd.Data) != expectedNumDataBytes {
		return ErrorBadChunkDataSize
	}

	buf := bytes.NewBuffer(make([]byte, 0, 4096))

	zWriter, err := zlib.NewWriter(buf)
	if err != nil {
		return
	}
	_, err = zWriter.Write(cd.Data)
	zWriter.Close()
	if err != nil {
		return
	}

	compressedBytes := buf.Bytes()
	if err = ps.writeUint32(writer, uint32(len(compressedBytes))); err != nil {
		return
	}

	_, err = writer.Write(compressedBytes)
	return
}

// MultiBlockChanges implements IMarshaler.
type MultiBlockChanges struct {
	// Coords are packed x,y,z block coordinates relative to a chunk origin. Note
	// that these differ from the value for BlockIndex, which supplies conversion
	// methods for this purpose.
	Coords    []int16
	TypeIds   []byte
	BlockData []byte
}

func (mbc *MultiBlockChanges) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	numBlocksUint16, err := ps.readUint16(reader)
	if err != nil {
		return
	}

	numBlocks := int16(numBlocksUint16)
	if numBlocks < 0 {
		return ErrorLengthNegative
	} else if numBlocks == 0 {
		// Odd case.
		return nil
	}

	mbc.Coords = make([]int16, numBlocks)
	if err = binary.Read(reader, binary.BigEndian, mbc.Coords); err != nil {
		return
	}

	mbc.TypeIds = make([]byte, numBlocks)
	if _, err = io.ReadFull(reader, mbc.TypeIds); err != nil {
		return
	}

	mbc.BlockData = make([]byte, numBlocks)
	_, err = io.ReadFull(reader, mbc.BlockData)

	return
}

func (mbc *MultiBlockChanges) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	numBlocks := len(mbc.Coords)
	if numBlocks != len(mbc.TypeIds) || numBlocks != len(mbc.BlockData) {
		return ErrorMismatchingValues
	}

	if err = ps.writeUint16(writer, uint16(numBlocks)); err != nil {
		return
	}

	if err = binary.Write(writer, binary.BigEndian, mbc.Coords); err != nil {
		return
	}

	if _, err = writer.Write(mbc.TypeIds); err != nil {
		return
	}

	_, err = writer.Write(mbc.BlockData)
	return
}

// BlocksDxyz contains 3 * number of block relative locations. [0:3] contains
// the first, [3:6] the second, etc.
type BlocksDxyz []byte

func (b *BlocksDxyz) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	numBlocksUint32, err := ps.readUint32(reader)
	if err != nil {
		return
	}

	numBlocks := int32(numBlocksUint32)
	if numBlocks < 0 {
		return ErrorLengthNegative
	}

	*b = make([]byte, 3*numBlocks)
	_, err = io.ReadFull(reader, *b)

	return
}

func (b *BlocksDxyz) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	if err = ps.writeUint32(writer, uint32(len(*b))/3); err != nil {
		return
	}

	_, err = writer.Write(*b)

	return
}

// MapData implements IMarshaler.
type MapData []byte

func (md *MapData) MinecraftUnmarshal(reader io.Reader, ps *PacketSerializer) (err os.Error) {
	length, err := ps.readUint8(reader)
	if err != nil {
		return
	}

	*md = make(MapData, length)
	_, err = io.ReadFull(reader, []byte(*md))
	return
}

func (md *MapData) MinecraftMarshal(writer io.Writer, ps *PacketSerializer) (err os.Error) {
	if err = ps.writeUint8(writer, byte(len(*md))); err != nil {
		return
	}

	_, err = writer.Write([]byte(*md))
	return
}
