package transport

import (
	"commentservice/service"
	"context"

	"github.com/Smart-Pot/pkg/adapter/amqp"
)

func RunDeletePostCommentsConsumer(consumer amqp.Consumer, service service.Service) {
	for {
		postID := string(consumer.Consume())
		service.DeleteMany(context.TODO(), postID)
	}
}
