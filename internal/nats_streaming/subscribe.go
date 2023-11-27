package nats_streaming

import (
	"log"

	"github.com/kiselevms01/wbProject_L0/internal/service"
	"github.com/nats-io/stan.go"
)

type Subscribe struct {
	channel    string
	natsClient stan.Conn
	client     *service.Client
}

func NewSubscribe(channel string, natsClient stan.Conn, client *service.Client) *Subscribe {

	nStClient := Subscribe{
		channel:    channel,
		natsClient: natsClient,
		client:     client,
	}
	return &nStClient
}

func (n *Subscribe) NStSubscribe() {

	_, err := n.natsClient.Subscribe(
		n.channel, func(m *stan.Msg) {
			err := n.client.AddOrder(m.Data)
			if err != nil {
				log.Printf("Error order adding %s", err)
			}
		})
	if err != nil {
		log.Fatal(err)
	}
}
