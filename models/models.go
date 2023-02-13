package models

import (
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Db *gorm.DB
}

type Code struct {
	Code string `json:"code"`
}

type Response struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
	Result  Result `json:"Result"`
}

type Result struct {
	MerchantID      string `json:"MerchantID"`
	Amount          int    `json:"Amt"`
	TradeNo         string `json:"TradeNo"`
	MerchantOrderNo string `json:"MerchantOrderNo"`
	Format          string `json:"RespondType"`
	IP              string `json:"IP"`
	Bank            string `json:"EscrowBank"`
	PaymentType     string `json:"PaymentType"`
	Time            string `json:"PayTime"`
	AccountCode     string `json:"PayBankCode"`
}

type Item struct {
	Amount      int    `json:"amount"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Order struct {
	ID        int64     // primary key, column name is `id`"
	Title     string    // column name is `title`"
	Amount    int       // column name is `amount`"
	Email     string    // column name is `email`"
	CreatedAt time.Time // column name is `created_at`"
	UpdatedAt time.Time // column name is `updated_at`"
}

func (model *Order) BeforeCreate(db *gorm.DB) (err error) {
	model.CreatedAt = time.Now()
	model.CreatedAt = time.Now()
	return
}

type AcOrder struct {
	ID          int64     // primary key, column name is `id`"
	State       string    // column name is `name`"
	Amount      int       // column name is `age`"
	UserId      int       // column name is `age`"
	createdTime time.Time // column name is `created_at`"
}
