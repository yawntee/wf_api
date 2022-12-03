package asset

import (
	"bytes"
	"fmt"
)

type PlayerCharacter struct {
}

type PlayerCharacterTable map[int]PlayerCharacter

func (a *Asset) GetPlayerCharacterTable() PlayerCharacterTable {
	if cache, ok := a.Cache["PlayerCharacter"].(PlayerCharacterTable); ok {
		return cache
	}
	reader := a.GetTableFile("/player/player_character")
	Map := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(PlayerCharacterTable)
	for i, params := range Map {
		fmt.Println(i, params)
		table[i] = PlayerCharacter{}
	}
	a.Cache["PlayerCharacter"] = table
	return table
}
