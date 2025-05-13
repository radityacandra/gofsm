package examples

import (
	"context"

	"github.com/radityacandra/gofsm"
)

func buildStateDiagram() map[string]gofsm.Transitions {
	return map[string]gofsm.Transitions{
		"CLOSED": {
			"APP_PASSIVE_OPEN": gofsm.TransitionAction{
				To: "LISTEN",
			},
			"APP_ACTIVE_OPEN": gofsm.TransitionAction{
				To: "SYN_SENT",
			},
		},
		"LISTEN": {
			"RCV_SYN": gofsm.TransitionAction{
				To: "SYN_RCVD",
			},
			"APP_SEND": gofsm.TransitionAction{
				To: "SYN_SENT",
			},
			"APP_CLOSE": gofsm.TransitionAction{
				To: "CLOSED",
			},
		},
		"SYN_RCVD": {
			"APP_CLOSE": gofsm.TransitionAction{
				To: "FIN_WAIT_1",
			},
			"RCV_ACK": gofsm.TransitionAction{
				To: "ESTABLISHED",
			},
		},
		"SYN_SENT": {
			"RCV_SYN": gofsm.TransitionAction{
				To: "SYN_RCVD",
			},
			"RCV_SYN_ACK": gofsm.TransitionAction{
				To: "ESTABLISHED",
			},
			"APP_CLOSE": gofsm.TransitionAction{
				To: "CLOSED",
			},
		},
		"ESTABLISHED": {
			"APP_CLOSE": gofsm.TransitionAction{
				To: "FIN_WAIT_1",
			},
			"RCV_FIN": gofsm.TransitionAction{
				To: "CLOSE_WAIT",
			},
		},
		"FIN_WAIT_1": {
			"RCV_FIN": gofsm.TransitionAction{
				To: "CLOSING",
			},
			"RCV_FIN_ACK": gofsm.TransitionAction{
				To: "TIME_WAIT",
			},
			"RCV_ACK": gofsm.TransitionAction{
				To: "FIN_WAIT_2",
			},
		},
		"CLOSING": {
			"RCV_ACK": gofsm.TransitionAction{
				To: "TIME_WAIT",
			},
		},
		"FIN_WAIT_2": {
			"RCV_FIN": gofsm.TransitionAction{
				To: "TIME_WAIT",
			},
		},
		"TIME_WAIT": {
			"APP_TIMEOUT": gofsm.TransitionAction{
				To: "CLOSED",
			},
		},
		"CLOSE_WAIT": {
			"APP_CLOSE": gofsm.TransitionAction{
				To: "LAST_ACK",
			},
		},
		"LAST_ACK": {
			"RCV_ACK": gofsm.TransitionAction{
				To: "CLOSED",
			},
		},
	}
}

func TraverseTCPStates(events []string) string {
	states := buildStateDiagram()

	sm := gofsm.NewStateMachine("CLOSED", states)

	var err error
	for _, event := range events {
		err = sm.Transition(context.Background(), event, nil)
		if err != nil {
			break
		}
	}

	if err == nil {
		return sm.GetStatus()
	}

	return "ERROR"
}
