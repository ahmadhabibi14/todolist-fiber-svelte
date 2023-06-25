package main

import (
	"fmt"
	"sort"
	"time"
)

type Item struct {
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	UserID    int
}

type ByUpdatedAt []Item

func (a ByUpdatedAt) Len() int           { return len(a) }
func (a ByUpdatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdatedAt) Less(i, j int) bool { return a[i].UpdatedAt.After(a[j].UpdatedAt) }

func main() {
	items := []Item{
		{
			Text:      "Item 1",
			CreatedAt: time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2023, 6, 5, 0, 0, 0, 0, time.UTC),
			Username:  "user1",
			UserID:    1,
		},
		{
			Text:      "Item 2",
			CreatedAt: time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2023, 6, 20, 0, 0, 0, 0, time.UTC),
			Username:  "user2",
			UserID:    2,
		},
		{
			Text:      "Item 3",
			CreatedAt: time.Date(2023, 6, 5, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2023, 6, 10, 0, 0, 0, 0, time.UTC),
			Username:  "user3",
			UserID:    3,
		},
	}

	sort.Sort(ByUpdatedAt(items))

	for _, item := range items {
		fmt.Printf("Text: %s, UpdatedAt: %s\n", item.Text, item.UpdatedAt)
	}
}
