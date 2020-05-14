package main

import fiber "github.com/hendratommy/fiber/v2"

var handler = func(c *fiber.Ctx) {
	c.Send(c.Path())
}

func routeEditor(router fiber.Router) {
	router.Get("/extra", handler)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/thomas", handler)

	grp := app.Group("/v1")
	grp.Get("/api", handler)

	routeEditor(app)

	routeEditor(grp)

	app.Listen(8080)
}
