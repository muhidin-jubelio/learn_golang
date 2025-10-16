package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Item struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `gorm:"not null" json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// init DB
func initDB() {
	dsn := "host=localhost user=postgres password=password01 dbname=learn_golang port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established")

	// Migrate the schema
	err = db.AutoMigrate(&Item{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}

func addItem(c fiber.Ctx) error {
	var input Item
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var existingItem Item
	result := db.Where("name = ?", input.Name).First(&existingItem)
	if result.Error == nil {
		existingItem.Quantity = input.Quantity
		existingItem.Price = input.Price
		if err := db.Save(&existingItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update item",
			})
		}
		return c.Status(fiber.StatusOK).JSON(existingItem)
	}

	if err := db.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create item",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(input)
}

func getInventory(c fiber.Ctx) error {
	var items []Item
	if err := db.Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve items",
		})
	}
	return c.Status(fiber.StatusOK).JSON(items)
}

func main() {
	initDB()
	app := fiber.New()

	app.Post("/item", addItem)
	app.Get("/inventory", getInventory)

	log.Fatal(app.Listen(":9200"))
}
