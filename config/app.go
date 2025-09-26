package config

import "gorm.io/gorm"

type App struct {
	Env *Env
	DB  *gorm.DB
}

func NewApp() (*App, error) {
	env, err := NewEnv()
	if err != nil {
		return nil, err
	}

	db := NewDatabaseConfig(env)

	return &App{
		Env: env,
		DB:  db,
	}, nil
}

func (app *App) CloseDatabase() {
	CloseDatabase(app.DB)
}
