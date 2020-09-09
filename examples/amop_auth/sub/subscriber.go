package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	privateKey1 = "b8ac1f6be271cd4ad5644615dea54823f7fa5d860fb8d3ae8b24141d9f1b9486"
	privateKey2 = "0fe5e3ce06d6d48ec806ea17d13ce3d80e74b85f23c32c38f2c8e4180f539a7e"
	privateKey3 = "13e3531ac291bcf5674acd1c8c7c77b725dc9bf56242b02ef76bf970190412aa"
)

func onPush(data []byte, response *[]byte) {
	log.Printf("received: %s\n", string(data))
}

var (
	c *client.Client
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("the number of arguments is not equal 1")
	}
	endpoint := os.Args[1]
	topic := os.Args[2]
	config := &conf.Config{IsHTTP: false, ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt", IsSMCrypto: false, GroupID: 1,
		PrivateKey: "145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58",
		NodeURL:    endpoint}
	c, err := client.Dial(config)
	if err != nil {
		log.Fatalf("init client failed, err: %v\n", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKey1)
	if err != nil {
		log.Fatalf("hex to ECDSA failed, err: %v", privateKey)
	}
	err = c.SubscribePrivateTopic(topic, privateKey, onPush)
	if err != nil {
		log.Fatalf("SubscribeAuthTopic failed, err: %v\n", err)
	}
	fmt.Println("Subscriber success")

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	<-killSignal
}
