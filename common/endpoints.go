package common

const (
	DefaultHttpPort = 8080

	PutEnvironmentInfoEndpoint = "/api/v1/environment"
	PutHumanHeartInfoEndpoint  = "/api/v1/human/heart"
	PutHumanCommonInfoEndpoint = "/api/v1/human/common"
	PutFlowerpotInfoEndpoint   = "/api/v1/flowerpot"
	PutRobotModeEndpoint       = "/api/v1/robot/mode"

	WebsocketEchoEndpoint                   = "/echo"
	WebsocketControllerEndpoint             = "/controller"
	WebsocketControllerSubscriptionEndpoint = "/controller/subscribe"
)
