package models

type Setting struct {
	ID                        uint `json:"id"`
	UserID                    uint `json:"user_id"`
	PushNotifications         bool `json:"push_notifications"`
	PromotionsAndOffers       bool `json:"promotions_and_offers"`
	PlatformUpdates           bool `json:"platform_updates"`
	OrderUpdates              bool `json:"order_updates"`
	TransactionNotifications  bool `json:"transaction_notifications"`
	VendorUpdates             bool `json:"vendor_updates"`
	AccountAndSecurityUpdates bool `json:"account_and_security_updates"`
}
