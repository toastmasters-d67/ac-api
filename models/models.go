package models

import (
	"gorm.io/gorm"
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
	Callback    string `json:"callback"`
	Early       int    `json:"early"`
	Double      int    `json:"double"`
	First       int    `json:"first"`
	Second      int    `json:"second"`
	Banquet     int    `json:"banquet"`
}
