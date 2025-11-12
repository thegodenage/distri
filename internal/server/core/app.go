package core

type AppConfiguration struct {
}

type App struct {
	cfg *AppConfiguration
}

func NewApp(cfg *AppConfiguration) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) SetHandler(h any) {}
