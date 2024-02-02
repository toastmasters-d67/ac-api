package models

import (
	"os"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           string         `json:"orderId" gorm:"type:varchar(24);primary_key"`
	Amount       int            `json:"amount" gorm:"type:uint;not null"`
	Email        string         `json:"-" gorm:"index;not null"`
	Description  string         `json:"-" gorm:"type:varchar(255)"`
	Early        int            `json:"early" gorm:"type:uint"`
	Double       int            `json:"double" gorm:"type:uint"`
	First        int            `json:"first" gorm:"type:uint"`
	Second       int            `json:"second" gorm:"type:uint"`
	Banquet      int            `json:"banquet" gorm:"type:uint"`
	UserID       string         `json:"-" gorm:"type:varchar(50)"`
	Tickets      []*Ticket      `json:"tickets"`
	Transactions []*Transaction `json:"transactions"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (model *Order) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.CreatedAt = time.Now().In(loc)
	return
}

func (model *Order) BeforeUpdate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.UpdatedAt = time.Now().In(loc)
	return
}

func GetOrdersByEmail(db *gorm.DB, email string) (*[]Order, error) {
	models := []Order{}
	if res := db.Where("email = ?", email).Find(&models); res.Error != nil {
		return nil, res.Error
	}
	return &models, nil
}

func GetOrderById(db *gorm.DB, id string) (*Order, error) {
	model := Order{}
	if res := db.Preload("Tickets").Where("id = ?", id).Find(&model); res.Error != nil {
		return nil, res.Error
	}
	return &model, nil
}

func CreateOrder(db *gorm.DB, model *Order) (err error) {
	err = db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(db *gorm.DB, model *[]Order) (err error) {
	err = db.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(db *gorm.DB, model *Order, id int) (err error) {
	err = db.Where("id = ?", id).First(model).Error
	if err != nil {
		return err
	}
	return nil
}
