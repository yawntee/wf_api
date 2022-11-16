package module

import (
	"wf_api/wf/internal/context"
)

type ShopModule struct {
	*context.ClientContext
}

func NewShopModule(clientContext *context.ClientContext) *ShopModule {
	return &ShopModule{ClientContext: clientContext}
}

type ShopStock struct {
	SalesList []struct {
		ShopItemId           int `json:"shop_item_id"`
		StockQuantity        int `json:"stock_quantity"`
		TodayPurchaseNum     int `json:"today_purchase_num"`
		ThisMonthPurchaseNum int `json:"this_month_purchase_num"`
		TotalPurchaseNum     int `json:"total_purchase_num"`
		GroupInfo            struct {
			GroupTotalStockQuantity int  `json:"group_total_stock_quantity"`
			GroupTotalPurchaseNum   int  `json:"group_total_purchase_num"`
			MultiStage              bool `json:"multi_stage"`
		} `json:"group_info"`
		ShopType int `json:"shop_type"`
	} `json:"sales_list"`
}

func (c *ShopModule) GetSaleList() {

}
