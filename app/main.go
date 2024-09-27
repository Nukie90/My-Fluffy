// @title			SDA My Fluffy API
// @version		1.0
// @description	This is the API documentation for the SDA My Fluffy API
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
// @host			localhost:3000
// @BasePath		/api/v1
package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/Nukie90/my-fluffy/app/internal/shared"

	"github.com/Nukie90/my-fluffy/app/internal/api"

	_ "github.com/Nukie90/my-fluffy/app/docs"
)

var File string

// App is the main application
type App struct {
	*fiber.App
}

// New creates a new App
func New() *App {
	return &App{fiber.New()}
}

// Start starts the application
func (a *App) Start(name, value, usage string) {
	envFlag := flag.String(name, value, usage)
	flag.Parse()

	configDetail, err := shared.ReadConfigFile(*envFlag)

	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}

	dbConfig := shared.NewGormConfig(configDetail)
	fmt.Println(dbConfig)
	db, err := dbConfig.Connector()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

	api.SetupRoutes(a.App, db)

	a.Get("/swagger/*", fiberSwagger.WrapHandler)

	err = a.App.Listen(":3000")
	if err != nil {
		return
	}
	fmt.Println("Server started on port 3000")
}

func main() {
	app := New()
	app.Start("env", File, "an environment variable")
}
