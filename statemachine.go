package statemachine

// MachineTransition transition map
type MachineTransition struct {
	Action []func()
	Cond   func()
	To     string
}

// TransitionMap map with transitions
type TransitionMap map[string]MachineTransition

// MachineState is State of machine
type MachineState struct {
	On TransitionMap
}

// StateMap maps state
type StateMap map[string]MachineState

// Machine datatype
type Machine struct {
	ID      string
	Initial string
	current string
	States  StateMap
}

// IMachine machine interface
type IMachine interface {
	Transition() string
	Current() string
}

// Current returns current state
func (m *Machine) Current() string {
	if m.current == "" {
		return m.Initial
	}
	return m.current
}

// Transition transitions to next state
func (m *Machine) Transition(event string) string {
	current := m.Current()
	transitions := m.States[current].On
	for evt := range transitions {
		if evt == event {
			next := transitions[evt].To
			m.current = next
			return next
		}
	}
	return current
}
