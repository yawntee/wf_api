package asset

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type ItemList struct {
	Name     string
	Rarity   uint8
	MaxCount int
}

type ItemListTable map[int]ItemList

func (a *Asset) GetItemListTable() ItemListTable {
	if cache, ok := a.Cache["ItemList"].(ItemListTable); ok {
		return cache
	}
	reader := a.GetTableFile("/item/item")
	intMap := parseIntMap(reader)
	table := make(ItemListTable)
	for i, params := range intMap {
		rarity, err := strconv.ParseUint(params[15], 10, 8)
		if err != nil {
			panic(errors.WithMessage(err, params[15]))
		}
		if rarity < 1 || rarity > 5 {
			panic(fmt.Sprintf("%v\n%v", ErrItemRarity, rarity))
		}
		maxCount, err := strconv.Atoi(params[16])
		if err != nil {
			panic(errors.WithMessage(err, params[16]))
		}
		table[i] = ItemList{
			Name:     params[1],
			Rarity:   uint8(rarity),
			MaxCount: maxCount,
		}
	}
	a.Cache["ItemList"] = table
	return table
}
