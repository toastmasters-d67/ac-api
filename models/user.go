package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"os"
	"time"
)

type Register struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID         string    `json:"-" gorm:"type:varchar(50);primary_key"`
	FirstName  string    `json:"firstName" gorm:"type:varchar(50);not null"`
	LastName   string    `json:"lastName" gorm:"type:varchar(50);not null"`
	Email      string    `json:"email" gorm:"uniqueIndex;not null"`
	Password   string    `json:"-" gorm:"type:varchar(255);not null"`
	IsVerified bool      `json:"isVerified" gorm:"default:false"`
	Orders     []*Order  `json:"orders"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

func (model *User) BeforeCreate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.ID = uuid.New().String()
	// model.IsVerified = false
	model.CreatedAt = time.Now().In(loc)
	// model.UpdatedAt = time.Now().In(loc)
	return
}

func (model *User) BeforeUpdate(db *gorm.DB) (err error) {
	loc, _ := time.LoadLocation(os.Getenv("DB_TIMEZONE"))
	model.UpdatedAt = time.Now().In(loc)
	return
}

func CreateUser(db *gorm.DB, model *User) (err error) {
	err = db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, model *[]User) (err error) {
	err = db.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *gorm.DB, id string) (*User, error) {
	user := User{}
	// if res := db.Where("id = ?", id).Find(&user); res.Error != nil {
	if res := db.Preload("Orders").Preload("Orders.Transactions").Where("id = ?", id).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}
	if res := db.Where("email = ?", email).Find(&user); res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
