package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rs/xid"
	"log"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unbale to connect to Redis", err)
	}

	subject := "transactions"
	consumersGroup := "transactions-report"

	err = redisClient.XGroupCreate(subject, consumersGroup, "0").Err()
	if err != nil {
		log.Println(err)
	}

	uniqueID := xid.New().String()

	for {

		entries, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
			Group:    consumersGroup,
			Consumer: uniqueID,
			Streams:  []string{subject, ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values

			transactionId := fmt.Sprintf("%v", values["transactionId"])
			phoneNumber := fmt.Sprintf("%v", values["phoneNumber"])
			amount := fmt.Sprintf("%v", values["amount"])
			time := fmt.Sprintf("%v", values["time"])

			err := handleNewTicket(transactionId, phoneNumber, amount, time)
			if err != nil {
				log.Fatal(err)
			}
			redisClient.XAck(subject, consumersGroup, messageID)
		}
	}

}

func handleNewTicket(transactionId string, phoneNumber string, amount string, time string) error {
	log.Printf("Handling new transaction id : %s, phone_number: %s, amount :%s, time:%s\n", transactionId, phoneNumber, amount, time)
	return nil
}
