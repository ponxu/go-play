package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	pool := newPool(":6379", "")
	defer pool.Close()

	fmt.Println(pool.Get().Do("SET", "name", "xwz"))
	fmt.Println(pool.Get().Do("DBSIZE"))
	fmt.Println(pool.Get().Do("KEYS", "*"))
	fmt.Println(pool.Get().Do("GET", "name"))
}

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
