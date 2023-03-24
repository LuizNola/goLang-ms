package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	kafka2 "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(msg)
		go kafka2.Produce(msg)
	}

}
