package asset

import (
	"bytes"
	"fmt"
)

type CharacterLevel struct {
	TotalExperience uint64
}

type CharacterLevelTable map[int]CharacterLevel

func (a *Asset) GetCharacterLevelTable() CharacterLevelTable {
	if cache, ok := a.Cache["CharacterLevel"].(CharacterLevelTable); ok {
		return cache
	}
	reader := a.GetTableFile("/character/character_level")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(CharacterLevelTable)
	for i, params := range intMap {
		fmt.Println(i, params)
		table[i] = CharacterLevel{}
	}
	a.Cache["CharacterLevel"] = table
	return table
}
