package repository

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/otsukatsuka/chat_line_bot/domain/model"
	"log"
)

var key = "message"

type Store interface {
	GetMessages() (model.Messages, error)
	SetMessages(value []byte) error
}

type store struct {
	conn redis.Conn
}

func (s store) GetMessages() (model.Messages, error) {
	var data []byte
	data, err := redis.Bytes(s.conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	var messages model.Messages
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return nil, err
	}
	return messages, err
}

func (s store) SetMessages(value []byte) error {
	_, err := s.conn.Do("SETEX", key, 300, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		log.Printf("error setting key %s to %s: %v", key, v, err)
		return err
	}
	return nil
}

func NewStore(conn redis.Conn) Store {
	return &store{conn: conn}
}
