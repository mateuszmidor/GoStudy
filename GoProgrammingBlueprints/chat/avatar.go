package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatar indicates that the avatar url can't be obtained
var ErrNoAvatar = errors.New("chat: can't fetch avatar url")

// Avatar represents chat user picture
type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatar
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// GetAvatarURL returns avatar url for given client
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatar
	}

	return url, nil
}

type GravatarAvatar struct{}

var UseGravatarAvatar GravatarAvatar

// GetAvatarURL returns avatar url for given client
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL returns avatar url for given client
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {

	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}

		}
	}
	return "", ErrNoAvatar
}
