package rabbit

import (
	"faceit-test/internal/config"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
)

func NewPublisher(cfg *config.Rabbit) (*rabbitmq.Publisher, error) {
	publisher, err := rabbitmq.NewPublisher(url(cfg), rabbitmq.Config{})
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

func url(cfg *config.Rabbit) string {
	return fmt.Sprintf("amqp://%s:%s@%s", cfg.Username, cfg.Password, cfg.Host)
}
