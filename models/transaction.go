package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          string    `json:"transactionId" gorm:"type:varchar(20);primary_key"`
	Valid       bool      `json:"-"`
	Status      string    `json:"status"`
	Amount      int       `json:"amount"`
	AccountCode string    `json:"accountCode"`
	TradeNo     string    `json:"-"`
	Ip          string    `json:"-"`
	Bank        string    `json:"-"`
	BankCode    string    `json:"-"`
	Time        string    `json:"-"`
	Message     string    `json:"message"`
	CodeNo      string    `json:"-"`
	StoreType   int       `json:"-"`
	StoreId     string    `json:"-"`
	Store       string    `json:"-"`
	OrderID     string    `json:"-" gorm:"type:varchar(24)"`
	CreatedAt   time.Time `json:"-"`
}

func (model *Transaction) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	model.CreatedAt = time.Now().In(loc)
	return
}

func CreateTransaction(db *gorm.DB, model *Transaction) (err error) {
	err = db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTransactions(db *gorm.DB, model *[]Transaction) (err error) {
	err = db.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTransaction(db *gorm.DB, model *Transaction, id int) (err error) {
	err = db.Where("id = ?", id).First(model).Error
	if err != nil {
		return err
	}
	return nil
}
