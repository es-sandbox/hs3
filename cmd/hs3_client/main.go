package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wcharczuk/go-chart" //exposes "chart"

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

	getMode := flag.Bool("getmode", false, "")
	mode := flag.Bool("mode", false, "")

	wsFlag := flag.Bool("ws", false, "")
	wsControllerFlag := flag.Bool("ws_ctrl", false, "")
	wsControllerSubscriptionFlag := flag.Bool("ws_ctrl_sub", false, "")

	wsSendRawImageFlag := flag.Bool("ws_send_raw_image", false, "")

	chartFlag := flag.Bool("chart", false, "")

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

	if *getMode {
		fmt.Println(common.GetMode())
	}

	if *mode {
		common.Mode()
	}

	if *wsFlag {
		ws()
	}

	if *wsControllerFlag {
		wsController()
	}

	if *wsControllerSubscriptionFlag {
		wsControllerSubscription()
	}

	if *wsSendRawImageFlag {
		wsSendRawImage()
	}

	if *chartFlag {
		chartFunc()
	}
}

func ws() {
	host := fmt.Sprintf("35.159.53.201:%v", common.DefaultHttpPort)
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketEchoEndpoint,
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

func wsController() {
	host := fmt.Sprintf("localhost:%v", common.DefaultHttpPort)
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketControllerEndpoint,
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		for i := 0;; i++ {
			err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("ctrl_%v", i)))
			if err != nil {
				log.Println("write:", err)
				return
			}

			time.Sleep(time.Second)
		}
	}()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func wsControllerSubscription() {
	host := fmt.Sprintf("localhost:%v", common.DefaultHttpPort)
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketControllerSubscriptionEndpoint,
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func () {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for i := 0;; i++ {
		err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("image_%v", i)))
		if err != nil {
			log.Println("write:", err)
			return
		}

		time.Sleep(time.Second)
	}
}

func wsSendRawImage() {
	host := fmt.Sprintf("localhost:%v", common.DefaultHttpPort)
	u := url.URL{
		Scheme: "ws",
		Host:   host,
		Path:   common.WebsocketControllerSubscriptionEndpoint,
	}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func () {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	for i := 0;; i++ {
		rawImage, err := ioutil.ReadFile("source.jpg")
		if err != nil {
			log.Fatalf("can't get raw image: %v\n", err)
		}
		rawImageBase64 := base64.StdEncoding.EncodeToString(rawImage)

		if err := c.WriteMessage(websocket.TextMessage, []byte(rawImageBase64)); err != nil {
			log.Println("write:", err)
			return
		}

		time.Sleep(time.Second * 5)
	}
}

func chartFunc() {
	env := common.DefaultEnvInfo
	env.EnvironmentTemp = 3
	common.ExtendedEnv(&env)

	envs := common.GetEnv()
	fmt.Println(envs)

	lenEnvs := len(envs)
	xValues := make([]float64, 0, lenEnvs)
	for i := 1; i <= lenEnvs; i++ {
		xValues = append(xValues, float64(i))
	}
	yValues := make([]float64, 0, lenEnvs)
	for _, env := range envs {
		yValues = append(yValues, float64(env.EnvironmentTemp))
	}

	fmt.Println(len(xValues), len(yValues))
	for _, x := range xValues {
		fmt.Print(x)
	}
	fmt.Println()

	for _, y := range yValues {
		fmt.Print(y)
	}
	fmt.Println()

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	if err := graph.Render(chart.PNG, buffer); err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(buffer)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("chart.png", raw, 0666); err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")
}