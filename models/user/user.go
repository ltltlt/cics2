package user

import (
	"encoding/json"

	"cytus2.rocks/cics2/dao/redis"
)

// User is a struct contains user related information
type User struct {
	ID   string
	Name string
	Pass string
}

// GetUserByID accept a user id, return a user struct, or a error
func GetUserByID(id string) (*User, error) {
	client := redis.MustGetClient("user")
	result, err := client.Get(id).Result()
	if err != nil {
		return nil, err
	}
	var u *User
	if err = json.Unmarshal([]byte(result), u); err != nil {
		return nil, err
	}
	return u, nil
}

// GetAllUsers gets all users
func GetAllUsers() []*User {
	var (
		client = redis.MustGetClient("user")
		cursor uint64
		users  []*User
	)
	for {
		keys, c, err := client.Scan(cursor, "*", 10).Result()
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			var u User
			us, err := client.Get(key).Result()
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal([]byte(us), &u)
			if err != nil {
				panic(err)
			}
			users = append(users, &u)
		}
		if c == 0 {
			break
		}
		cursor = c
	}
	return users
}
