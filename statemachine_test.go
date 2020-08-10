package statemachine_test

import (
	"testing"

	"github.com/dipeshdulal/statemachine"
	"github.com/stretchr/testify/assert"
)

func TestMachineStructure(t *testing.T) {
	machine := statemachine.Machine{
		ID:      "machine-1",
		Initial: "on",
		States: statemachine.StateMap{
			"on": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"TOGGLE": statemachine.MachineTransition{
						To: "off",
					},
				},
			},
			"off": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"TOGGLE": statemachine.MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	output := machine.Transition("TOGGLE")
	assert.Equal(t, output, "off", "Transition should occur on toggle.")
	output = machine.Transition("TOGGLE")
	assert.Equal(t, output, "on", "Transition should occurr on toggle.")
}

func TestMachineCondition(t *testing.T) {
	machine := statemachine.Machine{
		ID:      "machine-1",
		Initial: "on",
		States: statemachine.StateMap{
			"on": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"TOGGLE": statemachine.MachineTransition{
						To: "off",
						Cond: func(curr string) bool {
							return curr == ""
						},
					},
				},
			},
			"off": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"TOGGLE": statemachine.MachineTransition{
						To: "on",
					},
				},
			},
		},
	}
	output := machine.Transition("TOGGLE")
	assert.Equal(t, output, "on", "Transition should not occur on toggle. due to condition")
}
