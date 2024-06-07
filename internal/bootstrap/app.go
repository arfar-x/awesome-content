package bootstrap

import (
	"fmt"
	"sync"

	"awesome-content/config"
	"gorm.io/gorm"
)

type App struct {
	Config  *config.AppConfig
	DB      *gorm.DB
	Engines []EngineInterface
	wg      *sync.WaitGroup
}

func NewApp(config *config.AppConfig, db *gorm.DB, wg *sync.WaitGroup) *App {
	return &App{
		Config:  config,
		DB:      db,
		Engines: []EngineInterface{},
		wg:      wg,
	}
}

// Run the engines/servers; mostly defined within `cmd` package as application entry points.
func (app *App) Run() error {
	for _, engine := range app.Engines {

		e := engine

		if err := e.Boot(); err != nil {
			return fmt.Errorf("failed to boot engine %s: %w", e, err)
		}

		app.wg.Add(1)

		go func() {
			defer app.wg.Done()

			defer func() {
				if err := e.Shutdown(); err != nil {
					panic(fmt.Sprintf("Engine %s could not shutdown gracefully: %s", e, err.Error()))
				}
			}()

			if err := e.Run(); err != nil {
				panic(fmt.Sprintf("App could not run the engine %s: %s", e, err.Error()))
			}
		}()
	}

	return nil
}

func (app *App) Register(engine EngineInterface) {
	app.Engines = append(app.Engines, engine)
}
