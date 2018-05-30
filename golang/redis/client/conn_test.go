package client_test

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
	"github.com/wangming1993/share/golang/redis/client"
)

func TestSomething(t *testing.T) {
	s := miniredis.NewMiniRedis()
	s.RequireAuth(client.REDIS_PASSWD)
	s.DB(client.REDIS_DB)
	err := s.StartAddr(":6379")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// Run your code and see if it behaves.
	// An example using the redigo library from "github.com/garyburd/redigo/redis":
	c, err := redis.Dial("tcp", s.Addr())

	client.Conn = c

	_, err = c.Do("SET", "name", "mike")

	exists := client.NameExists()

	if !exists {
		t.Fail()
	}
}
