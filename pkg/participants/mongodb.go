package participants

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	Collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) Repository {
	return &repo{Collection: collection}
}

func (r *repo) FindByEmailAndEventName(ctx context.Context, email, event string) (*Participant, error) {
	filter := bson.M{
		"emailAddress": email,
		"eventName":    event,
	}
	participant := &Participant{}
	err := r.Collection.FindOne(ctx, filter).Decode(participant)
	switch err {
	case nil:
		break
	case mongo.ErrNoDocuments:
		return nil, nil
	default:
		return nil, err
	}
	return participant, nil
}

func (r *repo) SetPhoneNameReg(ctx context.Context, email, phone, name, reg string) error {
	return nil
}

func (r *repo) New(ctx context.Context, participant Participant) (ID interface{}, err error) {
	res, err := r.Collection.InsertOne(ctx, participant)
	if err != nil {
		return "", err
	}
	ID = res.InsertedID
	return ID, err
}

func (r *repo) FindByEventName(ctx context.Context, eventName string) ([]Participant, error) {
	filter := bson.M{
		"eventName": eventName,
	}
	cur, err := r.Collection.Find(ctx, filter)

	switch err {
	case nil:
		break
	case mongo.ErrNoDocuments:
		return nil, nil
	default:
		return nil, err
	}
	defer cur.Close(ctx)
	participants := []Participant{}

	for cur.Next(ctx) {
		participant := Participant{}
		cur.Decode(&participant)
		participants = append(participants, participant)
	}

	return participants, nil
}
