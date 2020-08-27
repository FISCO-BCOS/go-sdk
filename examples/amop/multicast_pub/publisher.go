package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
)

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
		log.Fatalf("init publisher failed, err: %v\n", err)
	}
	time.Sleep(3 * time.Second)

	message := "hello, FISCO BCOS, I am multi broadcast publisher!"
	for i := 0; i < 1000; i++ {
		log.Printf("publish message: %s ", message+" "+strconv.Itoa(i))
		err = c.SendAMOPMsg(topic, []byte(message+" "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
		if err != nil {
			log.Printf("PushTopicDataRandom failed, err: %v\n", err)
		}
	}
}
