package consumer

import (
	"Logs/app/handler/writer"
	"Logs/apptype"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

func internalReader(top string, partcons sarama.PartitionConsumer) {
	intrnl := new(apptype.Internal)
	for {
		select {
		case msg := <-partcons.Messages():
			err := json.Unmarshal(msg.Value, &intrnl)
			if err != nil {
				log.Print(err)
			}
			writer.Internal(intrnl)
		case err := <-partcons.Errors():
			log.Printf("Error while consuming (topic: %s): %s", top, err)
		}
	}
}

func clientActReader(top string, partcons sarama.PartitionConsumer) {
	apptype.Db = apptype.ConnectToDatabase()
	opnd := new(apptype.ClientAct)
	for {
		select {
		case msg := <-partcons.Messages():
			err := json.Unmarshal(msg.Value, &opnd)
			if err != nil {
				log.Print(err)
			}
			writer.ClientActivities(opnd)
		case err := <-partcons.Errors():
			log.Printf("Error while consuming (topic: %s): %s", top, err)
		}
	}
}
