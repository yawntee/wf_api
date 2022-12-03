package asset

type StarGrainShop = ShopItem

type StarGrainShopTable = map[int]StarGrainShop

func (a *Asset) GetStarGrainShopTable() StarGrainShopTable {
	if cache, ok := a.Cache["StarGrainShop"].(StarGrainShopTable); ok {
		return cache
	}
	reader := a.GetTableFile("/shop/star_grain_shop")
	intMap := parseIntMap(reader)
	table := make(StarGrainShopTable)
	for id, params := range intMap {
		item := a.parseShopItem(params, 1, 24, 6, 17, 20, 7)
		item.Id = id
		table[id] = *item
	}
	a.Cache["EventItemShop"] = table
	return table
}
