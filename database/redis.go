package database

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var client *redis.Client

func RedisOpen() *redis.Client {

	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetInt("REDIS_PORT")
	// redisUsername := viper.GetString("REDIS_USERNAME")
	redisPassword := viper.GetString("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     "" + redisHost + ":" + strconv.Itoa(redisPort) + "",
		Password: redisPassword, // no password set
		DB:       0,             // use default DB
	})

	return client
}

func SetRedis(Key string, Value string) error {
	client := RedisOpen()

	// SET KEY
	err := client.Set(Key, Value, 3600000*time.Second).Err()
	//err := client.Set(Key, Value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetRedis(Key string) string {
	client := RedisOpen()

	// GET KEY
	data, err := client.Get(Key).Result()

	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}

func DelRedis(Key string) error {
	client := RedisOpen()

	// DEL KEY
	err := client.Del(Key).Err()
	if err != nil {
		return err
	}
	return nil
}
