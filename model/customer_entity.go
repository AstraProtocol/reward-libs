package model

type CustomerEntity struct {
	BasicEntity
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	WalletAddress string `json:"walletAddress"`
}

func (e *CustomerEntity) TableName() string {
	return "customer"
}

// Check if customer's reward can be sent via wallet
func (e *CustomerEntity) ViaWallet() bool {
	return e.WalletAddress != ""
}

type CustomerSeller struct {
	BasicEntity
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	WalletAddress string `json:"walletAddress"`
	SellerId      int    `json:"sellerId"`
	CustomerId    int    `json:"customerId"`
}

func (e *CustomerSeller) TableName() string {
	return "customer_seller"
}

// Return changes from customer seller to customer
func (e *CustomerSeller) CompareWithCustomer(customer *CustomerEntity) map[string]interface{} {
	changes := make(map[string]interface{})
	if e.WalletAddress != customer.WalletAddress {
		changes["wallet_address"] = e.WalletAddress
	}
	if e.Email != customer.Email {
		changes["email"] = e.Email
	}
	if e.Phone != customer.Phone {
		changes["phone"] = e.Phone
	}
	return changes
}
