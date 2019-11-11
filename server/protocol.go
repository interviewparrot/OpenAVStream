package server

// Naming convention in the package.
// When a action is required we prefix it with CMD (Command)
// When a data or information is passed we prefix variable as MSG(Message)
// Telemetry command to send some telemetry data
const CMD_Telemetry = "Telemetry"
const CMD_Auth = "Auth"

//Client commands
const (
	CMD_StartSession = "StartSession"
	CMD_ReceiveChunk = "ReceiveChunk"
	CMD_EndSession = "EndSession"
	)

// Admin commands
const(
CMD_MonitorSession = "MonitorSession"
CMD_JoinSession = "JoinSession" // Join to an existing user session

)

// Server commands
const (
	CMD_ReceiveSessionId = "ReceiveSessionId"
)

type ClientMsg struct {
	Command string
	SessionId string
	AuthToken string
	Data string
}

type ServerMsg struct {
	Command string
	Data string
	SessionId string
}

type StartConversationData struct {
	Name string
	Type string
	Version string
}

type TelemetryData struct {
	Type string
	Data string

}

type ClientLogData struct {

}

type ClientMetricsData struct {

}