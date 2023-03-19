package asset

import (
	"bytes"
	"time"
	"wf_api/util"
)

/*
{
	1:{
		id:{
			chapter:{}
			}
	}
}
*/

type BossQuest struct {
	Name      string
	StartTime time.Time
	EndTime   *time.Time
}

type BossQuestTable map[int][]BossQuest

func (a *Asset) GetBossQuestTable() BossQuestTable {
	if cache, ok := a.Cache["BossQuest"].(BossQuestTable); ok {
		return cache
	}
	reader := a.GetTableFile("/quest/boss_battle_quest")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int]map[int][]string {
		return parseAnyMap(bytes.NewReader(data), intKeyParser, func(data []byte) map[int][]string {
			return parseIntMap(bytes.NewReader(data))
		})
	})
	table := make(BossQuestTable)
	for id, chapter := range intMap[1] {
		var quests []BossQuest
		for _, params := range chapter {
			startTime := util.ParseIso(params[5])
			var endTime *time.Time
			if endTimeStr := params[6]; endTimeStr != "(None)" {
				iso := util.ParseIso(endTimeStr)
				endTime = &iso
			}
			quests = append(quests, BossQuest{
				Name:      params[2],
				StartTime: startTime,
				EndTime:   endTime,
			})
		}
		table[id] = quests
	}
	a.Cache["BossQuest"] = table
	return table
}
