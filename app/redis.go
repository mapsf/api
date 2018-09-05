package app

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
	"os"
)

var (
	Redis *redis.Client
)

func InitRedis() {

	log.Printf("STARTING REDIS CONNECTION...")

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	log.Println("PING REDIS SERVER...")
	pong, err := Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println("REDIS RESPOND WITH " + pong)
}
