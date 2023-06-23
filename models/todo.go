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
		Text:       "Mau main bola",
		Created_At: time.Now().UTC(),
		Updated_At: time.Now().UTC(),
		Username:   "habi",
	},
	{
		Text:       "Rasa yang dulu pernah ada",
		Created_At: time.Date(2011, 2, 24, 10, 20, 0, 0, time.UTC),
		Updated_At: time.Date(2023, 2, 24, 10, 20, 0, 0, time.UTC),
		Username:   "hana",
	},
}
