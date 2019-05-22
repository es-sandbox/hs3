package main

import (
	"fmt"
	"github.com/es-sandbox/hs3/common"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestWs(t *testing.T) {
	server := start()
	defer server.shutdown()

	host := fmt.Sprintf("localhost:%v", common.DefaultHttpPort)
	androidUrl := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketControllerEndpoint,
	}
	log.Printf("connecting to %s", androidUrl.String())

	androidClient, _, err := websocket.DefaultDialer.Dial(androidUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer androidClient.Close()

	go func() {
		for i := 0;; i++ {
			err := androidClient.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ctrl_%v", i)))
			if err != nil {
				log.Fatal(err)
				return
			}

			time.Sleep(time.Second)
		}
	}()

	//for {
	//	_, message, err := c.ReadMessage()
	//	if err != nil {
	//		log.Println("read:", err)
	//		return
	//	}
	//	log.Printf("recv: %s", message)
	//}

	robotUrl := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketControllerSubscriptionEndpoint,
	}
	log.Printf("connecting to %s", robotUrl.String())

	robotClient, _, err := websocket.DefaultDialer.Dial(robotUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer robotClient.Close()

	log.Println("HERE")

	done := make(chan struct{}, 1)
	for {
		_, message, err := robotClient.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		_ = message

		done <- struct{}{}
		break
	}
	<-done

	log.Println("OK")
}