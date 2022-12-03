package asset

import "strconv"

type Character struct {
	StringId string
	Rarity   uint8
}

type CharacterTable map[int]Character

func (a *Asset) GetCharacterTable() CharacterTable {
	if cache, ok := a.Cache["Character"].(CharacterTable); ok {
		return cache
	}
	reader := a.GetTableFile("/character/character")
	intMap := parseIntMap(reader)
	table := make(CharacterTable)
	for i, params := range intMap {
		rarity, err := strconv.ParseUint(params[2], 10, 8)
		if err != nil {
			panic(err)
		}
		table[i] = Character{
			StringId: params[0],
			Rarity:   uint8(rarity),
		}
	}
	a.Cache["Character"] = table
	return table
}
