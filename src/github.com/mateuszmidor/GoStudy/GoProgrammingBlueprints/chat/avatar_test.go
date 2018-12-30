package main

import "testing"

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
	client.userData = map[string]interface{}{"email": "MyEmailAddress@example.com"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Errorf("GravatarAvatar.GetAvatarURL returned wrong: %s", url)
	}
}
