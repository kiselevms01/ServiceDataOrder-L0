package main

import (
	"log"

	"github.com/kiselevms01/wbProject_L0/internal/config"
	"github.com/kiselevms01/wbProject_L0/internal/http"
	"github.com/kiselevms01/wbProject_L0/internal/nats_streaming"
	"github.com/kiselevms01/wbProject_L0/internal/service"
)

func main() {
	cfg := config.MustLoad()

	// connecting to database
	database, err := config.DbConnect(cfg.Database)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}

	log.Printf("connecting to the database OK")

	defer func() {
		dbInstance, _ := database.DB()
		if err = dbInstance.Close(); err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}()

	client := service.NewOrderClient(database)
	err = client.Run()
	if err != nil {
		log.Fatalf("error at Run: %v", err)
	}

	// connecting to nats-streaming
	natsCon, err := config.NatsStreamingConnect(cfg.NatsStreaming)
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}

	log.Printf("connecting to the nats-streaming OK")
	defer natsCon.Close()

	natsSub := nats_streaming.NewSubscribe(cfg.NatsStreaming.Channel, natsCon, client)

	natsSub.NStSubscribe()

	// starting the http server
	httpRouter := http.NewHttpRouter(cfg.Http, client)
	log.Printf("start http-server")
	err = httpRouter.Start()
	if err != nil {
		log.Fatalf("failed to listen the port: %d %v", cfg.Http.Port, err)
	}
}
