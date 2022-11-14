package game

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ysawyers/Audd.io/tree/master/server/configs"
	"github.com/ysawyers/Audd.io/tree/master/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var lobbiesCollection *mongo.Collection = configs.GetCollection(configs.DB, "Game_Lobbies")

type GameLobby struct {
	NumberOfRounds int `json:"numberOfRounds"`
	LengthOfRound  int `json:"lengthOfRound"`
	PlayerMax      int `json:"playerMax"`
}

func FetchData(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var gameLobby models.GameLobby
	lobbyId, _ := primitive.ObjectIDFromHex(c.Params("lobbyId"))

	err := lobbiesCollection.FindOne(ctx, bson.M{"_id": lobbyId}).Decode(&gameLobby)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Lobby does not exist!",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"gameLobby": gameLobby,
	})
}

func CreateLobby(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gameLobby := new(GameLobby)

	if err := c.BodyParser(gameLobby); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not parse JSON body on request",
		})
	}

	new_lobby := models.GameLobby{
		NumberOfRounds: gameLobby.NumberOfRounds,
		LengthOfRound:  gameLobby.LengthOfRound,
		PlayerMax:      gameLobby.PlayerMax,
	}

	result, err := lobbiesCollection.InsertOne(ctx, new_lobby)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error: MongoDB INSERT",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"generatedLobbyId": result.InsertedID,
	})
}
