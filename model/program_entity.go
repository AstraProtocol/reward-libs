package model

import (
	"strconv"
	"time"
)

type ProgramAttributes struct {
	SpecialForNewCustomer  float64 // Special reward amount for new customer
	MinimumPurchasedAmount int64   // Minimum purchased amount for NEW_CUSTOMER 1st order or required purchased amount for ORDER_VALUE type
	MaximumLimitReward     float64 // Maximum reward can be given on ORDER_VALUE
	MinimumPurchasedOrder  int64   // Minimum purchased order count to be compute
	MinimumOrderValue      int64   // Minimum purchased order value to be count
}

const (
	PROGRAM_NEW_CUSTOMER   = "NEW_CUSTOMER"
	PROGRAM_ORDER_VALUE    = "ORDER_VALUE"
	PROGRAM_ORDER_QUANTITY = "ORDER_QUANTITY"

	PROGRAM_USAGE_ONE_TIME  = "ONE_TIME"
	PROGRAM_USAGE_UNLIMITED = "UNLIMITED"
)

type ProgramEntity struct {
	ID              int                          `json:"id" gorm:"primaryKey"`
	SellerId        int                          `json:"sellerId"`
	Status          string                       `json:"status"`
	ProgramType     string                       `json:"programType"`
	ProgramName     string                       `json:"programName"`
	Description     string                       `json:"description"`
	RewardAmount    float64                      `json:"rewardAmount"`
	Budget          float64                      `json:"budget"`
	ContractAddress string                       `json:"contractAddress"`
	CreatedAt       time.Time                    `json:"createdAt"`
	UpdatedAt       time.Time                    `json:"updatedAt"`
	Attributes      []*ProgramAttributeValue     `json:"attributes" gorm:"foreignKey:ProgramId;references:ID"`
	Duration        *ProgramDuration             `json:"duration" gorm:"foreignKey:ProgramId;references:ID"`
	UsageLimit      *ProgramUsageLimit           `json:"usageLimit" gorm:"foreignKey:ProgramId;references:ID"`
	Tiers           []ProgramCustomerEligibility `json:"tiers" gorm:"foreignKey:ProgramId;references:ID"`
}

func (e *ProgramEntity) TableName() string {
	return "cashback_program"
}

func (e *ProgramEntity) ParseProgramAttributes() ProgramAttributes {
	var (
		minimumPurchasedAmount int64
		maximumLimitReward     float64
		specialForNewCustomer  float64
		minimumPurchasedOrder  int64
		minimumOrderValue      int64
	)
	for _, attr := range e.Attributes {
		switch attr.AttributeCode {
		case "minimumPurchasedAmount":
			val, _ := strconv.ParseInt(attr.Value, 10, 64)
			minimumPurchasedAmount = val
		case "maximumLimitReward":
			val, _ := strconv.ParseFloat(attr.Value, 64)
			maximumLimitReward = val
		case "specialForNewCustomer":
			val, _ := strconv.ParseFloat(attr.Value, 64)
			specialForNewCustomer = val
		case "minimumPurchasedOrder":
			val, _ := strconv.ParseInt(attr.Value, 10, 64)
			minimumPurchasedOrder = val
		case "minimumOrderValue":
			val, _ := strconv.ParseInt(attr.Value, 10, 64)
			minimumOrderValue = val
		}
	}

	return ProgramAttributes{
		MinimumPurchasedAmount: minimumPurchasedAmount,
		MaximumLimitReward:     maximumLimitReward,
		SpecialForNewCustomer:  specialForNewCustomer,
		MinimumPurchasedOrder:  minimumPurchasedOrder,
		MinimumOrderValue:      minimumOrderValue,
	}
}

type ProgramAttribute struct {
	ID            int       `json:"id" gorm:"primaryKey"`
	AttributeCode string    `json:"attributeCode"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (e *ProgramAttribute) TableName() string {
	return "cashback_program_attribute"
}

type ProgramAttributeValue struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	AttributeCode string         `json:"attributeCode"`
	ProgramId     int            `json:"programId"`
	Value         string         `json:"value"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	Program       *ProgramEntity `json:"program" gorm:"foreignKey:ProgramId;references:ID"`
}

func (e *ProgramAttributeValue) TableName() string {
	return "cashback_program_attribute_value"
}

type ProgramDuration struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	DurationType string    `json:"durationType"`
	FromDate     time.Time `json:"fromDate"`
	ToDate       time.Time `json:"toDate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ProgramId    int       `json:"programId" gorm:"foreignKey:ProgramId;references:ID"`
}

func (e *ProgramDuration) TableName() string {
	return "cashback_program_duration"
}

type ProgramUsageLimit struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	UsageLimitType string    `json:"usageLimitType"`
	MaximumOrders  int       `json:"maximumOrders"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	ProgramId      int       `json:"programId" gorm:"foreignKey:ProgramId;references:ID"`
}

func (e *ProgramUsageLimit) TableName() string {
	return "cashback_program_usage_limit"
}

type ProgramCustomerEligibility struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	ProgramId    int       `json:"programId" gorm:"foreignKey:ProgramId;references:ID"`
	TierId       int       `json:"tierId"`
	TierName     string    `json:"tierName"`
	RewardAmount float64   `json:"rewardAmount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (e *ProgramCustomerEligibility) TableName() string {
	return "cashback_program_customer_eligibility"
}
