package asset

import "fmt"

type BoxGacha struct {
}

type BoxGachaTable map[int]BoxGacha

func (a *Asset) GetBoxGachaTable() BoxGachaTable {
	if cache, ok := a.Cache["BoxGacha"].(BoxGachaTable); ok {
		return cache
	}
	reader := a.GetTableFile("/box_gacha/box_gacha")
	intMap := parseIntMap(reader)
	table := make(BoxGachaTable)
	for i, params := range intMap {
		fmt.Println(i, params)
		table[i] = BoxGacha{}
	}
	a.Cache["BoxGacha"] = table
	return table
}
