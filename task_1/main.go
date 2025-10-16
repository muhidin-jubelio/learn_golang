package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/robfig/cron/v3"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("healthy")
	})
	c := cron.New()
	c.Start()
	_, err := c.AddFunc("@every 1s", func() { printName("John", 30) })
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	go func() {
		if err := app.Listen(":9200"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	quit := make(chan struct{})
	<-quit
	log.Println("Shutting down server...")
	_ = app.Shutdown()
	log.Println("Server shut down successfully")

}

func printName(name string, age int) {
	println("My name is "+name+" and Age ", age)
}
