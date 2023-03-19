package asset

import (
	"bytes"
	"fmt"
	"strconv"
)

type BoxReward = Item

type BoxRewardTable = map[int]map[int]map[int]BoxReward

func (a *Asset) GetBoxRewardTable() BoxRewardTable {
	if cache, ok := a.Cache["BoxReward"].(BoxRewardTable); ok {
		return cache
	}
	reader := a.GetTableFile("/box_gacha/box_reward")
	intMap := parseAnyMap(reader, intKeyParser, func(data []byte) map[int]map[int][]string {
		return parseAnyMap(bytes.NewReader(data), intKeyParser, func(data []byte) map[int][]string {
			return parseIntMap(bytes.NewReader(data))
		})
	})
	table := make(BoxRewardTable)
	for boxGachaId, boxes := range intMap {
		fmt.Println(boxGachaId, boxes)
		boxMap := make(map[int]map[int]BoxReward)
		for boxId, rewardes := range boxes {
			rewardMap := make(map[int]BoxReward)
			for rewardId, reward := range rewardes {
				//numberPerCapsule, err := strconv.Atoi(reward[4])
				//if err != nil {
				//	return nil
				//}
				capsuleNumber, err := strconv.Atoi(reward[5])
				if err != nil {
					return nil
				}
				rewardMap[rewardId] = BoxReward{
					Count: capsuleNumber,
				}
			}
			boxMap[boxId] = rewardMap
		}
		table[boxGachaId] = boxMap
	}
	a.Cache["BoxReward"] = table
	return table
}
