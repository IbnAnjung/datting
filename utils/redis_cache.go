package utils

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/IbnAnjung/datting/entity/util_entity"
	"github.com/redis/go-redis/v9"
)

type RedisCaching struct {
	conn     *redis.Client
	key      string
	commands map[int][]interface{}
}

func NewRedisCaching(
	conn *redis.Client,
) RedisCaching {
	return RedisCaching{
		conn:     conn,
		commands: make(map[int][]interface{}),
	}
}

func (r RedisCaching) Set(key string, value interface{}) util_entity.Caching {
	r.key = key
	r.commands[0] = []interface{}{"SET", key, value}

	return r
}

func (r RedisCaching) PushList(key string, value interface{}) util_entity.Caching {
	r.key = key
	r.commands[0] = []interface{}{"LPUSH", key, value}

	return r
}

func (r RedisCaching) Expire(duration time.Duration) util_entity.Caching {
	r.commands[1] = []interface{}{"EXPIRE", r.key, duration}

	return r
}

func (r RedisCaching) ExpireAt(t time.Time) util_entity.Caching {
	r.commands[1] = []interface{}{"EXPIREAT", r.key, t.Unix()}

	return r
}

func (r RedisCaching) Do(ctx context.Context) error {
	var err error

	for _, c := range r.commands {
		if err = r.conn.Do(ctx, c...).Err(); err != nil {
			log.Printf("fail excute redis command %v", c...)
			return err
		}
		log.Printf("excute redis command %v", c...)
	}

	return nil
}

func (r RedisCaching) GetList(ctx context.Context, key string, from, to int64) ([]string, error) {
	val, err := r.conn.LRange(ctx, key, from, to).Result()

	if err != nil && err != redis.Nil {
		log.Printf("GET LIST ERROR %s", err.Error())
		return []string{}, errors.New("fail get redis data")
	}

	return val, nil
}

func (r RedisCaching) Get(ctx context.Context, key string) (string, error) {
	val, err := r.conn.Get(ctx, key).Result()

	if err != redis.Nil {
		log.Printf("GET ERROR %s", err.Error())
		return "", errors.New("fail get redis data")
	}

	return val, nil
}
func (r RedisCaching) Del(ctx context.Context, key string) error {
	if err := r.conn.Del(ctx, key).Err(); err != nil {
		return errors.New("fail delete redis data")
	}

	return nil
}
