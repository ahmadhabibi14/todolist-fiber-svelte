package models

import "time"

type Todo struct {
	Id         string    `json:"id"`
	Text       string    `json:"text"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Username   string    `json:"username"`
	UserId     int       `json:"user_id"`
}

type AddTextStore struct {
	Text string `json:"text"`
}

type EditTextStore struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

var Todos = []Todo{}
