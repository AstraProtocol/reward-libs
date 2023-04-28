package model

import (
	"time"
)

const (
	NOTI_NOT_SEND_YET = "NotSendYet"
	NOTI_SUCCESS      = "Success"
	NOTI_FAIL         = "Failed"
)

type DeliveryEntity struct {
	//gorm.Model
	ID              uint           `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time      `json:"createdAt"`
	TxHash          string         `json:"txHash"`
	Response        string         `json:"response"`
	ClaimAt         int64          `json:"claimAt"`
	ClaimTxHash     string         `json:"claimTxHash"`
	WithdrawTxHash  string         `json:"withdrawTxHash"`
	HolderAddress   string         `json:"holderAddress"`
	CustomerAddress string         `json:"customerAddress"`
	RedeemCode      string         `json:"redeemCode"`
	RedeemExpiredAt int64          `json:"redeemExpiredAt"`
	TokenAddress    string         `json:"tokenAddress"`
	Amount          string         `json:"amount"`
	Email           string         `json:"email"`
	PhoneNumber     string         `json:"phoneNumber"`
	NotiStatus      string         `json:"notiStatus"`
	ShippingType    string         `json:"shippingType"`
	Receipt         *ReceiptEntity `json:"receipt" gorm:"foreignKey:TxHash;references:TxHash"`
}

func (e *DeliveryEntity) TableName() string {
	return "deliveries"
}
