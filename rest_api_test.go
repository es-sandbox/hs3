package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/es-sandbox/hs3/common"
	"github.com/es-sandbox/hs3/message"
)

func TestGetLastEnvironmentInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	diffObj := &message.EnvironmentInfo{
		Id:                 7,
		EnvironmentTemp:    6,
		AtmospherePressure: 5,
		Altitude:           4,
		Humidity:           3,
		RobotBatteryLvl:    2,
		Brightness:         1,
	}
	common.ExtendedEnv(diffObj)
	common.ExtendedEnv(diffObj)
	common.ExtendedEnv(diffObj)
	common.ExtendedEnv(diffObj)
	common.Env()
	obj := common.GetLastEnv()

	expected := common.DefaultEnvInfo
	expected.Id = 5
	fmt.Println(obj)
	assert(compareEnvObjects(obj, &expected), "TestEnvironmentInfoEndpoint: compareEnvObjects")
}

func TestGetLastHumanHeartInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	diffObj := &message.HumanHeartInfo{
		Id:               4,
		HeartRate:        3,
		HeartRhythm:      2,
		DeviceBatteryLvl: 1,
	}

	common.ExtendedHh(diffObj)
	common.ExtendedHh(diffObj)
	common.ExtendedHh(diffObj)
	common.ExtendedHh(diffObj)
	common.Hh()
	obj := common.GetLastHh()

	expected := common.DefaultHumanHeartInfo
	expected.Id = 5
	fmt.Println(obj)
	assert(compareHhObjects(obj, &expected), "TestEnvironmentInfoEndpoint: compareHhObjects")
}

func TestGetLastHumanCommonInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	diffObj := &message.HumanCommonInfo{
		Id:       3,
		Humidity: 2,
		Temp:     1,
	}

	common.ExtendedHc(diffObj)
	common.ExtendedHc(diffObj)
	common.ExtendedHc(diffObj)
	common.ExtendedHc(diffObj)
	common.Hc()
	obj := common.GetLastHc()


	expected := common.DefaultHcInfo
	expected.Id = 5
	fmt.Println(obj)
	assert(compareHcObjects(obj, &expected), "TestEnvironmentInfoEndpoint: compareHcObjects")
}

func TestEnvironmentInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Env()
	objectSlice := common.GetEnv()

	fmt.Printf("DEBUG: len(objectSlice): %v, expected: %v\n", len(objectSlice), 1)
	assert(compareInts(len(objectSlice), 1), "TestEnvironmentInfoEndpoint: compareInts")
	obj := objectSlice[0]

	expected := common.DefaultEnvInfo
	expected.Id = 1
	assert(compareEnvObjects(obj, &expected), "TestEnvironmentInfoEndpoint: compareEnvObjects")
}

func TestHumanHeartInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Hh()
	objectSlice := common.GetHh()

	assert(compareInts(len(objectSlice), 1), "TestHumanHeartInfoEndpoint: compareInts")
	obj := objectSlice[0]

	expected := common.DefaultHumanHeartInfo
	expected.Id = 1
	assert(compareHhObjects(obj, &expected), "TestHumanHeartInfoEndpoint: compareHhObjects")
}

func TestHumanCommonInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Hc()
	objectSlice := common.GetHc()

	assert(compareInts(len(objectSlice), 1), "TestHumanCommonInfoEndpoint: compareInts")
	obj := objectSlice[0]

	expected := common.DefaultHcInfo
	expected.Id = 1
	assert(compareHcObjects(obj, &expected), "TestHumanCommonInfoEndpoint: compareHcObjects")
}

func TestFlowerpotInfoInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Fp()
	objectSlice := common.GetFp()

	assert(compareInts(len(objectSlice), 1), "TestFlowerpotInfoInfoEndpoint: compareInts")
	obj := objectSlice[0]

	expected := common.DefaultFlowerpotInfo
	expected.Id = 1
	assert(compareFpObjects(obj, &expected), "TestFlowerpotInfoInfoEndpoint: compareFpObjects")
}

func TestRobotModeEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Mode()
	obj := common.GetMode()

	expected := common.DefaultRobotMode
	assert(compareRobotModeObjects(obj, &expected), "TestRobotModeEndpoint: compareRobotModeObjects")
}

func removeDBFile() {
	if err := os.Remove("my.db"); err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		log.Fatal(err)
	}
}

type server struct {
	cmd *exec.Cmd
}

func start() *server {
	cmd := exec.Command("hs3")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	return &server{
		cmd: cmd,
	}
}

func (s *server) shutdown() {
	if err := s.cmd.Process.Kill(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
}

func assert(value bool, message string) {
	if !value {
		log.Fatal("assertion failed", message)
	}
}

func compareStrings(actual, expected string) bool {
	return actual == expected
}

func compareInts(actual, expected int) bool {
	return actual == expected
}

func compareEnvObjects(actual, expected *message.EnvironmentInfo) bool {
	ok1 := actual.Id == expected.Id
	ok2 := actual.EnvironmentTemp == expected.EnvironmentTemp
	ok3 := actual.AtmospherePressure == expected.AtmospherePressure
	ok4 := actual.Altitude == expected.Altitude
	ok5 := actual.Humidity == expected.Humidity
	ok6 := actual.RobotBatteryLvl == expected.RobotBatteryLvl
	ok7 := actual.Brightness == expected.Brightness
	return ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7
}

func compareHhObjects(actual, expected *message.HumanHeartInfo) bool {
	ok1 := actual.Id == expected.Id
	ok2 := actual.HeartRate == expected.HeartRate
	ok3 := actual.HeartRhythm == expected.HeartRhythm
	ok4 := actual.DeviceBatteryLvl == expected.DeviceBatteryLvl
	return ok1 && ok2 && ok3 && ok4
}

func compareHcObjects(actual, expected *message.HumanCommonInfo) bool {
	ok1 := actual.Id == expected.Id
	ok2 := actual.Humidity == expected.Humidity
	ok3 := actual.Temp == expected.Temp
	return ok1 && ok2 && ok3
}

func compareFpObjects(actual, expected *message.FlowerpotInfo) bool {
	ok1 := actual.Id == expected.Id
	ok2 := actual.PouredOn == expected.PouredOn
	ok3 := actual.Humidity == expected.Humidity
	return ok1 && ok2 && ok3
}

func compareRobotModeObjects(actual, expected *message.RobotMode) bool {
	ok1 := actual.Mode == expected.Mode
	return ok1
}
