package main

import (
	"os"

	"github.com/ozonebg/juicetomator/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/ozonebg/juicetomator/internal/config"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	juiceController := controllers.NewJuiceController()

	r.POST("/get-juices", juiceController.HandleMultiJuices)
	r.POST("/get-juice", juiceController.HandleSingleJuice)

	return r
}

func run(c *cli.Context) error {
	config := config.NewJuiceConfig(c)

	r := setupRouter()

	return r.Run(config.HTTPPort)
}

func main() {
	app := &cli.App{
		Name:        "juice-server",
		Version:     "1.0.0",
		Description: "Server to handle juice requests",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "http-port",
				Value: ":8080",
				Usage: "the http port to listen on",
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("failed to run application")
	}
}
