package models

import (
	"time"
)

type (
	Todo struct {
		Id         string    `json:"id"`
		Text       string    `json:"text"`
		Created_At time.Time `json:"created_at"`
		Updated_At time.Time `json:"updated_at"`
		Username   string    `json:"username"`
		UserId     int       `json:"user_id"`
	}
	AddTextStore struct {
		Text string `json:"text"`
	}
	EditTextStore struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}
	DeleteTodoStore struct {
		Id string `json:"id"`
	}
	ByUpdatedAt []Todo
)

var (
	Todos                = []Todo{}
	TotalTodoDeleted int = 0
	TotalTodoCreated int = 0
)

func (a ByUpdatedAt) Len() int           { return len(a) }
func (a ByUpdatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdatedAt) Less(i, j int) bool { return a[i].Updated_At.After(a[j].Updated_At) }

func DeleteTodo(arr []Todo, index int) []Todo {
	// Create a new slice with length reduced by 1
	result := make([]Todo, len(arr)-1)

	// Copy elements before the index
	copy(result, arr[:index])

	// Copy elements after the index
	copy(result[index:], arr[index+1:])

	return result
}
