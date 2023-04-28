package model

import (
	"strconv"
	"time"
)

const (
	SELLER_DEFAULT_EXPIRED_TIME      uint64 = 604800
	SELLER_CONFIG_REDEEM_EXPIRE_TIME string = "redeem_expiration_time"
)

type SellerEntity struct {
	ID           int                 `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time           `json:"createdAt"`
	UpdatedAt    time.Time           `json:"updatedAt"`
	Logo         string              `json:"logo"`
	Name         string              `json:"name"`
	ConfigValues []SellerConfigValue `json:"configValues" gorm:"foreignKey:SellerId;references:ID"`
	Tiers        []SellerTier        `json:"tiers" gorm:"foreignKey:SellerId;references:ID"`
}

func (e *SellerEntity) TableName() string {
	return "sellers"
}

type SellerTier struct {
	SellerEntity
	Name       string
	SellerId   int
	Ranking    int
	Conditions string
	Logo       string
}

type SellerConfigValue struct {
	ID         int    `json:"id"`
	ConfigCode string `json:"configCode"`
	Value      string `json:"value"`
	SellerId   int    `json:"sellerId"`
}

func (e *SellerConfigValue) TableName() string {
	return "seller_config_value"
}

func (seller *SellerEntity) GetExpiredTime() uint64 {
	// Then parse seller configValues
	now := uint64(time.Now().Unix())
	var expiredTime uint64 = now + 604800
	for _, conf := range seller.ConfigValues {
		if conf.ConfigCode == SELLER_CONFIG_REDEEM_EXPIRE_TIME {
			parsed, _ := strconv.ParseUint(conf.Value, 10, 64)
			expiredTime = now + parsed
		}
	}
	return expiredTime
}
