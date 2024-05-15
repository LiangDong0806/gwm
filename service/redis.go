package service

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
	"zg5/Homework01/server/global"
)

var RedisDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	address := global.RpcALLConf.Redis.Host
	ip := global.RpcALLConf.Redis.Host
	RedisDB = redis.NewClient(&redis.Options{
		Addr: address + ":" + ip,
		DB:   5,
	})
	fmt.Println(address+":"+ip, "...................")
}

func RedisDBHSet(key string, values interface{}) error {
	err = RedisDB.HSet(ctx, key, values, time.Minute*90).Err()
	return err
}
