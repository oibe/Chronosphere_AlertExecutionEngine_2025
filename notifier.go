package main

import "fmt"

type Notifier interface {
	Notify(*Alert)
}

type NotifierStdout struct {
	AlertToState    map[int64]AlertState
	AlertToSeverity map[int64]AlertSeverity
}

func NewNotifier() *NotifierStdout {
	return &NotifierStdout{
		AlertToState:    make(map[int64]AlertState),
		AlertToSeverity: make(map[int64]AlertSeverity),
	}
}

func (n *NotifierStdout) RemoveAlert(id int64) {
	delete(n.AlertToState, id)
	delete(n.AlertToSeverity, id)
}

func (n *NotifierStdout) Notify(alert *Alert) {
	state, level := alert.GetState(), alert.GetLevel()
	n.AlertToState[alert.Id] = state
	n.AlertToSeverity[alert.Id] = level
	fmt.Printf("alert id: %d, state: %s, level: %s\n", alert.Id, state, level)
}
