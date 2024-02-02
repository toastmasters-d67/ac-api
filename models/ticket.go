package models

import (
	"os"
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID         uint64         `json:"-" gorm:"primary_key;auto_increment"`
	Type       string         `json:"type" gorm:"type:varchar(10);not null"`
	FirstName  string         `json:"firstName" gorm:"type:varchar(50);not null"`
	LastName   string         `json:"lastName" gorm:"type:varchar(50);not null"`
	Known      string         `json:"knownAs" gorm:"type:varchar(50)"`
	Club       string         `json:"club" gorm:"type:varchar(50);not null"`
	Vegetarian bool           `json:"vegetarian" gorm:"default:false"`
	Dtm        bool           `json:"dtm" gorm:"default:false"`
	Banquet    bool           `json:"banquet" gorm:"default:false"`
	OrderID    string         `json:"orderId" gorm:"type:varchar(24)"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (model *Ticket) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.CreatedAt = time.Now().In(loc)
	return
}

func (model *Ticket) BeforeUpdate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.UpdatedAt = time.Now().In(loc)
	return
}

func CreateTickets(db *gorm.DB, models *[]Ticket) (err error) {
	err = db.Create(models).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTickets(db *gorm.DB, models *[]Ticket, id string) (err error) {
	err = db.Where("order_id = ?", id).Find(models).Error
	if err != nil {
		return err
	}
	return nil
}
