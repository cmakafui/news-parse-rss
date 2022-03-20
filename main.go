package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Get news items on the root endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("http://ghheadlines.com/rss")
	
		items, _ := json.Marshal(feed.Items)

		err := c.SendString(string(items))

		return err
	})

	// Listen on PORT 300
	app.Listen(":8080")
}
