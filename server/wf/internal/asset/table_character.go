package asset

type Character struct {
	StringId string
}

type CharacterTable map[int]Character

func (a *Asset) GetCharacterTable() CharacterTable {
	if cache, ok := a.Cache["Character"].(CharacterTable); ok {
		return cache
	}
	reader := a.GetTableFile("/character/character")
	intMap := parseIntMap(reader)
	table := make(CharacterTable)
	for i, strings := range intMap {
		table[i] = Character{
			StringId: strings[0],
		}
	}
	a.Cache["Character"] = table
	return table
}
