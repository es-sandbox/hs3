package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/es-sandbox/hs3/common"
	"github.com/es-sandbox/hs3/message"
)

func TestEnvironmentInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Env()
	objectSlice := common.GetEnv()

	assert(compareInts(len(objectSlice), 1))
	obj := objectSlice[0]

	expected := common.DefaultEnvInfo
	expected.Id = 1
	assert(compareEnvObjects(obj, &expected))
}

func TestHumanHeartInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Hh()
	objectSlice := common.GetHh()

	assert(compareInts(len(objectSlice), 1))
	obj := objectSlice[0]

	expected := common.DefaultHumanHeartInfo
	expected.Id = 1
	assert(compareHhObjects(obj, &expected))
}

func TestHumanCommonInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Hc()
	objectSlice := common.GetHc()

	assert(compareInts(len(objectSlice), 1))
	obj := objectSlice[0]

	expected := common.DefaultHcInfo
	expected.Id = 1
	assert(compareHcObjects(obj, &expected))
}

func TestFlowerpotInfoInfoEndpoint(t *testing.T) {
	removeDBFile()
	server := start()
	defer server.shutdown()

	common.Fp()
	objectSlice := common.GetFp()

	assert(compareInts(len(objectSlice), 1))
	obj := objectSlice[0]

	expected := common.DefaultFlowerpotInfo
	expected.Id = 1
	assert(compareFpObjects(obj, &expected))
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

func assert(value bool) {
	if !value {
		log.Fatal("assertion failed")
	}
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