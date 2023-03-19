package asset

import "fmt"

type LevelCap struct {
}

type LevelCapTable map[int]LevelCap

func (a *Asset) GetLevelCapTable() LevelCapTable {
	if cache, ok := a.Cache["LevelCap"].(LevelCapTable); ok {
		return cache
	}
	reader := a.GetTableFile("/character/level_cap")
	intMap := parseIntMap(reader)
	table := make(LevelCapTable)
	for id, params := range intMap {
		fmt.Println(id, params)
		table[id] = LevelCap{}
	}
	a.Cache["LevelCap"] = table
	return table
}
