package main

import (
	"github.com/go-redis/redis"
)

type app struct {
	r *redis.Client
}

func (a *app) init() {
	a.r = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func (a *app) set(key string, value string) {
	a.r.Set(key, value, 0)
}

func (a *app) get(key string) string {
	return a.r.Get(key).Val()
}

func (a *app) del(key string) {
	a.r.Del(key)
}

func main() {
	var a app
	a.init()
	a.set("foo", "bar")
	defer println("Cleaned up")
	defer a.del("foo")
	println(a.get("foo"))
}
