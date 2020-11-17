package twlib_user

type ShopTypeKey int

type UserShop struct {
	UserData *UserShopInfo
}

type UserShopInfo struct {
	BuyCount          int
	Refresh_free      int
	Refresh_price_add int
	GoodsInfo         []int // 商品ID
}
