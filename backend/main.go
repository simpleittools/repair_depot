package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/simpleittools/repair_depot/backend/config"
	"log"
)

func main() {
	PORT := config.EnvVariables("PORT")

	// Dev Status
	status := config.EnvVariables("STATUS")
	switch {
	case status == "Development":
		fmt.Println("The current status is Development")
	case status == "Production":
		fmt.Println("You are in production")
	default:
		log.Fatalf("Status is undefined. Please modify your .env file and set your Status Variable. Please match " +
			"the variable options as defined in the case statements per the main.go Dev Status section")
	}

	// GoFiber template Engine
	engine := html.New("./views", ".gohtml")
	if status == "Development" {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Confirming",
		}, "layouts/main")
	})
	fmt.Println(fmt.Sprintf("Starting application, available at http://localhost%s", PORT))
	log.Fatal(app.Listen(PORT))

}
