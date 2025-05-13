package gofsm

import (
	"context"
	"errors"
)

var (
	ErrTransitionNotFound = errors.New("transition not found")
)

type TransitionAction struct {
	To     string
	Action func(ctx context.Context, input any) error
}

type Transitions map[string]TransitionAction

type StateMachine struct {
	CurrentState string
	States       map[string]Transitions
}

func NewStateMachine(currentState string, states map[string]Transitions) *StateMachine {
	return &StateMachine{
		CurrentState: currentState,
		States:       states,
	}
}

func (ss *StateMachine) Transition(ctx context.Context, event string, input any) error {
	transitionAction, ok := ss.States[ss.CurrentState][event]
	if !ok {
		return ErrTransitionNotFound
	}

	ss.CurrentState = transitionAction.To
	if transitionAction.Action != nil {
		return transitionAction.Action(ctx, input)
	}

	return nil
}

func (ss *StateMachine) GetStatus() string {
	return ss.CurrentState
}
