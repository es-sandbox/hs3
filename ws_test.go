package main

import (
	"fmt"
	"github.com/es-sandbox/hs3/common"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strings"
	"testing"
)

const useClosedNetworkConnectionErrorMessage = "use of closed network connection"

func TestAndroidWriteRobotRead(t *testing.T) {
	server := start()
	defer server.shutdown()

	const (
		messagesNum = 10
		messageText = "ctrl"
	)

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
			err := androidClient.WriteMessage(websocket.TextMessage, []byte(messageText))
			if err != nil && !strings.Contains(err.Error(), useClosedNetworkConnectionErrorMessage) {
				log.Fatal(err)
			}
		}
	}()

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

	done := make(chan string, messagesNum)
	for i := 0; i < messagesNum; i++ {
		_, message, err := robotClient.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		done <- string(message)
	}
	for i := 0; i < messagesNum; i++ {
		assert(compareStrings(string(<-done), messageText))
	}
}

func TestAndroidReadRobotWrite(t *testing.T) {
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

	done := make(chan struct{}, 10)
	go func() {
		for {
			_, message, err := androidClient.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)
			_ = message

			done <- struct{}{}
		}
	}()


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

	go func() {
		for i := 0;; i++ {
			err := robotClient.WriteMessage(websocket.TextMessage, []byte("ctrl"))
			if err != nil {
				//log.Fatal(err)
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		<-done
	}
}