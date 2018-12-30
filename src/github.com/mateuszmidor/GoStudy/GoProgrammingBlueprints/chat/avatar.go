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
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// GetAvatarURL returns avatar url for given client
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatar
}

type GravatarAvatar struct{}

var UseGravatarAvatar GravatarAvatar

// GetAvatarURL returns avatar url for given client
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			return "//gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatar
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

// GetAvatarURL returns avatar url for given client
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatar
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(useridStr+"*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}

			}
		}
	}
	return "", ErrNoAvatar
}
