package messaging

import (
	"github.com/forta-protocol/forta-node/config"
)

// Message types
const (
	SubjectAgentsVersionsLatest = "agents.versions.latest"
	SubjectAgentsActionRun      = "agents.action.run"
	SubjectAgentsActionStop     = "agents.action.stop"
	SubjectAgentsStatusRunning  = "agents.status.running"
	SubjectAgentsStatusStopped  = "agents.status.stopped"
	SubjectMetricAgent          = "metric.agent"
)

// AgentPayload is the message payload.
type AgentPayload []config.AgentConfig
