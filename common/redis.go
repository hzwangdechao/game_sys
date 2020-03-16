package common

import (
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var Client *redis.Client

func InitClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(`redis.addr`),
		Password: viper.GetString(`redis.password`),
		DB:       viper.GetInt(`redis.db`),
	})
	err := client.Ping().Err()
	if err != nil{
		panic(err)
	}

	Client = client
	return Client
}

func GetClient() *redis.Client {
	return Client

}
