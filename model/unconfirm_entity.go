package model

type UnconfirmTx struct {
	Txid string `gorm:"primaryKey"`
	Type int
}

func (entity *UnconfirmTx) TableName() string {
	return "unconfirmtx"
}
