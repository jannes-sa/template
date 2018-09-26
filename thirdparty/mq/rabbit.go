package mq

import (
	"os"
	"template/helper"
	"template/helper/timetn"

	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
)

// SendCommonMQ ...
func SendCommonMQ(post []byte, channel string, mq string, exchange string,
	eventName string) error {

	if os.Getenv("GOENV") != "devci" && os.Getenv("GOENV") != "local" &&
		os.Getenv("GOENV") != "mac" {

		beego.Debug(string(post))
		conn, err := ConnectMq(mq)
		if err != nil {
			return err
		}
		helper.CheckErr("Failed to connect to RabbitMQ @SendGLMQ", err)
		defer func() {
			err = conn.Close()
			helper.CheckErr("Error Close Connection @SendGLMQ", err)
		}()

		ch, err := conn.Channel()
		helper.CheckErr("Failed to open a channel @SendGLMQ", err)
		if err != nil {
			return err
		}
		defer func() {
			err = ch.Close()
			helper.CheckErr("Error Close Channel @SendGLMQ", err)
		}()

		_, err = ch.QueueDeclare(
			channel, // name
			true,    // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		helper.CheckErr("Failed to declare a queue @SendGLMQ", err)
		if err != nil {
			return err
		}
		// beego.Debug(q)
		err = ch.Publish(
			exchange,  // exchange
			eventName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				Timestamp:    timetn.Now(),
				ContentType:  "application/json",
				Body:         post,
			},
		)
		helper.CheckErr("Failed to publish a message @SendGLMQ", err)
		if err != nil {
			return err
		}
	}

	return nil

}

// ConnectMq ...
func ConnectMq(mq string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(mq)
	helper.CheckErr("Failed to connect to RabbitMQ", err)

	return conn, err
}
