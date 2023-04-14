package middleware

import (
	"encoding/json"
	"time"
)

type RedisLoginCache struct {
	Id              uint      `json:"id"`
	Token           string    `json:"token"`
	LoginTime       time.Time `json:"login_time"`
	LastRequestTime time.Time `json:"last_request_time"`
}

func (c RedisLoginCache) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

func (c *RedisLoginCache) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &c)
}
