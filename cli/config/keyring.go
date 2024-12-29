package config

import (
	"os/user"

	"github.com/zalando/go-keyring"
)

type keyringSource struct{}

func (s *keyringSource) Lookup(key string) (any, bool) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, false
	}

	value, err := keyring.Get(key, currentUser.Username)
	if err != nil {
		return nil, false
	}

	return value, true
}

func (s *keyringSource) String() string {
	return "OS Keyring"
}

func (s *keyringSource) GoString() string {
	return "&keyringSource{}"
}

var KeyringSource = &keyringSource{}

func SetKeyring(key string, value string) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	return keyring.Set(keyringPath(key), currentUser.Username, value)
}

func keyringPath(key string) string {
	return "markxus." + key
}
