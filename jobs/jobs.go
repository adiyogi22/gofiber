package jobs

import (
	"github.com/gofiber/fiber"
)

func GetJobs(c *fiber.Ctx) {
	c.Send("All Job")
}

func GetJob(c *fiber.Ctx) {
	c.Send("Single Job")
}

func NewJob(c *fiber.Ctx) {
	c.Send("New Job")
}

func DeleteJob(c *fiber.Ctx) {
	c.Send("Delete Job")
}
