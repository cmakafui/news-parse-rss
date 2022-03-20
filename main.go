package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Send a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("http://ghheadlines.com/rss")
	
		err := c.SendString(feed.String())
		return err
	})

	// Listen on PORT 300
	app.Listen(":8080")
}
