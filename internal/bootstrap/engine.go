package bootstrap

type EngineInterface interface {
	Boot() error
	Run() error
	Shutdown() error
}
