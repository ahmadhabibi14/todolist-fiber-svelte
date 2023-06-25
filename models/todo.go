package models

import "time"

type Todo struct {
	Text       string    `json:"text"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Username   string    `json:"username"`
	UserId     int       `json:"user_id"`
}

var Todos = []Todo{}
