package main

import (
	"fmt"
	"github.com/es-sandbox/hs3/bolt_db"
	"github.com/es-sandbox/hs3/bolt_db/storage"
	"log"
	"net/http"

	"github.com/boltdb/bolt"

	"github.com/es-sandbox/hs3/common"
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
	http.HandleFunc(common.PutRobotModeEndpoint, robotModeEndpoint)

	http.HandleFunc(common.GetLastEnvironmentInfoEndpoint, lastEnvironmentInfoEndpoint)
	http.HandleFunc(common.GetLastHumanHeartInfoEndpoint, lastHumanHeartInfoEndpoint)
	http.HandleFunc(common.GetLastHumanCommonInfoEndpoint, lastHumanCommonInfoEndpoint)
	http.HandleFunc(common.GetLastFlowerpotInfoEndpoint, lastFlowerpotInfoEndpoint)

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
