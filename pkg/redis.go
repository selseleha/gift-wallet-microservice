package pkg

import (
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"task/pkg/models"
	"time"
)

type RedisService struct {
	client *redis.Client
}

type RedisOption struct {
	RedisURL string
}

func NewRedis(input *RedisOption) *RedisService {
	if input == nil {
		log.Fatal("input is required")
	}
	client := redis.NewClient(&redis.Options{
		Addr: input.RedisURL,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return &RedisService{
		client: client,
	}
}

func (rs *RedisService) Set(key string, value interface{}) error {
	exp := time.Duration(600 * time.Second)
	return rs.client.Set(key, value, exp).Err()
}

func (rs *RedisService) Get(key string) (interface{}, error) {
	val, err := rs.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (rs *RedisService) PublishTransaction(transaction models.Transaction) error {

	err := rs.client.XAdd(&redis.XAddArgs{
		Stream:       "transactions",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values: map[string]interface{}{
			"transactionId": strconv.Itoa(int(transaction.Id)),
			"phoneNumber":   transaction.PhoneNumber,
			"amount":        transaction.Amount,
			"time":          transaction.CreatedAt.String(),
		},
	}).Err()

	return err
}
