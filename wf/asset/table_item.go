package asset

import (
	"fmt"
	"strconv"
)

type ItemList struct {
	Name     string
	Rarity   int
	MaxCount int
}

type ItemListTable map[int]ItemList

func (a *Asset) GetItemListTable() ItemListTable {
	if cache, ok := a.Cache["ItemList"].(ItemListTable); ok {
		return cache
	}
	reader := a.getTableFile("/item/item")
	intMap := parseIntMap(reader)
	table := make(ItemListTable)
	for i, strings := range intMap {
		rarity, err := strconv.Atoi(strings[14])
		if err != nil {
			panic(err)
		}
		if rarity < 1 || rarity > 5 {
			panic(fmt.Sprintf("%v\n%v", ErrItemRarity, rarity))
		}
		maxCount, err := strconv.Atoi(strings[15])
		if err != nil {
			panic(err)
		}
		table[i] = ItemList{
			Name:     strings[1],
			Rarity:   rarity,
			MaxCount: maxCount,
		}
	}
	a.Cache["ItemList"] = table
	return table
}
