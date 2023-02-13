package models

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID          int64 // primary key, column name is `id`"
	Amount      int
	Email       string
	Description string
	CreatedAt   time.Time // column name is `created_at`"
	// UpdatedAt   time.Time // column name is `updated_at`"
}

func (model *Cart) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	model.CreatedAt = time.Now().In(loc)
	// model.UpdatedAt = time.Now().In(loc)
	return
}

func CreateCart(db *gorm.DB, model *Cart) (err error) {
	err = db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCarts(db *gorm.DB, model *[]Cart) (err error) {
	err = db.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCart(db *gorm.DB, model *Cart, id int) (err error) {
	err = db.Where("id = ?", id).First(model).Error
	if err != nil {
		return err
	}
	return nil
}
