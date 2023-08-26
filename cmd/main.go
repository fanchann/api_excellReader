package main

import (
	"flag"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/patrickmn/go-cache"

	"github.com/fanchann/excelReader/app/controllers"
	"github.com/fanchann/excelReader/app/domain/models"
	"github.com/fanchann/excelReader/app/repositories"
	"github.com/fanchann/excelReader/app/services"
	"github.com/fanchann/excelReader/config"
	"github.com/fanchann/excelReader/lib"
)

var fileConfiguration *string

/*
you can change the configuration file
*/

func init() {
	fileConfiguration = flag.String("c", ".env", "Insert your configuration setting")
	flag.Parse()
}

func main() {
	dbConnection := lib.DatabaseConnection(*fileConfiguration)

	// AutoMigrate
	dbConnection.AutoMigrate(&models.Customers{})

	// repo
	customerRepo := repositories.NewCustomersRepoImpl(dbConnection)

	// cache
	cache := cache.New(5*time.Minute, 10*time.Minute)

	// service
	customerService := services.NewCustomerServiceImpl(customerRepo, cache)

	// controller
	customercontroller := controllers.NewCustomerCntrlImpl(customerService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	customercontroller.Route(app)

	err := app.Listen(":9000")
	if err != nil {
		panic(err)
	}
}
