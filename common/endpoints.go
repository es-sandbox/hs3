package common

const (
	DefaultHttpPort = 8080

	PutEnvironmentInfoEndpoint = "/api/v1/environment"
	PutHumanHeartInfoEndpoint  = "/api/v1/human/heart"
	PutHumanCommonInfoEndpoint = "/api/v1/human/common"
	PutFlowerpotInfoEndpoint   = "/api/v1/flowerpot"
	PutRobotModeEndpoint       = "/api/v1/robot/mode"
	PutHeadEndpoint            = "/api/v1/head"

	GetLastEnvironmentInfoEndpoint = "/api/v1/environment/last"
	GetLastHumanHeartInfoEndpoint  = "/api/v1/human/heart/last"
	GetLastHumanCommonInfoEndpoint = "/api/v1/human/common/last"
	GetLastFlowerpotInfoEndpoint   = "/api/v1/flowerpot/last"
	GetLastHeadInfoEndpoint   = "/api/v1/head/last"

	WebsocketEchoEndpoint                   = "/echo"
	WebsocketControllerEndpoint             = "/controller"
	WebsocketControllerSubscriptionEndpoint = "/controller/subscribe"
)
