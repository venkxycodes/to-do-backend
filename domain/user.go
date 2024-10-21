package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId      int64              `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Username    string             `json:"username" bson:"username"`
	Password    string             `json:"password" bson:"password"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
}

type UserMap struct {
	sync.RWMutex
	M          map[string]*User
	LastUserId int64
}

func (u *UserMap) Get(username string) (*User, int64) {
	u.RLock()
	defer u.RUnlock()
	if userRecord, isPresent := u.M[username]; isPresent {
		return userRecord, u.LastUserId
	} else {
		return nil, u.LastUserId
	}
}

func (u *UserMap) Set(user *User) {
	u.Lock()
	defer u.Unlock()
	if user != nil {
		u.M[user.Username] = user
		u.LastUserId += 1
	}
}
