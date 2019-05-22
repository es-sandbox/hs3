package main

import (
	"encoding/base64"
	"github.com/gorilla/websocket"
	"io/ioutil"
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

const (
	RobotConnected    = "connected"
	RobotDisconnected = "disconnected"
)

var chanRobotStatusEvents = make(chan string, 200)

var (
	controllerStatus             uint32
	controllerSubscriptionStatus uint32
)

func androidIsActive() bool {
	return atomic.LoadUint32(&controllerStatus) == 1
}

func robotIsActive() bool {
	return atomic.LoadUint32(&controllerSubscriptionStatus) == 1
}

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
		select {
		case msg := <-chanImageMsgs:
			err = c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				continue
			}
		case event := <-chanRobotStatusEvents:
			if err := c.WriteMessage(websocket.TextMessage, []byte(event)); err != nil {
				log.Println("write:", err)
				continue
			}
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
	if androidIsActive() {
		chanRobotStatusEvents <- RobotConnected
	}

	defer func() {
		atomic.StoreUint32(&controllerSubscriptionStatus, 0)
		if androidIsActive() {
			chanRobotStatusEvents <- RobotDisconnected
		}
	}()

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

		// send raw image to channel only if android is active
		if atomic.LoadUint32(&controllerStatus) == 1 {
			log.Println("android is active so pass to chanImageMsgs")
			chanImageMsgs <- string(message)
		}

		if string(message) == RobotConnected || string(message) == RobotDisconnected {
			continue
		}

		// save images in server's filesystem always (even android is inactive)
		rawImage, err := base64.StdEncoding.DecodeString(string(message))
		if err != nil {
			log.Println(err)
			continue
		}

		if err := ioutil.WriteFile("images/final.jpg", rawImage, 0666); err != nil {
			log.Printf("can't save image in server's filesystem: %v\n", err)
			continue
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
