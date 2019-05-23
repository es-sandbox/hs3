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

		log.Println("RAW", raw)

		var envInfo message.EnvironmentInfo
		if err := json.Unmarshal(raw, &envInfo); err != nil {
			log.Println(err)
			return
		}

		log.Println("PARSED", envInfo)

		if err := db.PutEnvironmentInfoRecord(&envInfo); err != nil {
			log.Println(err)
			return
		}
	}
}

func humanHeartInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("new GET request")

		hhInfo, err := db.GetAllHumanHeartInfoRecords()
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

		log.Println("RAW", raw)

		var humanHeartInfo message.HumanHeartInfo
		if err := json.Unmarshal(raw, &humanHeartInfo); err != nil {
			log.Println(err)
			return
		}

		log.Println("PARSED", humanHeartInfo)

		if err := db.PutHumanHeartInfo(&humanHeartInfo); err != nil {
			log.Println(err)
			return
		}
	}
}

func humanCommonInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("new GET request")

		hhInfo, err := db.GetAllHumanCommonInfoRecords()
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

		log.Println("RAW", raw)

		var hcInfo message.HumanCommonInfo
		if err := json.Unmarshal(raw, &hcInfo); err != nil {
			log.Println(err)
			return
		}

		log.Println("PARSED", hcInfo)

		if err := db.PutHumanCommonInfo(&hcInfo); err != nil {
			log.Println(err)
			return
		}
	}
}

func flowerpotInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("new GET request")

		fpInfo, err := db.GetAllFlowerpotInfoRecords()
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

		log.Println("RAW", raw)

		var flowerpotInfo message.FlowerpotInfo
		if err := json.Unmarshal(raw, &flowerpotInfo); err != nil {
			log.Println(err)
			return
		}

		log.Println("PARSED", flowerpotInfo)

		if err := db.PutFlowerpotInfo(&flowerpotInfo); err != nil {
			log.Println(err)
			return
		}
	}
}