package main

import (
	"encoding/json"
	"github.com/es-sandbox/hs3/message"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
)

//message Head {
//bool movement = 1;
//uint32 ambient = 2;
//float temperature = 3;
//float altitude_meters = 4;
//}

func headInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			requestType: POST,
			eventType:   headInfoEvent,
		}).Info("new request")

		raw, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   headInfoEvent,
			messageType: RAW,
		}).Info(string(raw))

		var head message.Head
		if err := json.Unmarshal(raw, &head); err != nil {
			log.Println(err)
			return
		}

		logrus.WithFields(logrus.Fields{
			subsystem:   HTTP,
			eventType:   headInfoEvent,
			messageType: PARSED,
		}).Info(head)
		//
		//if err := db.PutEnvironmentInfoRecord(&envInfo); err != nil {
		//	log.Println(err)
		//	return
		//}
	}
}