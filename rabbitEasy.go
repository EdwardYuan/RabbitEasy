package rabbiteasy

import (
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func init() {
	loadConfig("./")
}

type RabbitEasy struct {
	Addr  string
	Queue *amqp.Queue
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", err, msg)
	}
}

func loadConfig(path string) {
	viper.SetConfigName("rabbitConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	failOnError(err, "Read RabbitEasy Config error")
}

func RabbitSubscribeTopic(topic string) {

}

func RabbitSubscribeQueue(name string) {

}

func RabbitPublishTopic(topic string) {

}

func RabbitPublishQueue(name string) {

}
