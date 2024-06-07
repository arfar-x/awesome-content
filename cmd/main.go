package main

import (
	"sync"

	"awesome-content/cmd/http"
	"awesome-content/config"
	"awesome-content/internal/bootstrap"
	"awesome-content/internal/database"
)

var (
	appConfig *config.AppConfig
	app       *bootstrap.App
)

func init() {
	appConfig = config.InitConfig()
	if appConfig == nil {
		panic("Could not initialize configs.")
	}
}

func main() {
	var wg sync.WaitGroup

	defer database.CloseDbClient()
	db, err := database.InitDbClient(appConfig)
	if err != nil {
		panic("Could not open connection to the database")
	}

	app = bootstrap.NewApp(appConfig, db, &wg)

	app.Register(http.NewHttpServer(appConfig, db))

	if err := app.Run(); err != nil {
		panic(err)
	}

	wg.Wait()
}
