package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatar)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatar {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no url is available")
	}

	testURL := "http://gravatar_url_addrs/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return error when avatar url is available")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct avatar url")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL returned error while it shouldnt")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.GetAvatarURL returned wrong: %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	if err := os.MkdirAll("avatars", 0777); err != nil {
		t.Errorf("Could not create 'avatars' directory: %s", err)
	}
	if err := ioutil.WriteFile(filename, []byte{}, 0777); err != nil {
		t.Errorf("Could not create avatar file: %s", err)
	}
	defer os.Remove(filename)
	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return error when avatar url is available")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned wrong: %s", url)
	}
}
