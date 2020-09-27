package lib

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	defaultRedisClu *redis.ClusterClient
)

type RedisService interface {
	ServiceName() string
	Start()
	Stop()
	Exec()
}

func InitRedis() {
	defaultRedisClu = NewRedisClusterCli(GetDefaultConfRedis())
}

func NewRedisClusterCli(conf *RedisConf) *redis.ClusterClient {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    conf.ProxyList,
		Password: "1234",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("new redis cluster" + err.Error())
		return nil
	}
	return rdb
}

func DefaultRedisCluster() *redis.ClusterClient {
	return defaultRedisClu
}
