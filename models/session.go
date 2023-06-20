package models

import "time"

// Untuk key nya, menggunakan token
var Sessions = map[string]Session{}

// Each Session contains the username of the user and the time at which it expires
type Session struct {
	Username string
	Expiry   time.Time
}

// We'll use this method later to determine if the Session has expired
func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}
