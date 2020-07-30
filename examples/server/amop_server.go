package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

const (
	topic1 = "subscribe"
)

var (
	c *client.Client
)

func init() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("parse configuration failed, err: %v", err)
	}
	c, err = client.Dial(&configs[0])
	if err != nil {
		log.Fatalf("init client failed, err: %v\n", err)
	}
}

func onPush(data []byte) {
	fmt.Println("\n\n" + string(data))
}

func main() {
	err := c.SubscribeTopic(topic1, onPush)
	if err != nil {
		fmt.Printf("subscribe topic failed, err: %v\n", err)
		return
	}
	fmt.Println("subscribe success")

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	<-killSignal
}
