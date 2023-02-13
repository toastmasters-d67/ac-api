package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID          string    // primary key, column name is `id`"
	Valid       bool      // column name is `valid`"
	Status      string    // column name is `status`"
	Amount      int       // column name is `amount`"
	AccountCode string    // column name is `account_code`"
	TradeNo     string    // column name is `trade_no`"
	Ip          string    // column name is `ip`"
	Bank        string    // column name is `bank`"
	BankCode    string    // column name is `bank_code`"
	Time        string    // column name is `time`"
	Message     string    // column name is `message`"
	CodeNo      string    // column name is `code_no`"
	StoreType   int       // column name is `store_type`"
	StoreId     string    // column name is `store_id`"
	Store       string    // column name is `store`"
	CreatedAt   time.Time // column name is `created_at`"
	// UpdatedAt   time.Time // column name is `updated_at`"
}

func (model *Transaction) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	model.CreatedAt = time.Now().In(loc)
	// model.UpdatedAt = time.Now().In(loc)
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
