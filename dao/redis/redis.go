package redis

import (
	"log"
	"time"

	"cytus2.rocks/cics2/configloader"
	"github.com/go-redis/redis"
)

var clients map[string]*redis.Client

func init() {
	clients = make(map[string]*redis.Client)
	for name, addr := range configloader.GetByL1("redis") {
		clients[name] = redis.NewClient(&redis.Options{
			Addr:         addr,
			DB:           0,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
			// retry failed command(watch command need this, since i don't provide CICS level retry)
			MaxRetries: 10, // retry at most 11 times
		})
	}
}

// MustGetClient always return a non-nil Client,
// if it doesn't exist, log and panic
func MustGetClient(name string) *redis.Client {
	client := GetClient(name)
	if client == nil {
		log.Panicf("get [%s] redis client fail\n", name)
	}
	return client
}

// GetClient given a name, return a redis client
func GetClient(name string) *redis.Client {
	return clients[name]
}

// MustCmdSuccess makes sure the cmd is success(no error), or it will log and panic
func MustCmdSuccess(cmd redis.Cmder) {
	if err := cmd.Err(); err != nil {
		log.Panicln("redis cmd fail", err)
	}
}
