package participants

import (
	"context"

	pkg "github.com/GDGVIT/dsc-events-registration/pkg"
)

type Service interface {
	Save(ctx context.Context, participant *Participant) (ID interface{}, err error)
	CountParticipantsByEvent(ctx context.Context, eventName string) (*int, error)
	CountParticipantsByEvents(ctx context.Context) (interface{}, error)
	// ShowParticipantsByEvent(ctx context.Context, eventName string) ([]Participant, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Save(ctx context.Context, participant *Participant) (ID interface{}, err error) {

	p, err := s.repo.FindByEmailAndEventName(ctx, participant.EmailAddress, participant.EventName)

	if err != nil {
		return "", err
	}

	if p != nil {
		return "", pkg.ErrExists
	}
	id, err := s.repo.New(ctx, *participant)
	if err != nil {
		return "", err
	}
	return id, nil

}

func (s *service) CountParticipantsByEvent(ctx context.Context, eventName string) (*int, error) {
	p, err := s.repo.FindByEventName(ctx, eventName)

	if err != nil {
		return nil, err
	}
	count := len(p)
	return &count, nil
}

func (s *service) CountParticipantsByEvents(ctx context.Context) (interface{}, error) {
	data, err := s.repo.GroupByEventName(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
