package cache

import (
    "github.com/go-redis/redis/v8"
	"campusCard/config"
)

var Rdb *redis.Client

func init() {
	// 创建一个Redis客户端实例
    Rdb = redis.NewClient(&redis.Options{
        Addr:     config.RedisAddress, // Redis服务器地址
        Password: config.RedisPassword,               // 设置密码，如果没有密码则留空
        DB:       config.RedisDB,                // 使用默认的数据库
    })

}