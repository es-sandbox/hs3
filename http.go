package main

import (
	"encoding/json"
	"github.com/es-sandbox/hs3/message"
	"io/ioutil"
	"log"
	"net/http"
)

func environmentInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("new GET request")

		envInfo, err := db.GetAllEnvironmentInfoRecords()
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

		if err := db.PutEnvironmentInfoRecord(&envInfo); err != nil {
			log.Println(err)
			return
		}
	}
}