package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic     = "book"
)

/*  */

func Publish(message []byte, producer sarama.SyncProducer) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
