package controllers

import (
	"api/helpers"
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
)

func NewTickets() *Db {
	db := helpers.InitDb()
	db.AutoMigrate(&models.Ticket{})
	return &Db{Db: db}
}

func (db *Db) CreateTicket(c *gin.Context) {
	_, _, ok := helpers.GetSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	json := []models.Ticket{}
	c.BindJSON(&json)
	fmt.Printf("json: %v\n\n", json)

	err := models.CreateTickets(db.Db, &json)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, json)
}

func (db *Db) GetTickets(c *gin.Context) {
	fmt.Printf("GetTickets() id: %v\n\n", c.Param("id"))
	tickets := []models.Ticket{}
	err := models.GetTickets(db.Db, &tickets, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("GetTickets() tickets: %v\n\n", tickets)
	c.JSON(http.StatusOK, tickets)
}
