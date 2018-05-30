package client

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Redis exposed to access redis
var Redis = &redisClient{}

var (
	REDIS_HOST   = "47.100.114.173:6379"
	REDIS_PASSWD = "Bzyaiabc123"
	REDIS_DB     = 15
)

type redisClient struct {
	connPool *redis.Pool
}

func init() {
	pool, err := initRedisPool(REDIS_HOST, REDIS_PASSWD, REDIS_DB)
	if err != nil {
		panic(err)
	}

	Redis.connPool = pool
}

func initRedisPool(host, pass string, db int) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialTimeout("tcp", host, 1*time.Second, 0, 0)
			if err != nil {
				return nil, err
			}

			if "" != pass {
				if _, err := c.Do("AUTH", pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err = c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	//test Redis connection
	conn := pool.Get()
	defer conn.Close() //close it to pool after tested

	_, err := conn.Do("PING")
	if nil != err {
		return nil, err
	}
	return pool, nil
}

// provide redis function

func (c *redisClient) conn() redis.Conn {
	return c.connPool.Get()
}

func (c *redisClient) Get(key string) (string, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.String(conn.Do("GET", key))
}

func (c *redisClient) Expire(key string, time int64) (bool, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Bool(conn.Do("EXPIRE", key, time))
}

func (c *redisClient) TTL(key string) (int64, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Int64(conn.Do("TTL", key))
}

func (c *redisClient) Set(key string, val string) error {
	conn := c.conn()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)

	return err
}

func (c *redisClient) SetNX(key string, val string, ttl int) (bool, error) {
	conn := c.conn()
	defer conn.Close()

	ok, err := redis.String(conn.Do("SET", key, val, "EX", ttl, "NX"))
	if ok == "" {
		return false, err
	}

	return true, err
}

func (c *redisClient) GetSet(key string, val string) (string, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.String(conn.Do("GETSET", key, val))
}

func (c *redisClient) Incr(key string) (int, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Int(conn.Do("INCR", key))
}

func (c *redisClient) IncrBy(key string, increment int) (int, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Int(conn.Do("INCRBY", key, increment))
}

func (c *redisClient) Exists(key string) (bool, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Bool(conn.Do("EXISTS", key))
}

func (c *redisClient) Delete(key string) (int, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Int(conn.Do("DEL", key))
}

func (c *redisClient) ZAdd(keysAndArgs ...interface{}) (int, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.Int(conn.Do("ZADD", keysAndArgs...))
}

func (c *redisClient) ZRange(key string, start, stop int) (map[string]string, error) {
	conn := c.conn()
	defer conn.Close()

	return redis.StringMap(conn.Do("ZRANGE", key, start, stop, "WITHSCORES"))
}

var Conn redis.Conn

func init() {
	Conn = Redis.connPool.Get()
}

func NameExists() bool {
	v, err := Conn.Do("GET", "name")
	fmt.Println(v, err)
	return err == nil && v != nil
}

// func NameExists() bool {
// 	v, err := Redis.Get("name")

// 	return err == nil && v != ""
// }
