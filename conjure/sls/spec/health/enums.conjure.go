// This file was generated by Conjure and should not be manually edited.

package health

import (
	"encoding/json"
	"strings"
)

type HealthState string

const (
	// The service node is fully operational with no issues.
	HealthStateHealthy HealthState = "HEALTHY"
	// The service node is fully operational with no issues; however, it is requesting to defer shutdown or restart. A deferring node should not accept "new" jobs but should allow polling of existing jobs.
	HealthStateDeferring HealthState = "DEFERRING"
	// The service node is no longer serving requests and is ready to be shut down. Nodes in a deferring state are expected to change to a suspended state once they have completed any pending work. A suspended node must also indicate in its readiness probe that it should not receive incoming requests.
	HealthStateSuspended HealthState = "SUSPENDED"
	// The service node is operating in a degraded state, but is capable of automatically recovering. If any of the nodes in the service were to be restarted, it may result in correctness or consistency issues with the service. Ex: When a cassandra node decides it is not up-to-date and needs to repair, the node is operating in a degraded state. Restarting the node prior to the repair being complete might result in the service being unable to correctly respond to requests.
	HealthStateRepairing HealthState = "REPAIRING"
	// The service node is in a state that is trending towards an error. If no corrective action is taken, the health is expected to become an error.
	HealthStateWarning HealthState = "WARNING"
	// The service node is operationally unhealthy.
	HealthStateError HealthState = "ERROR"
	// The service node has entered an unrecoverable state. All nodes of the service should be stopped and no automated attempt to restart the node should be made. Ex: a service fails to migrate to a new schema and is left in an unrecoverable state.
	HealthStateTerminal HealthState = "TERMINAL"
	HealthStateUnknown  HealthState = "UNKNOWN"
)

func (e *HealthState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch strings.ToUpper(s) {
	default:
		*e = HealthStateUnknown
	case "HEALTHY":
		*e = HealthStateHealthy
	case "DEFERRING":
		*e = HealthStateDeferring
	case "SUSPENDED":
		*e = HealthStateSuspended
	case "REPAIRING":
		*e = HealthStateRepairing
	case "WARNING":
		*e = HealthStateWarning
	case "ERROR":
		*e = HealthStateError
	case "TERMINAL":
		*e = HealthStateTerminal
	}
	return nil
}
