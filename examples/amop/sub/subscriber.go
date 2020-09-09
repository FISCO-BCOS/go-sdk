package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

func onPush(data []byte, response *[]byte) {
	log.Printf("received: %s\n", string(data))
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("the number of arguments is not equal 4")
	}
	endpoint := os.Args[1]
	topic := os.Args[2]
	config := &conf.Config{IsHTTP: false, ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt", IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58",
		NodeURL:    endpoint}
	c, err := client.Dial(config)
	if err != nil {
		log.Fatalf("init subscriber failed, err: %v\n", err)
	}
	time.Sleep(3 * time.Second)

	err = c.SubscribeTopic(topic, onPush)
	if err != nil {
		fmt.Printf("SubscribeAuthTopic failed, err: %v\n", err)
		return
	}
	fmt.Println("Subscriber success")

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	<-killSignal
}
