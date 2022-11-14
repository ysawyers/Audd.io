package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ysawyers/Audd.io/tree/master/server/controllers/game"
)

func PublicRoutes(app *fiber.App) {

	app.Get("/lobbies/:lobbyId/fetch", game.FetchData)
	app.Post("/lobbies/create", game.CreateLobby)

	app.Get("/sessions/:lobbyId/fetch", game.FetchSessionState)
	app.Post("/sessions/:lobbyId/generate", game.CreateSession)
	app.Delete("/sessions/:lobbyId/drop", game.DropSession)
}
