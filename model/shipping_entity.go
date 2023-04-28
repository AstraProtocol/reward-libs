package model

import "time"

const (
	SHIPPING_QUEUE_NEW     = 0
	SHIPPING_QUEUE_PENDING = 1
	SHIPPING_QUEUE_DONE    = 2
)

// This is a saved for late shipping batch
type ShippingEntity struct {
	ID int `json:"id" gorm:"primaryKey"`
	// 0 - new, 1 - pending, 2 - done, deleting
	QueueStatus int       `json:"queueStatus"`
	CreatedAt   time.Time `json:"createdAt"`
	ViaWallet   bool      `json:"via_wallet"`
	Payload     string    `json:"payload"`
}

func (entity *ShippingEntity) TableName() string {
	return "shippingqueue"
}
