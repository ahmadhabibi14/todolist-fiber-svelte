package models

import "time"

type Todo struct {
	Text       string    `json:"text"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
	Username   string    `json:"username"`
}

var Todos = []Todo{
	{
		Text:       "Test 124",
		Created_At: time.Now().UTC(),
		Updated_At: time.Now().UTC(),
		Username:   "habi",
	},
}
