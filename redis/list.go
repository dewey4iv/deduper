package redisdeduper

import (
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

func NewList(logger *logrus.Logger, redis *redis.Pool, redisKey string) *List {
	return &List{
		logger,
		redis,
		redisKey,
	}
}

type List struct {
	logger   *logrus.Logger
	redis    *redis.Pool
	redisKey string
}

func (l *List) Add(option string) error {
	conn := l.redis.Get()
	defer conn.Close()

	if _, err := conn.Do("SADD", l.redisKey, option); err != nil {
		return err
	}

	return nil
}

func (l *List) List() ([]string, error) {
	conn := l.redis.Get()
	defer conn.Close()

	results, err := redis.Strings(conn.Do("SMEMBERS", l.redisKey))
	if err != nil {
		return nil, err
	}

	return results, nil
}
