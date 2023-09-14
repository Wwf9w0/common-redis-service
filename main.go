package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Go Redis.")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	/*pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)*/

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/redis", func(c *fiber.Ctx) error {
		pong, err := client.Ping().Result()
		if err != nil {
			fmt.Println(err)
		}
		return c.SendString(pong)
	})

	app.Listen(":3000")
}
