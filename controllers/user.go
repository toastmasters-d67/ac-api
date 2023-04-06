package controllers

import (
	"api/helpers"
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func NewUsers() *Db {
	db := helpers.InitDb()
	db.AutoMigrate(&models.User{})
	return &Db{Db: db}
}

func (db *Db) CreateUser(c *gin.Context) {
	fmt.Printf("CreateUser()\n\n")
	json := models.Register{}
	c.BindJSON(&json)
	fmt.Printf("json: %v\n\n", json)

	record, err := models.GetUserByEmail(db.Db, json.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if record.Email == json.Email {
		fmt.Printf("CreateUser() record: %v\n\n", record)
		c.JSON(409, gin.H{
			"error": fmt.Sprintf("email found"),
		})
		return
	}

	user := models.User{}
	user.FirstName = json.FirstName
	user.LastName = json.LastName
	user.Email = json.Email
	user.Password = helpers.GeneratePasswordHash(json.Password)

	err = models.CreateUser(db.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, json)
}

func (db *Db) GetUsers(c *gin.Context) {
	user := []models.User{}
	err := models.GetUsers(db.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (db *Db) AuthenticateUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	fmt.Printf("AuthenticateUser() req: %v\n\n", req)
	user, err := models.GetUserByEmail(db.Db, req.Email)
	if err != nil {
		fmt.Printf("AuthenticateUser() GetUserByEmail err: %v\n\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", req.Email),
		})
		return
	}
	if user.Email != req.Email {
		fmt.Printf("AuthenticateUser() GetUserByEmail: %v\n\n", user.Email)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect account or password",
		})
		return
	}
	err = helpers.PasswordCompare(req.Password, user.Password)
	if err != nil {
		fmt.Printf("AuthenticateUser() PasswordCompare err: %v\n\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect account or password",
		})
		return
	}
	token, err := helpers.GenerateToken(*user)
	if err != nil {
		fmt.Printf("AuthenticateUser() GenerateToken err: %v\n\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (db *Db) GetUser(c *gin.Context) {
	id, _, ok := helpers.GetSession(c)
	fmt.Printf("GetUser() id: %v\n\n", id)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	user, err := models.GetUserById(db.Db, id)
	if err != nil {
		fmt.Printf("GetUser() err: %v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, user)
}
