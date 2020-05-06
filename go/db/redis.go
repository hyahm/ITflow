package db

import (
	"github.com/go-redis/redis"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/xredis"
)

//var client *redis.Client

//type ConnectRedis struct {
//	Rediscli *redis.Client
//}

//type Rconfig struct {
//	RedisHost string `json:"redishost"`
//	RedisDb   int    `json:"redisdb"`
//	RedisPwd  string `json:"redispwd"`
//}

var RSconn *xredis.TypeString

func InitRedis() {
	rconn, err := xredis.Conn(&redis.Options{
		Addr:     goconfig.ReadString("redis.host", "127.0.0.1:6379"),
		Password: goconfig.ReadString("redis.pwd", ""),
		DB:       goconfig.ReadInt("redis.db", 0),
	})
	if err != nil {
		panic(err)
	}
	RSconn = rconn.NewStr()

}
