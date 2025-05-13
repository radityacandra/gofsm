package examples

import (
	"context"
	"time"

	"github.com/radityacandra/gofsm"
)

type Order struct {
	Status    string
	UpdatedAt int64
	UpdatedBy string

	StateMachine *gofsm.StateMachine
}

func NewOrder() *Order {
	sm := gofsm.NewStateMachine("CREATED", map[string]gofsm.Transitions{
		"CREATED": {
			"PAYMENT_CAPTURED": gofsm.TransitionAction{
				To: "PAID",
			},
			"REQUEST_CANCEL": gofsm.TransitionAction{
				To: "CANCELED",
			},
		},
		"PAID": {
			"DELIVERED_TO_CUSTOMER": gofsm.TransitionAction{
				To: "CLOSED",
			},
		},
	})

	return &Order{
		Status:       "CREATED",
		StateMachine: sm,
	}
}

func (o *Order) UpdateStatus(ctx context.Context, status, userId string) error {
	o.UpdatedAt = time.Now().UnixMilli()
	o.UpdatedBy = userId

	err := o.StateMachine.Transition(ctx, status, nil)

	o.Status = status

	return err
}
