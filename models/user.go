package models

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	LoggedIn bool   `json:"loggedin"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UserStores = []User{}
var IdAutoInc = uint64(1)

func (u *User) NewUser(data Credentials) {
	newUser := User{
		Id:       IdAutoInc,
		Username: data.Username,
		Password: data.Password,
		LoggedIn: true,
	}
	UserStores = append(UserStores, newUser)
	IdAutoInc++
}
