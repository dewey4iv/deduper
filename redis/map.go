package redisdeduper

import (
	"sort"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

const (
	sep = ":"
)

func NewMap(l *logrus.Logger, redis *redis.Pool, prefix string) *Map {
	return &Map{
		l,
		redis,
		prefix,
	}
}

type Map struct {
	l      *logrus.Logger
	redis  *redis.Pool
	prefix string
}

func (m *Map) Add(key string, option string) error {
	conn := m.redis.Get()
	defer conn.Close()

	if _, err := conn.Do("SADD", m.prefix+sep+key, option); err != nil {
		return err
	}

	return nil
}

func (m *Map) Map() (map[string][]string, error) {
	conn := m.redis.Get()
	defer conn.Close()

	allMembers := make(map[string][]string)

	redisKeys, err := redis.Strings(conn.Do("KEYS", m.prefix+sep+"*"))
	if err != nil {
		return nil, err
	}

	for _, rKey := range redisKeys {
		members, err := redis.Strings(conn.Do("SMEMBERS", rKey))
		if err != nil {
			return nil, err
		}

		sort.Strings(members)

		allMembers[strings.Replace(rKey, m.prefix+sep, "", -1)] = members
	}

	return allMembers, nil
}
