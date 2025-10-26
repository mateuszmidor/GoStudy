package api

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

type User struct {
	Key         *datastore.Key `json:"id" datastore:"-"`
	UserID      string         `json:"-"`
	DisplayName string         `json:"display_name"`
	AvatarURL   string         `json:"avatar_url"`
	Score       int            `json:"score"`
}

func UserFromAEUser(ctx context.Context) (*User, error) {
	aeuser := user.Current(ctx)
	if aeuser == nil {
		return nil, errors.New("No logged in user")
	}
	var appUser User
	appUser.Key = datastore.NewKey(ctx, "User", aeuser.ID, 0, nil)
	err := datastore.Get(ctx, appUser.Key, &appUser)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return nil, err
	}
	if err == nil {
		return &appUser, nil
	}
	appUser.UserID = aeuser.ID
	appUser.DisplayName = aeuser.String()
	appUser.AvatarURL = gravatarURL(aeuser.Email)
	log.Infof(ctx, "Saving new user: %s", aeuser.String())
	appUser.Key, err = datastore.Put(ctx, appUser.Key, &appUser)
	if err != nil {
		return nil, err
	}
	return &appUser, nil
}

func gravatarURL(email string) string {
	m := md5.New()
	io.WriteString(m, strings.ToLower(email))
	return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil))
}

type UserCard struct {
	Key         *datastore.Key `json:"id"`
	DisplayName string         `json:"display_name"`
	AvatarURL   string         `json:"avatar_url"`
}

func (u User) Card() UserCard {
	return UserCard{
		Key:         u.Key,
		DisplayName: u.DisplayName,
		AvatarURL:   u.AvatarURL,
	}
}
