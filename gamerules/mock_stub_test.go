// Automatically generated by MockGen. DO NOT EDIT!
// Source: gamerules/stub.go

package gamerules

import (
	proto "github.com/huin/chunkymonkey/proto"
	gomock "code.google.com/p/gomock/gomock"
	. "github.com/huin/chunkymonkey/types"
)

// Mock of IShardConnecter interface
type MockIShardConnecter struct {
	ctrl     *gomock.Controller
	recorder *_MockIShardConnecterRecorder
}

// Recorder for MockIShardConnecter (not exported)
type _MockIShardConnecterRecorder struct {
	mock *MockIShardConnecter
}

func NewMockIShardConnecter(ctrl *gomock.Controller) *MockIShardConnecter {
	mock := &MockIShardConnecter{ctrl: ctrl}
	mock.recorder = &_MockIShardConnecterRecorder{mock}
	return mock
}

func (_m *MockIShardConnecter) EXPECT() *_MockIShardConnecterRecorder {
	return _m.recorder
}

func (_m *MockIShardConnecter) PlayerShardConnect(entityId EntityId, player IPlayerClient, shardLoc ShardXz) IPlayerShardClient {
	ret := _m.ctrl.Call(_m, "PlayerShardConnect", entityId, player, shardLoc)
	ret0, _ := ret[0].(IPlayerShardClient)
	return ret0
}

func (_mr *_MockIShardConnecterRecorder) PlayerShardConnect(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PlayerShardConnect", arg0, arg1, arg2)
}

func (_m *MockIShardConnecter) ShardShardConnect(shardLoc ShardXz) IShardShardClient {
	ret := _m.ctrl.Call(_m, "ShardShardConnect", shardLoc)
	ret0, _ := ret[0].(IShardShardClient)
	return ret0
}

func (_mr *_MockIShardConnecterRecorder) ShardShardConnect(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShardShardConnect", arg0)
}

// Mock of IPlayerShardClient interface
type MockIPlayerShardClient struct {
	ctrl     *gomock.Controller
	recorder *_MockIPlayerShardClientRecorder
}

// Recorder for MockIPlayerShardClient (not exported)
type _MockIPlayerShardClientRecorder struct {
	mock *MockIPlayerShardClient
}

func NewMockIPlayerShardClient(ctrl *gomock.Controller) *MockIPlayerShardClient {
	mock := &MockIPlayerShardClient{ctrl: ctrl}
	mock.recorder = &_MockIPlayerShardClientRecorder{mock}
	return mock
}

func (_m *MockIPlayerShardClient) EXPECT() *_MockIPlayerShardClientRecorder {
	return _m.recorder
}

func (_m *MockIPlayerShardClient) Disconnect() {
	_m.ctrl.Call(_m, "Disconnect")
}

func (_mr *_MockIPlayerShardClientRecorder) Disconnect() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disconnect")
}

