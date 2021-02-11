package endpoints

import "github.com/Smart-Pot/pkg/adapter/amqp"

//ERROR HANDLE
func MakeDeletePostCommentsConsumer() (amqp.Consumer, error) {
	return amqp.MakeConsumer("comment1", "DeletePostComments")
}

func MakeDeleteUserCommentsConsumer() (amqp.Consumer, error) {
	return amqp.MakeConsumer("comment2", "DeleteUserComments")
}
