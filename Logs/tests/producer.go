package tests

import (
	"Logs/app/consumer"
	"Logs/types"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func producer() {
	topic := "Welcome-called"
	//topic := "Welcome-enter-data"
	//topic := "Welcome-exit-data"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true
	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		log.Fatalln("Failed to start producer:", err)
	}
	defer producer.Close()

	val := &types.Called{Timestamp: time.Now(),
		LogId:   1,
		Message: "producer()",
		Data:    "id = 2, name = 'Bogdan'"}

	jd, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jd),
	}
	select {
	case producer.Input() <- msg:
		log.Print("Message has been sent")
	case <-producer.Errors():
		log.Println("Failed to send message")
	}
}

func testStorage() {
	if consumer.KafkaStorage.Message != "producer()" {
		panic(fmt.Sprintf(`consumer.KafkaStorage.Message != "producer()" because consumer.KafkaStorage.Message = "%s"`, consumer.KafkaStorage.Message))
	}
	if consumer.KafkaStorage.Data != "id = 2, name = 'Bogdan'" {
		panic(fmt.Sprintf(`consumer.KafkaStorage.Data != "id = 2, name = 'Bogdan'" because consumer.KafkaStorage.Data = "%s"`, consumer.KafkaStorage.Data))
	}
	if consumer.KafkaStorage.LogId != 1 {
		panic(fmt.Sprintf("consumer.KafkaStorage.LogId != 1 because consumer.KafkaStorage.LogId = %d", consumer.KafkaStorage.LogId))
	}
	log.Print("EVERYTHING WAS OK. TEST WAS COMPLETED")
}

func StartTest() {
	producer()
	testStorage()
}
