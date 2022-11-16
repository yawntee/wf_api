package asset

import (
	"strconv"
)

type BossCoinShop struct {
	BossShopId int
	ShopItem
}

type BossCoinShopTable map[int]BossCoinShop

func (a *CacheAsset) GetBossCoinShopTable() BossCoinShopTable {
	if cache, ok := a.Cache["BossCoinShop"].(BossCoinShopTable); ok {
		return cache
	}
	reader := a.getTableFile("/shop/boss_coin_shop")
	intMap := parseIntMap(reader)
	table := make(BossCoinShopTable)
	for i, strings := range intMap {
		//id
		id, err := strconv.Atoi(strings[0])
		if err != nil {
			panic(err)
		}
		shopItem := a.parseShopItem(strings, 29, 11, 22, 12)
		if err != nil {
			panic(err)
		}
		table[i] = BossCoinShop{
			BossShopId: id,
			ShopItem:   *shopItem,
		}
	}
	a.Cache["BossCoinShop"] = table
	return table
}

func (t BossCoinShopTable) GetShopType() int {
	return 7
}
