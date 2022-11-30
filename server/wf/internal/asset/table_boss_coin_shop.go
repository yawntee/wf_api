package asset

import (
	"strconv"
)

type BossCoinShop struct {
	BossShopId int
	ShopItem
}

type BossCoinShopTable map[int]BossCoinShop

func (a *Asset) GetBossCoinShopTable() BossCoinShopTable {
	if cache, ok := a.Cache["BossCoinShop"].(BossCoinShopTable); ok {
		return cache
	}
	reader := a.GetTableFile("/shop/boss_coin_shop")
	intMap := parseIntMap(reader)
	table := make(BossCoinShopTable)
	for id, strings := range intMap {
		//shopId
		shopId, err := strconv.Atoi(strings[0])
		if err != nil {
			panic(err)
		}
		shopItem := a.parseShopItem(strings, 4, 29, 11, 22, 25, 12)
		shopItem.Id = id
		if err != nil {
			panic(err)
		}
		table[id] = BossCoinShop{
			BossShopId: shopId,
			ShopItem:   *shopItem,
		}
	}
	a.Cache["BossCoinShop"] = table
	return table
}
