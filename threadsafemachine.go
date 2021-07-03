package statemachine

import (
	"errors"
	"strings"
	"sync"
)

// ParallelMachine to start a parallel state machine
type ThreadsafeMachine struct {
	Machines    Machines
	Subscribers []func(curr, next ParallelState)
	mu          sync.RWMutex
}

// Current returns current state of parallel machines, rlock
func (m *ThreadsafeMachine) Current() ParallelState {
	m.mu.RLock()
	defer m.mu.RUnlock()
	currentStateMap := make(ParallelState)
	for machine := range (*m).Machines {
		currentStateMap[machine] = (*m).Machines[machine].Current()
	}
	return currentStateMap
}

// current returns current state of parallel machines, no lock
func (m *ThreadsafeMachine) current() ParallelState {
	currentStateMap := make(ParallelState)
	for machine := range (*m).Machines {
		currentStateMap[machine] = (*m).Machines[machine].Current()
	}
	return currentStateMap
}

// Transition transitions to next state. rwlock
// event format is
//  m.Transition("machinekey.eventName")
func (m *ThreadsafeMachine) Transition(event string) (ParallelState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s := strings.Split(event, ".")
	if len(s) != 2 {
		return m.current(), errors.New("event format doesn't match")
	}

	if _, ok := (*m).Machines[s[0]]; ok {
		current := m.current()
		(*m).Machines[s[0]].Transition(s[1])
		for _, funct := range (*m).Subscribers {
			funct(current, m.current())
		}
		return m.current(), nil
	}

	return m.current(), errors.New("machine key doesnot match")
}
