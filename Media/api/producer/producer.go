package producer

import (
	"Media/apptype"
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func send(jd []byte, topic string, producer sarama.AsyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jd),
	}
	select {
	case producer.Input() <- msg:
		log.Print("Message has been sent")
	case <-producer.Errors():
		log.Print("Failed to send message")
	}
}

func ActLogs(text string, userId int) {
	topic := "clientActivities-logs"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true
	brokers := []string{"logs-kafka-1:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		log.Printf("Failed to start producer: %s", err)
	} else {
		defer producer.Close()

		val := &apptype.ClientAct{
			Timestamp: time.Now(),
			UserId:    userId,
			Action:    "Registration-To-Games",
			Message:   text}
		jd, err := json.Marshal(val)
		if err != nil {
			log.Printf("KAFKA ERROR ActLogs(): %s", err)
		}
		send(jd, topic, producer)
	}
}

func InterLogs(text, data string) {
	topic := "internal-logs"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true
	brokers := []string{"logs-kafka-1:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		log.Printf("Failed to start producer: %s", err)
	} else {
		defer producer.Close()
		val := &apptype.Internal{
			Timestamp: time.Now(),
			Message:   text,
			Data:      data}

		jd, err := json.Marshal(val)
		if err != nil {
			log.Printf("KAFKA ERROR InterLogs(): %s", err)
		}
		send(jd, topic, producer)
	}
}