func (_m *MockIPlayerShardClient) ReqSubscribeChunk(chunkLoc ChunkXz, notify bool) {
	_m.ctrl.Call(_m, "ReqSubscribeChunk", chunkLoc, notify)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqSubscribeChunk(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqSubscribeChunk", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqUnsubscribeChunk(chunkLoc ChunkXz) {
	_m.ctrl.Call(_m, "ReqUnsubscribeChunk", chunkLoc)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqUnsubscribeChunk(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqUnsubscribeChunk", arg0)
}

func (_m *MockIPlayerShardClient) ReqMulticastPlayers(chunkLoc ChunkXz, exclude EntityId, packet []byte) {
	_m.ctrl.Call(_m, "ReqMulticastPlayers", chunkLoc, exclude, packet)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqMulticastPlayers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqMulticastPlayers", arg0, arg1, arg2)
}

func (_m *MockIPlayerShardClient) ReqAddPlayerData(chunkLoc ChunkXz, name string, position AbsXyz, look LookBytes, held ItemTypeId) {
	_m.ctrl.Call(_m, "ReqAddPlayerData", chunkLoc, name, position, look, held)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqAddPlayerData(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqAddPlayerData", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockIPlayerShardClient) ReqRemovePlayerData(chunkLoc ChunkXz, isDisconnect bool) {
	_m.ctrl.Call(_m, "ReqRemovePlayerData", chunkLoc, isDisconnect)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqRemovePlayerData(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqRemovePlayerData", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqSetPlayerPosition(chunkLoc ChunkXz, position AbsXyz) {
	_m.ctrl.Call(_m, "ReqSetPlayerPosition", chunkLoc, position)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqSetPlayerPosition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqSetPlayerPosition", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqSetPlayerLook(chunkLoc ChunkXz, look LookBytes) {
	_m.ctrl.Call(_m, "ReqSetPlayerLook", chunkLoc, look)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqSetPlayerLook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqSetPlayerLook", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqHitBlock(held Slot, target BlockXyz, digStatus DigStatus, face Face) {
	_m.ctrl.Call(_m, "ReqHitBlock", held, target, digStatus, face)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqHitBlock(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqHitBlock", arg0, arg1, arg2, arg3)
}

func (_m *MockIPlayerShardClient) ReqInteractBlock(held Slot, target BlockXyz, face Face) {
	_m.ctrl.Call(_m, "ReqInteractBlock", held, target, face)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqInteractBlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqInteractBlock", arg0, arg1, arg2)
}

func (_m *MockIPlayerShardClient) ReqPlaceItem(target BlockXyz, slot Slot) {
	_m.ctrl.Call(_m, "ReqPlaceItem", target, slot)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqPlaceItem(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqPlaceItem", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqTakeItem(chunkLoc ChunkXz, entityId EntityId) {
	_m.ctrl.Call(_m, "ReqTakeItem", chunkLoc, entityId)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqTakeItem(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqTakeItem", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqDropItem(content Slot, position AbsXyz, velocity AbsVelocity, pickupImmunity Ticks) {
	_m.ctrl.Call(_m, "ReqDropItem", content, position, velocity, pickupImmunity)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqDropItem(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqDropItem", arg0, arg1, arg2, arg3)
}

func (_m *MockIPlayerShardClient) ReqInventoryClick(block BlockXyz, click Click) {
	_m.ctrl.Call(_m, "ReqInventoryClick", block, click)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqInventoryClick(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqInventoryClick", arg0, arg1)
}

func (_m *MockIPlayerShardClient) ReqInventoryUnsubscribed(block BlockXyz) {
	_m.ctrl.Call(_m, "ReqInventoryUnsubscribed", block)
}

func (_mr *_MockIPlayerShardClientRecorder) ReqInventoryUnsubscribed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqInventoryUnsubscribed", arg0)
}

// Mock of IShardShardClient interface
type MockIShardShardClient struct {
	ctrl     *gomock.Controller
	recorder *_MockIShardShardClientRecorder
}

// Recorder for MockIShardShardClient (not exported)
type _MockIShardShardClientRecorder struct {
	mock *MockIShardShardClient
}

func NewMockIShardShardClient(ctrl *gomock.Controller) *MockIShardShardClient {
	mock := &MockIShardShardClient{ctrl: ctrl}
	mock.recorder = &_MockIShardShardClientRecorder{mock}
	return mock
}

func (_m *MockIShardShardClient) EXPECT() *_MockIShardShardClientRecorder {
	return _m.recorder
}

func (_m *MockIShardShardClient) Disconnect() {
	_m.ctrl.Call(_m, "Disconnect")
}

func (_mr *_MockIShardShardClientRecorder) Disconnect() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Disconnect")
}

func (_m *MockIShardShardClient) ReqSetActiveBlocks(blocks []BlockXyz) {
	_m.ctrl.Call(_m, "ReqSetActiveBlocks", blocks)
}

func (_mr *_MockIShardShardClientRecorder) ReqSetActiveBlocks(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqSetActiveBlocks", arg0)
}

func (_m *MockIShardShardClient) ReqTransferEntity(loc ChunkXz, entity INonPlayerEntity) {
	_m.ctrl.Call(_m, "ReqTransferEntity", loc, entity)
}

func (_mr *_MockIShardShardClientRecorder) ReqTransferEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReqTransferEntity", arg0, arg1)
}

// Mock of IGame interface
type MockIGame struct {
	ctrl     *gomock.Controller
	recorder *_MockIGameRecorder
}

// Recorder for MockIGame (not exported)
type _MockIGameRecorder struct {
	mock *MockIGame
}

func NewMockIGame(ctrl *gomock.Controller) *MockIGame {
	mock := &MockIGame{ctrl: ctrl}
	mock.recorder = &_MockIGameRecorder{mock}
	return mock
}

func (_m *MockIGame) EXPECT() *_MockIGameRecorder {
	return _m.recorder
}

func (_m *MockIGame) BroadcastPacket(packet []byte) {
	_m.ctrl.Call(_m, "BroadcastPacket", packet)
}

func (_mr *_MockIGameRecorder) BroadcastPacket(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BroadcastPacket", arg0)
}

func (_m *MockIGame) BroadcastMessage(msg string) {
	_m.ctrl.Call(_m, "BroadcastMessage", msg)
}

func (_mr *_MockIGameRecorder) BroadcastMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BroadcastMessage", arg0)
}

func (_m *MockIGame) PlayerByName(name string) IPlayerClient {
	ret := _m.ctrl.Call(_m, "PlayerByName", name)
	ret0, _ := ret[0].(IPlayerClient)
	return ret0
}

func (_mr *_MockIGameRecorder) PlayerByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PlayerByName", arg0)
}

func (_m *MockIGame) PlayerByEntityId(id EntityId) IPlayerClient {
	ret := _m.ctrl.Call(_m, "PlayerByEntityId", id)
	ret0, _ := ret[0].(IPlayerClient)
	return ret0
}

func (_mr *_MockIGameRecorder) PlayerByEntityId(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PlayerByEntityId", arg0)
}

func (_m *MockIGame) ItemTypeById(id int) (ItemType, bool) {
	ret := _m.ctrl.Call(_m, "ItemTypeById", id)
	ret0, _ := ret[0].(ItemType)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockIGameRecorder) ItemTypeById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ItemTypeById", arg0)
}

// Mock of IPlayerClient interface
type MockIPlayerClient struct {
	ctrl     *gomock.Controller
	recorder *_MockIPlayerClientRecorder
}

// Recorder for MockIPlayerClient (not exported)
type _MockIPlayerClientRecorder struct {
	mock *MockIPlayerClient
}

func NewMockIPlayerClient(ctrl *gomock.Controller) *MockIPlayerClient {
	mock := &MockIPlayerClient{ctrl: ctrl}
	mock.recorder = &_MockIPlayerClientRecorder{mock}
	return mock
}

func (_m *MockIPlayerClient) EXPECT() *_MockIPlayerClientRecorder {
	return _m.recorder
}

func (_m *MockIPlayerClient) GetEntityId() EntityId {
	ret := _m.ctrl.Call(_m, "GetEntityId")
	ret0, _ := ret[0].(EntityId)
	return ret0
}

func (_mr *_MockIPlayerClientRecorder) GetEntityId() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEntityId")
}

func (_m *MockIPlayerClient) TransmitPacket(packet []byte) {
	_m.ctrl.Call(_m, "TransmitPacket", packet)
}

func (_mr *_MockIPlayerClientRecorder) TransmitPacket(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TransmitPacket", arg0)
}

func (_m *MockIPlayerClient) NotifyChunkLoad() {
	_m.ctrl.Call(_m, "NotifyChunkLoad")
}

func (_mr *_MockIPlayerClientRecorder) NotifyChunkLoad() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NotifyChunkLoad")
}

func (_m *MockIPlayerClient) InventorySubscribed(block BlockXyz, invTypeId InvTypeId, slots []proto.WindowSlot) {
	_m.ctrl.Call(_m, "InventorySubscribed", block, invTypeId, slots)
}

func (_mr *_MockIPlayerClientRecorder) InventorySubscribed(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventorySubscribed", arg0, arg1, arg2)
}

func (_m *MockIPlayerClient) InventorySlotUpdate(block BlockXyz, slot Slot, slotId SlotId) {
	_m.ctrl.Call(_m, "InventorySlotUpdate", block, slot, slotId)
}

func (_mr *_MockIPlayerClientRecorder) InventorySlotUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventorySlotUpdate", arg0, arg1, arg2)
}

