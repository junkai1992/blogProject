package common

// redis连接
import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RDB *redis.Client

func InitRedisClient() (err error) {
	host := viper.GetString("redissource.host")
	port := viper.GetString("redissource.port")
	addr := fmt.Sprintf("%s:%s", host, port)
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedisDB() *redis.Client {
	return RDB
}
