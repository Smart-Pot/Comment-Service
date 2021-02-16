package transport

import (
	"commentservice/service"
	"context"

	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func RunDeleteUserCommentsConsumer(consumer amqp.Consumer, service service.Service) {
	for {
		userID := string(consumer.Consume())
		service.DeleteUsersComments(context.TODO(), userID)

	}
}
