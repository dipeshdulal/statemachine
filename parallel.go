package statemachine

import (
	"errors"
	"strings"
)

// ParallelMachine to start a parallel state machine
type ParallelMachine map[string]IMachine

// Current returns current state of parallel machines
func (m *ParallelMachine) Current() map[string]string {
	currentStateMap := make(map[string]string)
	for machine := range *m {
		currentStateMap[machine] = (*m)[machine].Current()
	}
	return currentStateMap
}

// Transition transitions to next state.
// event format is
//  m.Transition("machinekey.eventName")
func (m *ParallelMachine) Transition(event string) (map[string]string, error) {
	s := strings.Split(event, ".")
	if len(s) != 2 {
		return m.Current(), errors.New("event format doesn't match")
	}

	if _, ok := (*m)[s[0]]; ok {
		(*m)[s[0]].Transition(s[1])
		return m.Current(), nil
	}

	return m.Current(), errors.New("machine key doesnot match")
}
