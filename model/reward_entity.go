package model

const (
	REWARD_STATUS_PENDING   = "pending"
	REWARD_STATUS_COMPLETED = "completed"
	REWARD_STATUS_FAILED    = "failed"
)

type RewardEntity struct {
	BasicEntity
	ProgramId     int     `json:"programId"`
	Amount        float64 `json:"amount"`
	AmountAastra  string  `json:"amountAastra"`
	TierAmount    float64 `json:"tierAmount"`
	Status        string  `json:"status"`
	Email         string  `json:"email"`
	Phone         string  `json:"phone"`
	WalletAddress string  `json:"walletAddress"`
	TxHash        string  `json:"txHash"`
	SellerId      int     `json:"sellerId"`
	CustomerId    int     `json:"customerId"`
	TxError       string  `json:"txError"`
	ImportId      int     `json:"importId"`
	RequestId     string  `json:"requestId"`
	DeliveryId    int     `json:"deliveryId"`
	QtyTraceback  string  `json:"qtyTraceback"`
	Type          string  `json:"type"`
	// TokenAddress string `json:"tokenAddress"`
}

func (e *RewardEntity) TableName() string {
	return "reward"
}
