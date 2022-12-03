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
	for i, params := range intMap {
		table[i] = CharacterText{
			Name:        params[0],
			Description: params[1],
			Nickname:    params[2],
		}
	}
	a.Cache["CharacterText"] = table
	return table
}
