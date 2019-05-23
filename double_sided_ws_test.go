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

func TestDoubleSidedWs(t *testing.T) {
	server := start()
	defer server.shutdown()

	const (
		messagesNum = 10
		messageText = "ctrl"
	)

	// ------------------------------ ANDROID CLIENT CONNECT ------------------------------
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
	// ------------------------------------------------------------------------------------

	// ------------------------------ ANDROID CLIENT INFINITELY WRITE ------------------------------
	go func() {
		for i := 0;; i++ {
			err := androidClient.WriteMessage(websocket.TextMessage, []byte(messageText))
			if err != nil && !strings.Contains(err.Error(), useClosedNetworkConnectionErrorMessage) {
				log.Fatal(err)
			}
		}
	}()
	// ---------------------------------------------------------------------------------------------

	// ------------------------------ ROBOT CLIENT CONNECT ------------------------------
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
	// ----------------------------------------------------------------------------------

	// ------------------------------ ROBOT CLIENT INFINITELY WRITE ------------------------------
	go func() {
		for i := 0;; i++ {
			err := robotClient.WriteMessage(websocket.TextMessage, []byte(messageText))
			if err != nil && !strings.Contains(err.Error(), useClosedNetworkConnectionErrorMessage) {
				log.Fatal(err)
			}
		}
	}()
	// -------------------------------------------------------------------------------------------

	// ------------------------------ ROBOT CLIENT READ messagesNum MESSAGES ------------------------------
	done := make(chan string, messagesNum)
	for i := 0; i < messagesNum; i++ {
		_, message, err := robotClient.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		done <- string(message)
	}
	// ----------------------------------------------------------------------------------------------------
	//
	//
	// USE DONE CHANNELS AS INTERMEDIARY BETWEEN ROBOT AND VERIFICATION
	//
	//
	// ------------------------------ VERIFY ROBOT RECEIVED MESSAGES ------------------------------
	for i := 0; i < messagesNum; i++ {
		assert(compareStrings(string(<-done), messageText))
	}
	// --------------------------------------------------------------------------------------------

	// ------------------------------ ANDROID CLIENT READ messagesNum MESSAGES ----------------------------
	go func() {
		for i := 0;; i++ {
			_, message, err := androidClient.ReadMessage()
			if err != nil && !strings.Contains(err.Error(), useClosedNetworkConnectionErrorMessage) {
				log.Fatal(err)
				return
			}

			if string(message) == "connected" || string(message) == "disconnected" {
				continue
			}

			done <- string(message)
		}
	}()
	// ----------------------------------------------------------------------------------------------------
	//
	//
	// USE DONE CHANNELS AS INTERMEDIARY BETWEEN ROBOT AND VERIFICATION
	//
	//
	// ------------------------------ VERIFY ANDROID RECEIVED MESSAGES ----------------------------
	for i := 0; i < messagesNum; i++ {
		assert(compareStrings(string(<-done), messageText))
	}
	// --------------------------------------------------------------------------------------------
}