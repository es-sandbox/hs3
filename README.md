# hs3
## Messages

```go
type EnvironmentInfo struct {
	Id                 uint64
	EnvironmentTemp    float32
	AtmospherePressure int64
	Altitude           float32
	Humidity           uint8
	RobotBatteryLvl    uint8
	Brightness         uint16
}

type HumanHeartInfo struct {
	Id               uint64
	HeartRate        uint8
	HeartRhythm      float32
	DeviceBatteryLvl uint8
}

type HumanCommonInfo struct {
	Id       uint64
	Humidity uint8
	Temp     float32
}

type FlowerpotInfo struct {
	Id       uint64
	PouredOn bool
}

type RobotMode struct {
	Mode uint8
}
```
	
## HTTP calls
```
Get/Post /api/v1/environment
Get/Post /api/v1/human/heart
Get/Post /api/v1/human/common
Get/Post /api/v1/flowerpot
Get/Post /api/v1/robot/mode

GET /api/v1/environment/last
GET /api/v1/human/heart/last
GET /api/v1/human/common/last
GET /api/v1/flowerpot/last
```

## Websocket calls
```
WebsocketEchoEndpoint                   = "/echo"
WebsocketControllerEndpoint             = "/controller"
WebsocketControllerSubscriptionEndpoint = "/controller/subscribe"
```
