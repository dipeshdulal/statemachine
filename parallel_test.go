package statemachine_test

import (
	"errors"
	"testing"

	"github.com/dipeshdulal/statemachine"
	"github.com/stretchr/testify/assert"
)

func TestParallelToggle(t *testing.T) {
	times := 0
	machineOne := statemachine.Machine{
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
	machineTwo := statemachine.Machine{
		ID:      "machine-2",
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

	parallel := statemachine.ParallelMachine{
		Machines: statemachine.Machines{
			"machine-1": &machineOne,
			"machine-2": &machineTwo,
		},
		Subscribers: statemachine.ParallelSubscribers{
			func(c, n statemachine.ParallelState) { times++ },
			func(c, n statemachine.ParallelState) { times++ },
			func(c, n statemachine.ParallelState) { times++ },
		},
	}

	next, err := parallel.Transition("machine-1.TOGGLE")
	assert.Equal(t, statemachine.ParallelState{"machine-1": "off", "machine-2": "on"}, next, "Transition should occur on toggle.")
	assert.Equal(t, nil, err, "Error should not occurr in correct transition")
	assert.Equal(t, 3, times, "Subscribers should be called on transition")

	next, err = parallel.Transition("machine-one")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("event format doesn't match"), err)
		assert.Equal(t, next, statemachine.ParallelState{"machine-1": "off", "machine-2": "on"}, "Transition should not occur on error.")
	} else {
		t.Error("error should occurr when key format doesn't match")
	}

	next, err = parallel.Transition("machine-one.TOGGLE")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("machine key doesnot match"), err)
		assert.Equal(t, next, statemachine.ParallelState{"machine-1": "off", "machine-2": "on"}, "Transition should not occur on error.")
	} else {
		t.Error("error should occurr when machine key doesnot exist")
	}

	assert.Equal(t, 3, times, "Subscribers shouldnot be called on error on transition")
}
