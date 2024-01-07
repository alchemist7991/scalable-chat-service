package redisHelper

import (
	"fmt"
	"log"
	"github.com/alchemist7991/scalable-chat-service/constant"
	"github.com/go-redis/redis"
)

// create new connection to redis
func GetRedisClient() *redis.Client {
	redisUrl := fmt.Sprintf("%s:%s", constant.REDIS_HOST, constant.REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr: redisUrl,
		Password: "",
		DB: int(constant.REDIS_CLUSTER),
	})
	err := IsSuccessfullyConnected(client)
	if err != nil {
		log.Panic(err)
	}
	return client
}

func IsSuccessfullyConnected(client *redis.Client) error{
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("Failed to make successful connction")
		return err
	}
	log.Println("Successfully connected to redis ", pong)
	return nil
}
