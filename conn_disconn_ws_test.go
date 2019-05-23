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

func TestConnDisconnWs(t *testing.T) {
	server := start()
	defer server.shutdown()

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

	// ------------------------------ VERIFY DISCONNECTED MESSAGE ------------------------------
	_, message, err := androidClient.ReadMessage()
	if err != nil {
		return
	}

	assert(compareStrings(string(message), "disconnected"), "VERIFY DISCONNECTED MESSAGE")
	// -----------------------------------------------------------------------------------------

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
	// ----------------------------------------------------------------------------------

	// ------------------------------ VERIFY CONNECTED MESSAGE ------------------------------
	_, message, err = androidClient.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return
	}

	assert(compareStrings(string(message), "connected"), "VERIFY CONNECTED MESSAGE")
	// --------------------------------------------------------------------------------------
	if err := robotClient.Close(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 500)


	done := make(chan string, 5)
	go func (){
		for i := 0; i < 5; i++ {
			_, message, err = androidClient.ReadMessage()
			if err != nil {
				log.Fatal(err)
				return
			}

			done <- string(message)
		}
	}()


	// ------------------------------ ROBOT RELAUNCHING ------------------------------
	go func() {
		for i := 0; i < 2; i++ {
			robotClient, _, err = websocket.DefaultDialer.Dial(robotUrl.String(), nil)
			if err != nil {
				log.Fatal("dial:", err)
			}

			time.Sleep(time.Millisecond * 500)

			if err := robotClient.Close(); err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Millisecond * 500)
		}
	}()
	// -------------------------------------------------------------------------------

	for i := 0; i < 5; i++ {
		expected := "disconnected"
		if i % 2 == 1 {
			expected = "connected"
		}

		actual := <-done
		fmt.Printf("DEBUG: actual: %v, expected: %v\n", actual, expected)
		assert(compareStrings(actual, expected), "VERIFY CONN/DISCONN MESSAGE")
	}
}