package statemachine

import (
	"errors"
	"strings"
)

// Machines bundle up multiple state machine definitions
type Machines map[string]IMachine

// ParallelState describes parallel state machines state
type ParallelState map[string]string

// ParallelSubscribers type declaration for parallel machine subscribers
type ParallelSubscribers []func(ParallelState, ParallelState)

// ParallelMachine to start a parallel state machine
type ParallelMachine struct {
	Machines    Machines
	Subscribers []func(curr, next ParallelState)
}

// Current returns current state of parallel machines
func (m *ParallelMachine) Current() ParallelState {
	currentStateMap := make(ParallelState)
	for machine := range (*m).Machines {
		currentStateMap[machine] = (*m).Machines[machine].Current()
	}
	return currentStateMap
}

// Transition transitions to next state.
// event format is
//  m.Transition("machinekey.eventName")
func (m *ParallelMachine) Transition(event string) (ParallelState, error) {
	s := strings.Split(event, ".")
	if len(s) != 2 {
		return m.Current(), errors.New("event format doesn't match")
	}

	if _, ok := (*m).Machines[s[0]]; ok {
		current := m.Current()
		(*m).Machines[s[0]].Transition(s[1])
		for _, funct := range (*m).Subscribers {
			funct(current, m.Current())
		}
		return m.Current(), nil
	}

	return m.Current(), errors.New("machine key doesnot match")
}
