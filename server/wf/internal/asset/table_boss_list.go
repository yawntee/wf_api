package asset

import (
	"bytes"
	"strings"
)

type BossList struct {
	Name string
}

type BossListTable map[int]BossList

func (a *Asset) GetBossListTable() BossListTable {
	if cache, ok := a.Cache["BossList"].(BossListTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/boss_battle_stage_node")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(BossListTable)
	for id, boss := range intMap[1] {
		table[id] = BossList{
			Name: strings.TrimSuffix(boss[0], "讨伐"),
		}
	}
	a.Cache["BossList"] = table
	return table
}
