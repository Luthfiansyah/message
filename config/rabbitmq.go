package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func RabbitMqConnection() (*amqp.Connection, error){
	conn, err := amqp.Dial("amqp://"+ viper.GetString("RABBITMQ_USERNAME")+":"+viper.GetString("RABBITMQ_PASSWORD")+"@"+viper.GetString("RABBITMQ_HOST")+":"+viper.GetString("RABBITMQ_PORT")+"/")
	if err != nil {
		fmt.Println(err.Error())
	}
	//defer conn.Close()
	return conn, err
}