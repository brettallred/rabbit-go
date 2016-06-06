package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection

func Handlers() map[string]func(b []byte) bool {
	return handlers
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
