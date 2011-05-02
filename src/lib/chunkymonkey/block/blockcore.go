package block

import (
	"rand"

	"chunkymonkey/item"
	"chunkymonkey/itemtype"
	. "chunkymonkey/types"
)

// IBlockPlayer defines the interactions that a block aspect may have upon a
// player.
type IBlockPlayer interface {
	OpenWindow(invTypeId InvTypeId)
}

// Defines the behaviour of a block.
type IBlockAspect interface {
	Name() string
	Hit(chunk IChunkBlock, blockLoc *BlockXyz, digStatus DigStatus) (destroyed bool)
	Interact(player IBlockPlayer)
}

type BlockAttrs struct {
	Name         string
	Opacity      int8
	defined      bool
	Destructable bool
	Solid        bool
	Replaceable  bool
	Attachable   bool
}

// The core information about any block type.
type BlockType struct {
	BlockAttrs
	Aspect IBlockAspect
}

// Lookup table of blocks.
type BlockTypeList []BlockType

// Get returns the requested BlockType by ID. ok = false if the block type does
// not exist.
func (btl *BlockTypeList) Get(id BlockId) (block *BlockType, ok bool) {
	if id < 0 || int(id) > len(*btl) {
		ok = false
		return
	}
	block = &(*btl)[id]
	ok = block.defined
	return
}

// MergeBlockItems creates default item types from a defined list of block
// types. It does not override any pre-existing items types.
func (btl *BlockTypeList) CreateBlockItemTypes(itemTypes itemtype.ItemTypeMap) {
	for id := range *btl {
		blockType := &(*btl)[id]
		if !blockType.defined {
			continue
		}
		if _, exists := itemTypes[ItemTypeId(id)]; exists {
			continue
		}

		itemTypes[ItemTypeId(id)] = &itemtype.ItemType{
			Id:       ItemTypeId(id),
			Name:     blockType.Name,
			MaxStack: itemtype.MaxStackDefault,
		}
	}
}

// The interface required of a chunk by block behaviour.
type IChunkBlock interface {
	GetRand() *rand.Rand
	GetItemType(itemTypeId ItemTypeId) (itemType *itemtype.ItemType, ok bool)
	AddItem(item *item.Item)
}

// The distance from the edge of a block that items spawn at in fractional
// blocks.
const blockItemSpawnFromEdge = 4.0 / PixelsPerBlock
