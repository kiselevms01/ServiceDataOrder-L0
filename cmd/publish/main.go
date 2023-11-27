package main

import (
	"fmt"
	"log"

	"github.com/kiselevms01/wbProject_L0/internal/config"
	"github.com/kiselevms01/wbProject_L0/internal/nats_streaming"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// if err != nil {
	// 	log.Fatalf("failed to Config %v", err)
	// }

	natsCon, err := config.NatsStreamingConnect(cfg.NatsStreaming)
	if err != nil {
		log.Fatalf("failed to connect to nats-streaming: %v", err)
	}

	log.Printf("connecting to the nats-streaming: success")

	defer func() {
		if err = natsCon.Close(); err != nil {
			log.Fatalf("failed to close nats-streaming: %v", err)
		}
	}()

	natsStreaming := nats_streaming.NewNatsSteamingPublish(cfg.NatsStreaming.Channel, natsCon)
	natsStreaming.NatsSteamingSubscribe()

}
