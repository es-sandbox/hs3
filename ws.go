package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
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
			return
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

var chanRobotModeСhanges = make(chan uint8, 200)

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

	msg := RobotDisconnected
	if robotIsActive() {
		msg = RobotConnected
	}
	if err = c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Println("write:", err)
		return
	}

	quit := make(chan struct{}, 0)
	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				quit <- struct{}{}
				break
			}
			_ = mt

			logrus.WithFields(logrus.Fields{
				subsystem: WS,
				Source:    AndroidSource,
				Direction: Read,
			}).Info(string(message))

			if atomic.LoadUint32(&controllerSubscriptionStatus) == 1 {
				chanCtrlMsgs <- string(message)
			}
		}
	}()

	for {
		select {
		case msg := <-chanImageMsgs:
			//logBigMessage(fmt.Sprintf("WRITE TO ANDROID: %v", msg))
			logrus.WithFields(logrus.Fields{
				subsystem: WS,
				Source:    AndroidSource,
				Direction: Write,
			}).Info(cutMessage(msg))
			err = c.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case event := <-chanRobotStatusEvents:
			log.Printf("WRITE TO ANDROID %v", event)
			logrus.WithFields(logrus.Fields{
				subsystem:    WS,
				Source:       AndroidSource,
				Direction:    Write,
				"additional": "robot status event",
			}).Info(event)
			if err := c.WriteMessage(websocket.TextMessage, []byte(event)); err != nil {
				log.Println("write:", err)
				return
			}
		//case mode := <-chanRobotModeСhanges:
		//	//{"method":mode, "mode": int }
		//	type robotMode struct {
		//		Method string
		//		Mode   uint8
		//	}
		//	obj := robotMode{
		//		Method: "mode",
		//		Mode:   mode,
		//	}
		//	raw, err := json.Marshal(obj)
		//	if err != nil {
		//		log.Println(err)
		//	}
		//	if err := c.WriteMessage(websocket.TextMessage, []byte(raw)); err != nil {
		//		log.Println("write:", err)
		//		return
		//	}
		case <-quit:
			return
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
			select {
			case msg := <-chanCtrlMsgs:
				err = c.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("write:", err)
					return
				}

				logrus.WithFields(logrus.Fields{
					subsystem: WS,
					Source:    RobotSource,
					Direction: Write,
				}).Info(msg)
			case mode := <-chanRobotModeСhanges:
				//{"method":mode, "mode": int }
				type robotMode struct {
					Method string `json:"method"`
					Mode   uint8  `json:"mode"`
				}
				obj := robotMode{
					Method: "mode",
					Mode:   mode,
				}
				raw, err := json.Marshal(obj)
				if err != nil {
					log.Println(err)
				}
				if err := c.WriteMessage(websocket.TextMessage, []byte(raw)); err != nil {
					log.Println("write:", err)
					return
				}
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

		//log.Printf("recv: %s", message)
		//logBigMessage(fmt.Sprintf("recv: %s", message))
		logrus.WithFields(logrus.Fields{
			subsystem: WS,
			Source:    RobotSource,
			Direction: Read,
		}).Info(cutMessage(string(message)))

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

func cutMessage(msg string) string {
	if len(msg) <= 50 {
		return msg
	} else {
		return msg[:50]
	}
}

func logBigMessage(msg string) {
	if len(msg) <= 50 {
		log.Println(msg)
	} else {
		log.Println(msg[:50])
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
