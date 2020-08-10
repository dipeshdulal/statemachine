## Simple Finite State Machines in GO

Learn more about finite state machines [here](https://xstate.js.org/docs/about/concepts.html#finite-state-machines)

Inspired from [@davidkpiano](https://github.com/davidkpiano) X-State library. (minimal finite state machine library in javascript) `@xstate/fsm`

![basic-toggle](basic_toggle.png)
Figure from xstate.js.org/viz/


```go
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
// output will contain state of the machine
// after sending TOGGLE event
```

##### Checklist

- [x] Implement basic finite state machine with transition
- [x] Add condition `cond` function calling to get the output.
- [x] Add multiple actions and call the function on state change.
- [ ] Add state change listeners. 
    - May be use channels and concurrency concepts of golang [publisher-subscriber pattern]
    - Also implement basic callback function as well. 

