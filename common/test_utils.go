package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/es-sandbox/hs3/message"
)

var (
	DefaultEnvInfo = message.EnvironmentInfo{
		EnvironmentTemp:    1,
		AtmospherePressure: 2,
		Altitude:           3,
		Humidity:           4,
		RobotBatteryLvl:    5,
		Brightness:         6,
	}

	DefaultHumanHeartInfo = message.HumanHeartInfo{
		HeartRate:        1,
		HeartRhythm:      2,
		DeviceBatteryLvl: 3,
	}

	DefaultHcInfo = message.HumanCommonInfo{
		Humidity: 1,
		Temp:     2,
	}

	DefaultFlowerpotInfo = message.FlowerpotInfo{
		PouredOn: false,
		Humidity: 1,
	}
	DefaultRobotMode = message.RobotMode{
		Mode: 1,
	}
)

func PrintEnv() {
	for _, envInfo := range GetEnv() {
		fmt.Println(envInfo)
	}
}

func GetLastEnv() *message.EnvironmentInfo {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, GetLastEnvironmentInfoEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result *message.EnvironmentInfo
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func GetEnv() []*message.EnvironmentInfo {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutEnvironmentInfoEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result []*message.EnvironmentInfo
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func Env() {
	raw, err := json.Marshal(DefaultEnvInfo)
	if err != nil {
		log.Fatal(err)
	}
	rawBuff := bytes.NewReader(raw)

	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutEnvironmentInfoEndpoint)
	contentType := "application/json"
	if _, err := http.Post(url, contentType, rawBuff); err != nil {
		log.Fatal(err)
	}
}

func PrintHh() {
	for _, hhInfo := range GetHh() {
		fmt.Println(hhInfo)
	}
}


func GetHh() []*message.HumanHeartInfo {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutHumanHeartInfoEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result []*message.HumanHeartInfo
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func Hh() {
	raw, err := json.Marshal(DefaultHumanHeartInfo)
	if err != nil {
		log.Fatal(err)
	}
	rawBuff := bytes.NewReader(raw)

	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutHumanHeartInfoEndpoint)
	contentType := "application/json"
	if _, err := http.Post(url, contentType, rawBuff); err != nil {
		log.Fatal(err)
	}
}

func PrintHc() {
	for _, hcInfo := range GetHc() {
		fmt.Println(hcInfo)
	}
}

func GetHc() []*message.HumanCommonInfo {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutHumanCommonInfoEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result []*message.HumanCommonInfo
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func Hc() {
	raw, err := json.Marshal(DefaultHcInfo)
	if err != nil {
		log.Fatal(err)
	}
	rawBuff := bytes.NewReader(raw)

	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutHumanCommonInfoEndpoint)
	contentType := "application/json"
	if _, err := http.Post(url, contentType, rawBuff); err != nil {
		log.Fatal(err)
	}
}

func PrintFp() {
	for _, fpInfo := range GetFp() {
		fmt.Println(fpInfo)
	}
}

func GetFp() []*message.FlowerpotInfo {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutFlowerpotInfoEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result []*message.FlowerpotInfo
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func Fp() {
	fpInfo := message.FlowerpotInfo{
		PouredOn: false,
		Humidity: 1,
	}
	raw, err := json.Marshal(fpInfo)
	if err != nil {
		log.Fatal(err)
	}
	rawBuff := bytes.NewReader(raw)

	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutFlowerpotInfoEndpoint)
	contentType := "application/json"
	if _, err := http.Post(url, contentType, rawBuff); err != nil {
		log.Fatal(err)
	}
}

func GetMode() *message.RobotMode {
	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutRobotModeEndpoint)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	raw, err := ioutil.ReadAll(resp.Body)

	var result message.RobotMode
	if err := json.Unmarshal(raw, &result); err != nil {
		log.Fatal(err)
	}

	return &result
}

func Mode() {
	raw, err := json.Marshal(DefaultRobotMode)
	if err != nil {
		log.Fatal(err)
	}
	rawBuff := bytes.NewReader(raw)

	url := fmt.Sprintf("http://localhost:%v%v", DefaultHttpPort, PutRobotModeEndpoint)
	contentType := "application/json"
	if _, err := http.Post(url, contentType, rawBuff); err != nil {
		log.Fatal(err)
	}
}