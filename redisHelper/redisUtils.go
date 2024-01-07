package redisHelper

import (
	"encoding/json"
	"log"
	"time"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// Json Marshalling only works with Upper Case Keys
type Message struct {
	UserId string	`json:"userId"`
	Message string `json:"message"`
	CreatedAt string `json:"created_at"`
	RemoteAddr string `json:"remote_addr"`
}

func GetMessage(message string, userId string, remoteAddr string) *Message {
	return &Message{
		UserId: userId,
		Message: message,
		CreatedAt: time.Now().String(),
		RemoteAddr: remoteAddr,
	}
}

func StoreMessage(message string, userId string, remoteAddr string) {
	messageStruct := GetMessage(message, userId, remoteAddr)
	msgJson, err := json.Marshal(messageStruct)
	if err != nil {
		log.Panic("Unable to marshal message, Err: ", err)
	}
	key := messageStruct.UserId
	status := redisClient.Set(key, msgJson, 0)
	if status.Err() != nil {
		log.Panic("Error in pushing message to redis, Err: ", status.Err())
	}
	RetriveMessages(key)
}

func RetriveMessages(key string) {
	messages, err := redisClient.Get(key).Result()
	if err != nil {
		log.Println("Error while retrieving data from DB")
	}
	log.Println(messages)
}

func SetClientInstance() {
	if redisClient == nil {
		redisClient = GetRedisClient()
	}
}
