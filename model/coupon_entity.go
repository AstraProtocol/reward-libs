package model

import (
	"time"
)

const ()

type CouponEntity struct {
	ID                     int       `json:"id" gorm:"primaryKey"`
	SellerId               int       `json:"sellerId"`
	Status                 string    `json:"status"`
	ProgramType            string    `json:"programType"`
	ProgramName            int8      `json:"programName"`
	Description            int       `json:"description"`
	CouponPrice            int       `json:"couponPrice"`
	TotalCouponCode        int       `json:"totalCouponCode"`
	FromDate               time.Time `json:"fromDate"`
	ToDate                 time.Time `json:"toDate"`
	RedeemLimitType        string    `json:"redeemLimitType"`
	DiscountCodeType       string    `json:"discountCodeType"`
	MinimumPurchasedAmount float64   `json:"minimumPurchasedAmount"`
	ContractAddress        string    `json:"contractAddress"`
}
