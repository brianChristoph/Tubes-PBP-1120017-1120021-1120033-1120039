package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func SetRedis(c *gin.Context, name string) {
	// Initialized Redis Client
	rdb := newRedisClient(LoadEnv("REDIS_HOST"), LoadEnv("REDIS_PASS"))
	data := name
	expirationTime := time.Duration(1) * time.Hour

	// Store data to redis
	op := rdb.Set(context.Background(), LoadEnv("REDIS_KEY"), data, expirationTime)
	if err := op.Err(); err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{
			"status":  http.StatusServiceUnavailable,
			"message": "Unable to Set data",
			"data":    err,
		})
		return
	}
}

func GetRedis(c *gin.Context) string {
	rdb := newRedisClient(LoadEnv("REDIS_HOST"), LoadEnv("REDIS_PASS"))

	// Get data from redis
	op := rdb.Get(context.Background(), LoadEnv("REDIS_KEY"))
	if err := op.Err(); err != nil {
		c.IndentedJSON(http.StatusGone, gin.H{
			"message": "Redis Expired",
			"status":  http.StatusGone,
		})
		return ""
	}
	res, err := op.Result()
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{
			"message": "Redis Nil Value",
			"status":  http.StatusNoContent,
		})
		return ""
	}
	return res
}

func DeleteRedis(c *gin.Context) {
	rdb := newRedisClient(LoadEnv("REDIS_HOST"), LoadEnv("REDIS_PASS"))
	err := rdb.Del(context.Background(), LoadEnv("REDIS_KEY")).Err()
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{
			"message": "Redis Nil Value",
			"status":  http.StatusNoContent,
		})
		return
	}
}
