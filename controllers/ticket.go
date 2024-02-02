package controllers

import (
	"api/helpers"
	"api/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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

	escapedJson := strings.ReplaceAll(fmt.Sprintf("%v", json), "\n", "")
	escapedJson = strings.ReplaceAll(escapedJson, "\r", "")
	fmt.Printf("json: %v\n", escapedJson)

	err := models.CreateTickets(db.Db, &json)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, json)
}

func (db *Db) GetTickets(c *gin.Context) {
	escapedId := strings.ReplaceAll(c.Param("id"), "\n", "")
	escapedId = strings.ReplaceAll(escapedId, "\r", "")
	fmt.Printf("GetTickets() id: %v\n", escapedId)

	tickets := []models.Ticket{}
	err := models.GetTickets(db.Db, &tickets, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("GetTickets() tickets: %v\n", tickets)
	c.JSON(http.StatusOK, tickets)
}
