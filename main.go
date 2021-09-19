package main

import (
	"github.com/gofiber/fiber"
	"github.com/adiyogi22/gofiber/jobs"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, Worldd!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/jobs", jobs.GetJobs)
	app.Get("/api/v1/jobs/:id", jobs.GetJob)
	app.Post("/api/v1/jobs", jobs.NewJob)
	app.Delete("/api/v1/jobs/:id", jobs.DeleteJob)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
