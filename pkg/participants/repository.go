package participants

import "context"

type Repository interface {
	FindByEmailAndEventName(ctx context.Context, email, event string) (*Participant, error)
	SetPhoneNameReg(ctx context.Context, email, phone, name, reg string) error
	New(ctx context.Context, participant Participant) (ID interface{}, err error)
	FindByEventName(ctx context.Context, eventName string) ([]Participant, error)
}
