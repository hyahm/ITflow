package db

import (
	"github.com/go-redis/redis/v7"
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
		Addr:     goconfig.ReadString("redis.host"),
		Password: goconfig.ReadString("redis.pwd"),
		DB:       goconfig.ReadInt("redis.db"),
	})
	if err != nil {
		panic(err)
	}
	RSconn = rconn.NewStr()

}

//
//func NewRedis() (*redis.Client, error) {
//	conf := &redis.Options{
//			Addr: fmt.Sprintf("%s:%d", goconfig.ReadString("redishost")),
//			DB: goconfig.ReadInt("redisdb"),
//			Password: goconfig.ReadString("redispwd"),
//		}
//	client := redis.NewClient(conf)
//	if err := client.Ping().Err(); err != nil {
//		return nil, err
//	}
//	return client, nil
//}

//func (rd *Rconfig) Connect() (*redis.Client, error) {
//
//	RedisClient := redis.NewClient(&redis.Options{
//		Addr:     rd.RedisHost,
//		Password: rd.RedisPwd, // no password set
//		DB:       rd.RedisDb,  // use default DB
//	})
//
//	_, err := RedisClient.Ping().Result()
//	if err != nil {
//		return nil, err
//	}
//	return RedisClient, nil
//}
