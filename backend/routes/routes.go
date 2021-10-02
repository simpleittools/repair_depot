package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/simpleittools/repair_depot/backend/controllers"
)

func Routing(app *fiber.App) {
	app.Get("/", controllers.Index)

	dashboard := app.Group("/dashboard")
	// TODO: Look into the status monitor more
	//https://github.com/gofiber/fiber/tree/master/middleware/monitor
	dashboard.Get("status", monitor.New())
}
