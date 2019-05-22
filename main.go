package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/boltdb/bolt"

	"github.com/es-sandbox/hs3/common"
	"github.com/es-sandbox/hs3/message"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	boltDb, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDb.Close()

	db := newDb(boltDb)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "Welcome to my website!"); err != nil {
			log.Println(err)
			return
		}
	})

	http.HandleFunc(common.PutEnvironmentInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			envInfo, err := db.getAllEnvironmentInfoRecords()
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
		case "POST":
			log.Println("new POST request")

			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}

			var envInfo message.EnvironmentInfo
			if err := json.Unmarshal(raw, &envInfo); err != nil {
				log.Println(err)
				return
			}

			log.Println(envInfo)

			if err := db.putEnvironmentInfoRecord(&envInfo); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.HandleFunc(common.PutHumanHeartInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			hhInfo, err := db.getAllHumanHeartInfoRecords()
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
		case "POST":
			log.Println("new POST request")

			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}

			var humanHeartInfo message.HumanHeartInfo
			if err := json.Unmarshal(raw, &humanHeartInfo); err != nil {
				log.Println(err)
				return
			}

			log.Println(humanHeartInfo)

			if err := db.putHumanHeartInfo(&humanHeartInfo); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.HandleFunc(common.PutHumanCommonInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			hhInfo, err := db.getAllHumanCommonInfoRecords()
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
		case "POST":
			log.Println("new POST request")

			raw, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}

			var hcInfo message.HumanCommonInfo
			if err := json.Unmarshal(raw, &hcInfo); err != nil {
				log.Println(err)
				return
			}

			log.Println(hcInfo)

			if err := db.putHumanCommonInfo(&hcInfo); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.HandleFunc(common.PutFlowerpotInfoEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			fpInfo, err := db.getAllFlowerpotInfoRecords()
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

			var flowerpotInfo message.FlowerpotInfo
			if err := json.Unmarshal(raw, &flowerpotInfo); err != nil {
				log.Println(err)
				return
			}

			log.Println(flowerpotInfo)

			if err := db.putFlowerpotInfo(&flowerpotInfo); err != nil {
				log.Println(err)
				return
			}
		}
	})

	http.HandleFunc(common.PutRobotModeEndpoint, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Println("new GET request")

			fpInfo, err := db.getRobotMode()
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

			if err := db.putRobotMode(&mode); err != nil {
				log.Println(err)
				return
			}
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

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
