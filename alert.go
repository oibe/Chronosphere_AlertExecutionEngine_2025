package main

type AlertState string
type AlertSeverity string

type SeverityStateType string

const (
	AlertStateFiring   AlertState = "firing"
	AlertStateInactive AlertState = "inactive"

	AlertSeverityNormal   AlertSeverity = "normal"
	AlertSeverityWarning  AlertSeverity = "warning"
	AlertSeverityCritical AlertSeverity = "critical"

	ThresholdTypeAbove SeverityStateType = "above"
	ThresholdTypeBelow SeverityStateType = "below"
)

type SeverityState struct {
	warn          float32
	critical      float32
	thresholdType SeverityStateType
}

func NewSeverityState(warn, critical float32, thresholdType SeverityStateType) *SeverityState {
	return &SeverityState{
		warn:          float32(warn),
		critical:      float32(critical),
		thresholdType: thresholdType,
	}
}

func (at *SeverityState) above(val float32) (AlertState, AlertSeverity) {
	if val >= at.critical {
		return AlertStateFiring, AlertSeverityCritical
	}
	if val >= at.warn {
		return AlertStateFiring, AlertSeverityWarning
	}
	return AlertStateInactive, AlertSeverityNormal
}

func (at *SeverityState) below(val float32) (AlertState, AlertSeverity) {
	if val <= at.critical {
		return AlertStateFiring, AlertSeverityCritical
	}
	if val <= at.warn {
		return AlertStateFiring, AlertSeverityWarning
	}
	return AlertStateInactive, AlertSeverityNormal
}

func (at *SeverityState) eval(val float32) (AlertState, AlertSeverity) {
	switch at.thresholdType {
	case ThresholdTypeAbove:
		return at.above(val)
	case ThresholdTypeBelow:
		return at.below(val)
	default:
		panic("Invalid threshold type")
	}
	return "", ""
}

type Alert struct {
	Id            int64
	state         AlertState
	level         AlertSeverity
	SeverityState *SeverityState
	metricChannel chan float32
}

func NewAlert(id int64, SeverityState *SeverityState, metricChannel chan float32) *Alert {
	return &Alert{
		Id:            id,
		state:         AlertStateInactive,
		level:         AlertSeverityNormal,
		SeverityState: SeverityState,
		metricChannel: metricChannel,
	}
}

func (alert *Alert) GetState() AlertState {
	return alert.state
}

func (alert *Alert) GetLevel() AlertSeverity {
	return alert.level
}

func (ag *Alert) Eval() {
	state, level := ag.SeverityState.eval(<-ag.metricChannel)
	ag.state = state
	ag.level = level
}
