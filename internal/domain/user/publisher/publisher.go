package publisher

import (
	"context"
	"encoding/json"
	"faceit-test/internal/domain/user/model"
	"github.com/pkg/errors"
	"github.com/wagslane/go-rabbitmq"
)

type User interface {
	Notify(context.Context, model.User, model.UserModificationAction) error
}

type publisherService struct {
	publisher *rabbitmq.Publisher
}

func NewUser(publisher *rabbitmq.Publisher) User {
	return &publisherService{publisher: publisher}
}

func (s *publisherService) Notify(_ context.Context, user model.User, action model.UserModificationAction) error {
	parsedUser, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "failed to parse user to json")
	}
	return s.publisher.Publish(parsedUser, nil, addActionHeader(string(action)))
}

func addActionHeader(action string) func(options *rabbitmq.PublishOptions) {
	return func(options *rabbitmq.PublishOptions) {
		options.Headers = map[string]interface{}{"action": action}
	}
}
