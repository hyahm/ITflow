package gadb

import (
	"github.com/go-redis/redis"
	"github.com/hyahm/goconfig"
)

//var client *redis.Client

type ConnectRedis struct {
	Rediscli *redis.Client
}

type Rconfig struct {
	RedisHost string `json:"redishost"`
	RedisDb   int    `json:"redisdb"`
	RedisPwd  string `json:"redispwd"`
}

func NewRedis() *Rconfig {
	return &Rconfig{
		RedisDb:   goconfig.ReadInt("redisdb"),
		RedisHost: goconfig.ReadString("redishost"),
		RedisPwd:  goconfig.ReadString("redispwd"),
	}
}

func (rd *Rconfig) Connect() (*redis.Client, error) {

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     rd.RedisHost,
		Password: rd.RedisPwd, // no password set
		DB:       rd.RedisDb,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return RedisClient, nil
}
