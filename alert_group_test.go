package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func alertFixture(id int64) *Alert {
	series := make(chan float32, 3)
	series <- 6.0
	series <- 11.0
	series <- 2.0
	return NewAlert(id, NewSeverityState(
		5.0, 10.0, ThresholdTypeAbove,
	), series)
}

func TestAlertGroup(t *testing.T) {
	alert1 := alertFixture(1)
	alert2 := alertFixture(2)
	alert3 := alertFixture(3)
	alerts := []*Alert{alert1, alert2, alert3}
	alertGroup := NewAlertGroup(alerts)
	alertGroup.evaluate()
	notifier := alertGroup.GetNotifier().(*NotifierStdout)

	assert.Equal(t, notifier.AlertToState, map[int64]AlertState{
		1: AlertStateFiring,
		2: AlertStateFiring,
		3: AlertStateFiring,
	})
	assert.Equal(t, notifier.AlertToSeverity, map[int64]AlertSeverity{
		1: AlertSeverityWarning,
		2: AlertSeverityWarning,
		3: AlertSeverityWarning,
	})

	silenceId := int64(1)
	alertGroup.Silence(silenceId, []string{}, []int64{alert2.Id})
	alertGroup.evaluate()
	assert.Equal(t, notifier.AlertToSeverity, map[int64]AlertSeverity{
		1: AlertSeverityCritical,
		3: AlertSeverityCritical,
	})
}
