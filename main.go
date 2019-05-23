package main

import (
	"encoding/json"
	"fmt"
	"github.com/es-sandbox/hs3/bolt_db"
	"github.com/es-sandbox/hs3/bolt_db/storage"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"

	"github.com/es-sandbox/hs3/common"
	"github.com/es-sandbox/hs3/message"
)

var db bolt_db.Store

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	boltDb, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDb.Close()

	db = storage.New(boltDb)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Welcome to my website!"); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.PutEnvironmentInfoEndpoint, environmentInfoEndpoint)

	http.HandleFunc(common.PutHumanHeartInfoEndpoint, humanHeartInfoEndpoint)

	http.HandleFunc(common.PutHumanCommonInfoEndpoint, humanCommonInfoEndpoint)

	http.HandleFunc(common.PutFlowerpotInfoEndpoint, flowerpotInfoEndpoint)

	http.HandleFunc(common.PutRobotModeEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			fpInfo, err := db.GetRobotMode()
			if err != nil {
				log.Println(err)
				return
			}

			raw, err := json.Marshal(fpInfo)
			if err != nil {
				log.Println(err)
				return
			}

			if _, err := w.Write(raw); err != nil {
				log.Println(err)
				return
			}
		case "POST":
			log.Println("new POST request")

			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}

			var mode message.RobotMode
			if err := json.Unmarshal(raw, &mode); err != nil {
				log.Println(err)
				return
			}

			log.Println(mode)

			if err := db.PutRobotMode(&mode); err != nil {
				log.Println(err)
				return
			}

			chanRobotModeСhanges <- mode.Mode
		}
	})

	http.HandleFunc(common.GetLastEnvironmentInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Println("new GET request")

		envInfo, err := db.GetEnvironmentInfoRecord()
		if err != nil {
			log.Println(err)
			return
		}

		raw, err := json.Marshal(envInfo)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := w.Write(raw); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.GetLastHumanHeartInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Println("new GET request")

		hhInfo, err := db.GetHumanHeartInfoRecord()
		if err != nil {
			log.Println(err)
			return
		}

		raw, err := json.Marshal(hhInfo)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := w.Write(raw); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.GetLastHumanCommonInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Println("new GET request")

		hcInfo, err := db.GetHumanCommonInfoRecord()
		if err != nil {
			log.Println(err)
			return
		}

		raw, err := json.Marshal(hcInfo)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := w.Write(raw); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.GetLastFlowerpotInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Println("new GET request")

		flowerpotInfo, err := db.GetFlowerpotInfoRecord()
		if err != nil {
			log.Println(err)
			return
		}

		raw, err := json.Marshal(flowerpotInfo)
		if err != nil {
			log.Println(err)
			return
		}

		if _, err := w.Write(raw); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.WebsocketEchoEndpoint, echo)
	http.HandleFunc(common.WebsocketControllerEndpoint, controller)
	http.HandleFunc(common.WebsocketControllerSubscriptionEndpoint, controllerSubscription)

	fs := http.FileServer(http.Dir("images/"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	log.Println("Start HTTP Server")
	httpAddr := fmt.Sprintf("0.0.0.0:%v", common.DefaultHttpPort)
	if err := http.ListenAndServe(httpAddr, nil); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
