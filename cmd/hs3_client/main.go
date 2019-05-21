package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"github.com/es-sandbox/hs3/common"
)

func main() {
	getEnv := flag.Bool("getenv", false, "")
	env := flag.Bool("env", false, "")
	getHh := flag.Bool("gethh", false, "")
	hh := flag.Bool("hh", false, "")
	getHc := flag.Bool("gethc", false, "")
	hc := flag.Bool("hc", false, "")
	getFp := flag.Bool("getfp", false, "")
	fp := flag.Bool("fp", false, "")
	wsFlag := flag.Bool("ws", false, "")

	flag.Parse()

	if *getEnv {
		common.PrintEnv()
	}

	if *env {
		common.Env()
	}

	if *getHh {
		common.PrintHh()
	}

	if *hh {
		common.Hh()
	}

	if *getHc {
		common.PrintHc()
	}

	if *hc {
		common.Hc()
	}

	if *getFp {
		common.PrintFp()
	}

	if *fp {
		common.Fp()
	}

	if *wsFlag {
		ws()
	}
}

func ws() {
	host := fmt.Sprintf("35.159.53.201:%v", common.DefaultHttpPort)
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   "/echo",
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		err := c.WriteMessage(websocket.TextMessage, []byte("Hello world!"))
		if err != nil {
			log.Println("write:", err)
			return
		}

		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)

		time.Sleep(time.Second)
	}
}
