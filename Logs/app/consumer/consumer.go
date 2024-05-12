package consumer

import (
	"fmt"

	"github.com/IBM/sarama"
)

func internalConsumer() {
	topic := "internal-logs"
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	brokers := []string{"kafka:9092"}
	consumer, err := sarama.NewConsumer(brokers, consumerConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to start consumer: %s", err))
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(fmt.Sprintf("Failed to start consumer partition: %s", err))
	}
	internalReader(topic, partitionConsumer)
}

func clientActConsumer() {
	topic := "clientActivities-logs"
	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	brokers := []string{"kafka:9092"}
	consumer, err := sarama.NewConsumer(brokers, consumerConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to start consumer: %s", err))
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(fmt.Sprintf("Failed to start consumer partition: %s", err))
	}
	clientActReader(topic, partitionConsumer)
}

func Start() {
	internalConsumer()
	clientActConsumer()
}
