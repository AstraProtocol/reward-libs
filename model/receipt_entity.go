package model

import "time"

const (
	SHIPPING_TYPE_WALLET = "WALLET"
	SHIPPING_TYPE_OTHER  = "EMAIL-SMS"
)

// This is a receipt for a batch shipping rewards
// only created when success
type ReceiptEntity struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// Encoded data send from client
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"time"`
	SellerId  uint      `json:"sellerId"`
	// Calling source
	Source string `json:"source"`
	// Receipt transaction hash
	TxHash string `json:"txHash"`
	// Shipping id for checking if shipping has been success or not
	ShippingId uint `json:"shippingId"`
	// Calling program contract
	ProgramAddress string `json:"programAddress"`
	// Type of shipping batch, WALLET or EMAIL-SMS
	ShippingType string `json:"shippingType"`
	// Transaction Status: 1 - Waiting, 2 - Success, 3 - Failure
	TxStatus  uint8  `json:"txStatus"`
	RequestId string `json:"requestId"`
}

func (e *ReceiptEntity) TableName() string {
	return "receipts"
}
