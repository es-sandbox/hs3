package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync/atomic"
)

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

const chanCtrlMsgsSize = 200

var chanCtrlMsgs = make(chan string, chanCtrlMsgsSize)

const chanImageMsgsSize = 200

var chanImageMsgs = make(chan string, chanImageMsgsSize)

var (
	controllerStatus             uint32
	controllerSubscriptionStatus uint32
)

func controller(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	atomic.StoreUint32(&controllerStatus, 1)
	defer atomic.StoreUint32(&controllerStatus, 0)

	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			_ = mt

			log.Printf("recv: %s", message)

			if atomic.LoadUint32(&controllerSubscriptionStatus) == 1 {
				chanCtrlMsgs <- string(message)
			}
		}
	}()

	for {
		msg := <-chanImageMsgs
		err = c.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func controllerSubscription(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	atomic.StoreUint32(&controllerSubscriptionStatus, 1)
	defer atomic.StoreUint32(&controllerSubscriptionStatus, 0)

	go func() {
		for {
			msg := <-chanCtrlMsgs
			err = c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		_ = mt

		log.Printf("recv: %s", message)

		if atomic.LoadUint32(&controllerStatus) == 1 {
			log.Println("android is active so pass to chanImageMsgs")
			chanImageMsgs <- string(message)
		}
	}
}

// ws://localhost:8080/controller
// android -> ws:
// "method": "move", "accel": int, "vector": int
// "method":  "arm", "x": int, "y": int, "z": int, "r": int
// "method": "grab": int
// "method": "song": int
// "method": "water": int

// ws -> android:
// "method": "image", "status":int, "image":string(base64)
