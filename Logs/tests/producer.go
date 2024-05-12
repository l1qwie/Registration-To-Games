package tests

import (
	"Logs/tests/database"
	"Logs/types"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func createInterProd() {
	topic := "internal-logs"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true
	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to start producer: %s", err))
	}
	defer producer.Close()

	timestampStr := "2024-05-12T10:48:34.670096883Z"
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге времени: %s", err))
	}

	val := &types.Internal{
		Com:     types.Common{Timestamp: timestamp, LogId: 1},
		Message: "function Receiving in Media microservice was called",
		Data:    "req *types.Request was given"}

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

func testInternal() {
	createInterProd()
	time.Sleep(time.Second * 4)
	testInterData()
}

func createClProd() {
	topic := "clientActivities-logs"
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true
	brokers := []string{"kafka:9092"}
	producer, err := sarama.NewAsyncProducer(brokers, producerConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to start producer: %s", err))
	}
	defer producer.Close()

	timestampStr := "2024-05-12T10:48:34.670096883Z"
	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	if err != nil {
		panic(fmt.Sprintf("Ошибка при парсинге времени: %s", err))
	}

	val := &types.ClientAct{
		Com:     types.Common{Timestamp: timestamp, LogId: 1},
		UserId:  129839,
		Message: "ASL:KDKL:ASDLK:"}

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

func testClientAct() {
	createClProd()
	time.Sleep(time.Second)
	testClData()
	database.TestDatabase()
}

func StartTest() {
	testInternal()
	//testClientAct()
	//time.Sleep(time.Second * 3)
	//testStorage()
}
