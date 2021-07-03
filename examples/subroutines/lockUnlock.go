package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/dipeshdulal/statemachine"
)

func main() {
	machine := statemachine.Machine{
		ID:      "lock-unlock",
		Initial: "locked",
		States: statemachine.StateMap{
			"locked": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"COIN": statemachine.MachineTransition{
						To: "unlocked",
					},
					"PUSH": statemachine.MachineTransition{
						To: "locked",
					},
				},
			},
			"unlocked": statemachine.MachineState{
				On: statemachine.TransitionMap{
					"COIN": statemachine.MachineTransition{
						To: "unlocked",
					},
					"PUSH": statemachine.MachineTransition{
						To: "locked",
					},
				},
			},
		},
	}

	threadSafe := statemachine.ThreadsafeMachine{
		Machines: statemachine.Machines{
			"m1": &machine,
		},
		Subscribers: []func(curr statemachine.ParallelState, next statemachine.ParallelState){},
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		ticker := time.Tick(1 * time.Second)
		for {
			select {
			case <-interrupt:
				return
			case <-ticker:
				fmt.Printf("transition: m1.COIN \n")
				fmt.Printf("current: %v \t", threadSafe.Current())
				output, _ := threadSafe.Transition("m1.COIN")
				fmt.Printf("next: %v \n", output)
			}
		}
	}()

	go func() {
		ticker := time.Tick(2 * time.Second)
		for {
			select {
			case <-interrupt:
				return
			case <-ticker:
				fmt.Printf("transition: m1.PUSH \n")
				fmt.Printf("current: %v \t", threadSafe.Current())
				output, _ := threadSafe.Transition("m1.PUSH")
				fmt.Printf("next: %v \n", output)
			}
		}
	}()

	<-interrupt
}