func (_m *MockIPlayerClient) InventoryProgressUpdate(block BlockXyz, prgBarId PrgBarId, value PrgBarValue) {
	_m.ctrl.Call(_m, "InventoryProgressUpdate", block, prgBarId, value)
}

func (_mr *_MockIPlayerClientRecorder) InventoryProgressUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventoryProgressUpdate", arg0, arg1, arg2)
}

func (_m *MockIPlayerClient) InventoryCursorUpdate(block BlockXyz, cursor Slot) {
	_m.ctrl.Call(_m, "InventoryCursorUpdate", block, cursor)
}

func (_mr *_MockIPlayerClientRecorder) InventoryCursorUpdate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventoryCursorUpdate", arg0, arg1)
}

func (_m *MockIPlayerClient) InventoryTxState(block BlockXyz, txId TxId, accepted bool) {
	_m.ctrl.Call(_m, "InventoryTxState", block, txId, accepted)
}

func (_mr *_MockIPlayerClientRecorder) InventoryTxState(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventoryTxState", arg0, arg1, arg2)
}

func (_m *MockIPlayerClient) InventoryUnsubscribed(block BlockXyz) {
	_m.ctrl.Call(_m, "InventoryUnsubscribed", block)
}

func (_mr *_MockIPlayerClientRecorder) InventoryUnsubscribed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InventoryUnsubscribed", arg0)
}

