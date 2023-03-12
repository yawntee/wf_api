package asset

import (
	"bytes"
	"strconv"
	"strings"
)

type Boss struct {
	Name       string
	BossShopId int
}

type BossTable map[int]Boss

func (a *Asset) GetBossTable() BossTable {
	if cache, ok := a.Cache["Boss"].(BossTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/boss_battle_stage_node")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int][]string {
		return parseIntMap(bytes.NewReader(data))
	})
	table := make(BossTable)
	for id, params := range intMap[1] {
		bossShopId, err := strconv.Atoi(params[6])
		if err != nil {
			panic(err)
		}
		table[id] = Boss{
			Name:       strings.TrimSuffix(params[1], "讨伐"),
			BossShopId: bossShopId,
		}
	}
	a.Cache["Boss"] = table
	return table
}
