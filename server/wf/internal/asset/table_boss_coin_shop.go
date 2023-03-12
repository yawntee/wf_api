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
	for id, params := range intMap {
		//shopId
		shopId, err := strconv.Atoi(params[0])
		if err != nil {
			panic(err)
		}
		shopItem := a.parseShopItem(params, 6, 31, 13, 24, 27, 14)
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
