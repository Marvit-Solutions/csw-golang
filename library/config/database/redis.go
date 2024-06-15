package database

import (
	"encoding/json"
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	Address   string `json:"address"`
	Port      string `json:"port"`
	Password  string `json:"password"`
	MaxIdle   string `json:"max_idle"`
	MaxActive string `json:"max_active"`
}

func InitDBRedis(env config.Config) (*redis.Pool, error) {
	var client Client

	c := env.GetStringMap("redis")

	jsonb, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(jsonb, &client); err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", client.Address, client.Port)

	// idle, err := strconv.Atoi(client.MaxIdle)
	idle := 1
	if err != nil {
		return nil, err
	}

	// active, err := strconv.Atoi(client.MaxActive)
	active := 1
	if err != nil {
		return nil, err
	}

	return &redis.Pool{
		MaxIdle:   idle,
		MaxActive: active,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", client.Password); err != nil {
				return nil, err
			}

			return c, nil
		},
	}, nil
}
