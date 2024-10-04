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
	"github.com/Nukie90/my-fluffy/app/internal/business"
	"github.com/Nukie90/my-fluffy/app/internal/presentation"
	"github.com/Nukie90/my-fluffy/app/internal/repository"
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
	db, err := dbConfig.Connector()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
	}

	//Initialize the repository, notifier and service
	userRepo := repository.UserRepo{DB: db}
	postRepo := repository.PostRepo{DB: db}
	notifier := shared.UserCreationNotifier{}
	adminNotifier := business.NewAdminNotifier(&userRepo)
	notifier.Register(adminNotifier)

	userUsecase := business.NewUserUsecase(&userRepo, &notifier)
	userHandler := presentation.UserHandler{UserUsecase: userUsecase}

	postUsecase := business.NewPostUsecase(&postRepo)
	postHandler := presentation.PostHandler{PostUsecase: postUsecase}

	router := api.NewRouter(&userHandler, &postHandler)

	router.SetupRoutes(a.App)

	a.Get("/swagger/*", fiberSwagger.WrapHandler)

	err = a.Listen(":3000")
	if err != nil {
		return
	}

}

func main() {
	app := New()
	app.Start("env", File, "an environment variable")
}
