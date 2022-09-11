package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

//ConnectDB connects to the database
func ConnectDB() *sqlx.DB {
	return sqlx.MustConnect("postgres", os.Getenv("DB_CONNECTION"))
}

func ConnectCacheDB() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println("🔴 Redis ping: ", pong, err)

	return rdb
}
