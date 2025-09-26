package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/go/authGo/api/route"
	"github.com/naol86/go/authGo/config"
)

func main() {
	app, err := config.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	db := app.DB
	defer app.CloseDatabase()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second
	r := gin.Default()

	route.SetUpRoutes(app.Env, timeout, db, r)
	r.Run(app.Env.ServerPort)
}
