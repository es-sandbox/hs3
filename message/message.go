package message

import "encoding/json"

type EnvironmentInfo struct {
	Id                 uint64
	EnvironmentTemp    float32
	AtmospherePressure int64
	Altitude           float32
	Humidity           uint8
	RobotBatteryLvl    uint8
	Brightness         uint16
}

func (e *EnvironmentInfo) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func NewEnvironmentInfoFromBytes(bytes []byte) (*EnvironmentInfo, error) {
	var envInfo EnvironmentInfo
	if err := json.Unmarshal(bytes, &envInfo); err != nil {
		return nil, err
	}
	return &envInfo, nil
}

type HumanHeartInfo struct {
	Id               uint64
	HeartRate        uint8
	HeartRhythm      float32
	DeviceBatteryLvl uint8
}

func (e *HumanHeartInfo) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func NewHumanHeartInfoFromBytes(bytes []byte) (*HumanHeartInfo, error) {
	var hhInfo HumanHeartInfo
	if err := json.Unmarshal(bytes, &hhInfo); err != nil {
		return nil, err
	}
	return &hhInfo, nil
}

type HumanCommonInfo struct {
	Id       uint64
	Humidity uint8
	Temp     float32
}

func (hc *HumanCommonInfo) Encode() ([]byte, error) {
	return json.Marshal(hc)
}

func NewHumanCommonInfoFromBytes(bytes []byte) (*HumanCommonInfo, error) {
	var hcInfo HumanCommonInfo
	if err := json.Unmarshal(bytes, &hcInfo); err != nil {
		return nil, err
	}
	return &hcInfo, nil
}

type FlowerpotInfo struct {
	Id       uint64
	PouredOn bool
	Humidity uint8
}

func (f *FlowerpotInfo) Encode() ([]byte, error) {
	return json.Marshal(f)
}

func NewFlowerpotInfoFromBytes(bytes []byte) (*FlowerpotInfo, error) {
	var fpInfo FlowerpotInfo
	if err := json.Unmarshal(bytes, &fpInfo); err != nil {
		return nil, err
	}
	return &fpInfo, nil
}

type RobotMode struct {
	Mode uint8
}

func (mode *RobotMode) Encode() ([]byte, error) {
	return json.Marshal(mode)
}

func NewRobotModeFromBytes(bytes []byte) (*RobotMode, error) {
	var mode RobotMode
	if err := json.Unmarshal(bytes, &mode); err != nil {
		return nil, err
	}
	return &mode, nil
}

type Image struct {
	Raw []byte
}

type RobotMovement struct {
	Acceleration int16
	Vector       int8
}

type ManipulatorMovement struct {
	X        int16
	Y        int16
	Z        int16
	Rotation int16
	Touch    bool
}
