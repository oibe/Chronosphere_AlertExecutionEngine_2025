package main

type AlertManager struct {
	alertGroups []*AlertGroup
}

func (am *AlertManager) generateSeries(values []float32) chan float32 {
	series := make(chan float32, len(values))
	for _, value := range values {
		series <- value
	}
	return series
}
func (am *AlertManager) Init() {
	series := am.generateSeries([]float32{6.0, 11.0, 2.0})
	alerts := []*Alert{
		NewAlert(
			1,
			NewSeverityState(5, 10, ThresholdTypeAbove),
			series,
		),
		NewAlert(
			2,
			NewSeverityState(5, 10, ThresholdTypeAbove),
			series,
		),
	}

	alerts2 := []*Alert{
		NewAlert(
			3,
			NewSeverityState(5, 10, ThresholdTypeAbove),
			series,
		),
		NewAlert(
			4,
			NewSeverityState(5, 10, ThresholdTypeAbove),
			series,
		),
	}

	alertGroups := []*AlertGroup{NewAlertGroup(alerts), NewAlertGroup(alerts2)}
	for _, alertGroup := range alertGroups {
		go alertGroup.Run()
	}
}
