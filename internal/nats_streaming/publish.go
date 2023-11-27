package nats_streaming

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kiselevms01/wbProject_L0/model"
	"github.com/nats-io/stan.go"
)

type NatsSteamingPublish struct {
	natsClient stan.Conn
	channel    string
}

func NewNatsSteamingPublish(channel string, natsClient stan.Conn) *NatsSteamingPublish {

	natsSteamingClient := NatsSteamingPublish{
		natsClient: natsClient,
		channel:    channel,
	}
	return &natsSteamingClient
}

func (n *NatsSteamingPublish) NatsSteamingSubscribe() {
	file, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}

	var order model.Order

	err = json.Unmarshal(file, &order)
	if err != nil {
		log.Fatalf("failed Unmarshal: %v", err)
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano())) // create a new Rand-type random number generator with its own source
	for {
		genNum := generator.Intn(1000) // use its methods to generate random numbers
		orderUid := fmt.Sprintf("%d", genNum)
		order.OrderUid = orderUid

		fmt.Println("Send Order Uid: ", order.OrderUid)

		bytes, err := json.Marshal(order)
		if err != nil {
			log.Println("ERROR: json.Marshal:", err)
			continue
		}

		err = n.natsClient.Publish(n.channel, bytes)
		if err != nil {
			log.Println("ERROR: conn.Publish:", err)
			continue
		}
		time.Sleep(time.Second * 5)
	}
}
