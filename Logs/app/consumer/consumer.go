package consumer

import (
	"Logs/types"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

var KafkaStorage types.Called

func read(partcons sarama.PartitionConsumer) {
	for {
		select {
		case msg := <-partcons.Messages():
			fmt.Println("Received message:", string(msg.Value))
			err := json.Unmarshal(msg.Value, &KafkaStorage)
			if err != nil {
				panic(err)
			}
		case err := <-partcons.Errors():
			fmt.Println("Received error:", err)
			panic(fmt.Sprintf("Error while consuming: %s", err))
		}
	}
}

func StartConsumer() {
	topic := "Welcome-called"
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
	//defer partitionConsumer.Close()
	read(partitionConsumer)
}
