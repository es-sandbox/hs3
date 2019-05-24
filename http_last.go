package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func lastEnvironmentInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	//log.Println("new GET request")
	logrus.WithFields(logrus.Fields{
		subsystem:    HTTP,
		requestType:  GET,
		eventType:    environmentInfo,
		"additional": "last",
	}).Info("new request")

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

	logrus.WithFields(logrus.Fields{
		subsystem:   HTTP,
		eventType:   environmentInfo,
		messageType: RAW,
	}).Infof("try to send: %v", string(raw))

	if _, err := w.Write(raw); err != nil {
		log.Println(err)
		return
	}
}

func lastHumanHeartInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	//log.Println("new GET request")
	logrus.WithFields(logrus.Fields{
		subsystem:    HTTP,
		requestType:  GET,
		eventType:    humanHeartInfoEvent,
		"additional": "last",
	}).Info("new request")

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

	logrus.WithFields(logrus.Fields{
		subsystem:   HTTP,
		eventType:   humanHeartInfoEvent,
		messageType: RAW,
	}).Infof("try to send: %v", string(raw))

	if _, err := w.Write(raw); err != nil {
		log.Println(err)
		return
	}
}

func lastHumanCommonInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	//log.Println("new GET request")
	logrus.WithFields(logrus.Fields{
		subsystem:    HTTP,
		requestType:  GET,
		eventType:    humanCommonInfoEvent,
		"additional": "last",
	}).Info("new request")

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

	logrus.WithFields(logrus.Fields{
		subsystem:   HTTP,
		eventType:   humanCommonInfoEvent,
		messageType: RAW,
	}).Infof("try to send: %v", string(raw))

	if _, err := w.Write(raw); err != nil {
		log.Println(err)
		return
	}
}

func lastFlowerpotInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		subsystem:    HTTP,
		requestType:  GET,
		eventType:    flowerpotInfoEvent,
		"additional": "last",
	}).Info("new request")

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

	logrus.WithFields(logrus.Fields{
		subsystem:   HTTP,
		eventType:   flowerpotInfoEvent,
		messageType: RAW,
	}).Infof("try to send: %v", string(raw))

	if _, err := w.Write(raw); err != nil {
		log.Println(err)
		return
	}
}

func lastHeadInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		subsystem:    HTTP,
		requestType:  GET,
		eventType:    headInfoEvent,
		"additional": "last",
	}).Info("new request")
	//
	//envInfo, err := db.GetEnvironmentInfoRecord()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//raw, err := json.Marshal(envInfo)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//logrus.WithFields(logrus.Fields{
	//	subsystem:   HTTP,
	//	eventType:   environmentInfo,
	//	messageType: RAW,
	//}).Infof("try to send: %v", string(raw))
	//
	//if _, err := w.Write(raw); err != nil {
	//	log.Println(err)
	//	return
	//}
}