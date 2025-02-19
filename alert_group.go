package main

import "time"

type AlertGroup struct {
	notifier        Notifier
	alerts          []*Alert
	alertToSilences map[int64]int64
	silences        []*Silence
	clock           *time.Ticker
	done            chan bool
}

func NewAlertGroup(alerts []*Alert) *AlertGroup {
	return &AlertGroup{
		notifier:        NewNotifier(),
		alerts:          alerts,
		clock:           time.NewTicker(1 * time.Second),
		done:            make(chan bool),
		alertToSilences: make(map[int64]int64),
	}
}

func (ag *AlertGroup) GetNotifier() Notifier {
	return ag.notifier
}

func (ag *AlertGroup) Silence(silenceId int64, teams []string, alertIds []int64) {
	silence := &Silence{
		id:    silenceId,
		teams: teams,
	}
	ag.silences = append(ag.silences, silence)

	for _, alertId := range alertIds {
		ag.alertToSilences[int64(alertId)] = silenceId
		notifier := ag.notifier.(*NotifierStdout)
		notifier.RemoveAlert(int64(alertId))
	}
}

func (ag *AlertGroup) evaluate() {
	for _, alert := range ag.alerts {
		alert.Eval()

		if _, exists := ag.alertToSilences[alert.Id]; !exists {
			// todo: maybe add more nuanced notifying for silenced alerts and subrouting by team??
			ag.notifier.Notify(alert)
		}
	}
}

func (ag *AlertGroup) Run() {
	defer ag.clock.Stop() // Stop the ticker when the function exits

	func() {
		for {
			select {
			case <-ag.clock.C:
				ag.evaluate() // Call the function to be executed
			case <-ag.done:
				return
			}
		}
	}()
}

func (ag *AlertGroup) Stop() {
	ag.done <- true
}
