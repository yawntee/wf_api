package asset

import (
	"bytes"
	"fmt"
)

type Box struct {
}

type BoxTable map[int]Box

func (a *Asset) GetBoxTable() BoxTable {
	if cache, ok := a.Cache["Box"].(BoxTable); ok {
		return cache
	}
	reader := a.GetTableFile("/box_gacha/box")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(BoxTable)
	for i, strings := range intMap {
		fmt.Println(i, strings)
		table[i] = Box{}
	}
	a.Cache["Box"] = table
	return table
}