func (_m *MockIPlayerClient) PlaceHeldItem(target BlockXyz, wasHeld Slot) {
	_m.ctrl.Call(_m, "PlaceHeldItem", target, wasHeld)
}

func (_mr *_MockIPlayerClientRecorder) PlaceHeldItem(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PlaceHeldItem", arg0, arg1)
}

func (_m *MockIPlayerClient) OfferItem(fromChunk ChunkXz, entityId EntityId, item Slot) {
	_m.ctrl.Call(_m, "OfferItem", fromChunk, entityId, item)
}

func (_mr *_MockIPlayerClientRecorder) OfferItem(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OfferItem", arg0, arg1, arg2)
}

func (_m *MockIPlayerClient) GiveItemAtPosition(atPosition AbsXyz, item Slot) {
	_m.ctrl.Call(_m, "GiveItemAtPosition", atPosition, item)
}

func (_mr *_MockIPlayerClientRecorder) GiveItemAtPosition(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GiveItemAtPosition", arg0, arg1)
}

func (_m *MockIPlayerClient) GiveItem(item Slot) {
	_m.ctrl.Call(_m, "GiveItem", item)
}

func (_mr *_MockIPlayerClientRecorder) GiveItem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GiveItem", arg0)
}

func (_m *MockIPlayerClient) PositionLook() (AbsXyz, LookDegrees) {
	ret := _m.ctrl.Call(_m, "PositionLook")
	ret0, _ := ret[0].(AbsXyz)
	ret1, _ := ret[1].(LookDegrees)
	return ret0, ret1
}

func (_mr *_MockIPlayerClientRecorder) PositionLook() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PositionLook")
}

func (_m *MockIPlayerClient) SetPositionLook(_param0 AbsXyz, _param1 LookDegrees) {
	_m.ctrl.Call(_m, "SetPositionLook", _param0, _param1)
}

func (_mr *_MockIPlayerClientRecorder) SetPositionLook(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPositionLook", arg0, arg1)
}

func (_m *MockIPlayerClient) EchoMessage(msg string) {
	_m.ctrl.Call(_m, "EchoMessage", msg)
}

func (_mr *_MockIPlayerClientRecorder) EchoMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EchoMessage", arg0)
}

// Mock of ICommandFramework interface
type MockICommandFramework struct {
	ctrl     *gomock.Controller
	recorder *_MockICommandFrameworkRecorder
}

// Recorder for MockICommandFramework (not exported)
type _MockICommandFrameworkRecorder struct {
	mock *MockICommandFramework
}

func NewMockICommandFramework(ctrl *gomock.Controller) *MockICommandFramework {
	mock := &MockICommandFramework{ctrl: ctrl}
	mock.recorder = &_MockICommandFrameworkRecorder{mock}
	return mock
}

func (_m *MockICommandFramework) EXPECT() *_MockICommandFrameworkRecorder {
	return _m.recorder
}

func (_m *MockICommandFramework) Prefix() string {
	ret := _m.ctrl.Call(_m, "Prefix")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockICommandFrameworkRecorder) Prefix() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prefix")
}

func (_m *MockICommandFramework) Process(player IPlayerClient, cmd string, game IGame) {
	_m.ctrl.Call(_m, "Process", player, cmd, game)
}

func (_mr *_MockICommandFrameworkRecorder) Process(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Process", arg0, arg1, arg2)
}