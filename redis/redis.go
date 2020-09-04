package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// Rdb redis
var Rdb *redis.Client

// ConnectRedis 链接redis
func ConnectRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	Rdb = rdb
}

// SaveToken 保存userToken
func SaveToken(uID uint, token string) error {
	return saveData("userInfo_id"+string(uID), token)
}

// GetToken 获取token
func GetToken(uID uint) (string, error) {
	return getData("userInfo_id" + string(uID))
}

// SaveData 保存数据
func saveData(key string, value string) error {
	if Rdb == nil {
		ConnectRedis()
	}
	err := Rdb.Set(ctx, key, value, 0).Err()
	return err
}

//GetData 获取数据
func getData(key string) (string, error) {
	val, err := Rdb.Get(ctx, key).Result()
	return val, err
}
