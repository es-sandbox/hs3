package main

import (
	"encoding/json"
	"github.com/es-sandbox/hs3/message"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	subsystem   = "subsystem"
	requestType = "request_type"
	eventType   = "event_type"
	messageType = "message_type"

	HTTP = "HTTP"
	GET  = "GET"
	POST = "POST"

	RAW    = "RAW"
	PARSED = "PARSED"

	environmentInfo = "environment_info"
	humanHeartInfoEvent  = "human_heart_info"
)

func environmentInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			requestType: GET,
			eventType:   environmentInfo,
		}).Info("new request")

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
		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			requestType: POST,
			eventType:   environmentInfo,
		}).Info("new request")

		raw, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   environmentInfo,
			messageType: RAW,
		}).Info(string(raw))

		var envInfo message.EnvironmentInfo
		if err := json.Unmarshal(raw, &envInfo); err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   environmentInfo,
			messageType: PARSED,
		}).Info(envInfo)

		if err := db.PutEnvironmentInfoRecord(&envInfo); err != nil {
			log.Println(err)
			return
		}
	}
}

func humanHeartInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			requestType: GET,
			eventType:   humanHeartInfoEvent,
		}).Info("new request")

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
		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			requestType: POST,
			eventType:   humanHeartInfoEvent,
		}).Info("new request")

		raw, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   humanHeartInfoEvent,
			messageType: RAW,
		}).Info(string(raw))

		var humanHeartInfo message.HumanHeartInfo
		if err := json.Unmarshal(raw, &humanHeartInfo); err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   humanHeartInfo,
			messageType: PARSED,
		}).Info(humanHeartInfo)

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

func robotModeEndpoint(w http.ResponseWriter, r *http.Request) {
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

		log.Println("RAW", raw)

		var mode message.RobotMode
		if err := json.Unmarshal(raw, &mode); err != nil {
			log.Println(err)
			return
		}

		log.Println("PARSED", mode)

		if err := db.PutRobotMode(&mode); err != nil {
			log.Println(err)
			return
		}

		chanRobotModeСhanges <- mode.Mode
	}
}
