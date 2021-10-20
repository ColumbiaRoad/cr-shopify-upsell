package merchant

type Profile struct {
	AccessToken    string `json:"access_token"`
	ShopURL        string `json:"shop_url"`
	SubscriptionID int64  `json:"subscription_id"`
	ShouldRender   bool   `json:"should_render"`
}
