package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatar {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatar when no avatar url is available")
	}

	testURL := "http://gravatar_url_addrs/"
	client.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should not return error when avatar url is available")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct avatar url")
	}

}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "0bc83cb571cd1c50ba6f3e8a78ef1346"} // generated md5 from lowercase MyEmailAddress@example.com
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Errorf("GravatarAvatar.GetAvatarURL returned wrong: %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer os.Remove(filename)
	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return error when avatar url is available")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL returned wrong: %s", url)
	}
}
