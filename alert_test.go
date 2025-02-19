package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlertEvaluation(t *testing.T) {
	series := make(chan float32, 3)
	series <- 6.0
	series <- 11.0
	series <- 2.0
	alert := NewAlert(1, NewSeverityState(
		5.0, 10.0, ThresholdTypeAbove,
	), series)
	assert.Equal(t, alert.GetLevel(), AlertSeverityNormal)
	assert.Equal(t, alert.GetState(), AlertStateInactive)
	alert.Eval()
	assert.Equal(t, alert.GetLevel(), AlertSeverityWarning)
	assert.Equal(t, alert.GetState(), AlertStateFiring)
	alert.Eval()
	assert.Equal(t, alert.GetLevel(), AlertSeverityCritical)
	assert.Equal(t, alert.GetState(), AlertStateFiring)
	alert.Eval()
	assert.Equal(t, alert.GetLevel(), AlertSeverityNormal)
	assert.Equal(t, alert.GetState(), AlertStateInactive)
}
