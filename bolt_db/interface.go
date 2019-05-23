package bolt_db

import "github.com/es-sandbox/hs3/message"

type getAll interface {
	GetAllEnvironmentInfoRecords() ([]*message.EnvironmentInfo, error)
	GetAllHumanHeartInfoRecords() ([]*message.HumanHeartInfo, error)
	GetAllHumanCommonInfoRecords() ([]*message.HumanCommonInfo, error)
	GetAllFlowerpotInfoRecords() ([]*message.FlowerpotInfo, error)
}

type get interface {
	GetEnvironmentInfoRecord() (*message.EnvironmentInfo, error)
	GetHumanHeartInfoRecord() (*message.HumanHeartInfo, error)
	GetHumanCommonInfoRecord() (*message.HumanCommonInfo, error)
	GetFlowerpotInfoRecord() (*message.FlowerpotInfo, error)
	GetRobotMode() (*message.RobotMode, error)
}

type put interface {
	PutEnvironmentInfoRecord(envInfo *message.EnvironmentInfo) error
	PutHumanHeartInfo(humanHeartInfo *message.HumanHeartInfo) error
	PutHumanCommonInfo(hcInfo *message.HumanCommonInfo) error
	PutFlowerpotInfo(flowerpotInfo *message.FlowerpotInfo) error
	PutRobotMode(mode *message.RobotMode) error
}

type Store interface {
	getAll
	get
	put
}
