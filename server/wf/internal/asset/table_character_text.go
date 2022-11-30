package asset

type CharacterText struct {
	Name        string
	Description string
	Nickname    string
}

type CharacterTextTable map[int]CharacterText

func (a *Asset) GetCharacterTextTable() CharacterTextTable {
	if cache, ok := a.Cache["CharacterText"].(CharacterTextTable); ok {
		return cache
	}
	reader := a.GetTableFile("/character/character_text")
	intMap := parseIntMap(reader)
	table := make(CharacterTextTable)
	for i, strings := range intMap {
		table[i] = CharacterText{
			Name:        strings[0],
			Description: strings[1],
			Nickname:    strings[2],
		}
	}
	a.Cache["CharacterText"] = table
	return table
}
