package main

import (
	"context"
	"encoding/hex"
	"os"
	"os/signal"
	"time"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 2 {
		logrus.Fatalf("parameters are not enough, example \n%s 127.0.0.1:20200 hello", os.Args[0])
	}
	endpoint := os.Args[1]
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	config := &conf.Config{IsHTTP: false, ChainID: 1, CAFile: "ca.crt", Key: "sdk.key", Cert: "sdk.crt",
		IsSMCrypto: false, GroupID: "group0", PrivateKey: privateKey, NodeURL: endpoint}
	var c *client.Client
	var err error
	for i := 0; i < 3; i++ {
		logrus.Printf("%d try to connect\n", i)
		c, err = client.Dial(config)
		if err != nil {
			logrus.Printf("init subscriber failed, err: %v, retrying\n", err)
			continue
		}
		break
	}
	if err != nil {
		logrus.Fatalf("init subscriber failed, err: %v\n", err)
	}

	timeout := 100 * time.Second
	queryTicker := time.NewTicker(timeout)
	defer queryTicker.Stop()
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = c.SubscribeBlockNumberNotify(ctx, func(blockNUmber int64) {
		logrus.Printf("received: %d\n", blockNUmber)
		//queryTicker.Stop()
		//queryTicker = time.NewTicker(timeout)
		//done <- true
	})

	if err != nil {
		logrus.Printf("subscribe event failed, err: %v\n", err)
		return
	}

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	for {
		select {
		case <-done:
			logrus.Println("Done!")
			os.Exit(0)
		case <-queryTicker.C:
			logrus.Printf("can't receive message after 10s, %s\n", time.Now().String())
			os.Exit(1)
		case <-killSignal:
			logrus.Println("user exit")
			os.Exit(0)
		}
	}
}
