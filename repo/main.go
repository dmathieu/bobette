package repo

import (
	"errors"
	"strings"
)

// Pull fetches the code for a repository into the specified directory
func Pull(dir, url, auth string) error {
	var user, password string
	if auth != "" {
		a := strings.Split(auth, ":")
		user, password = a[0], a[1]
	}

	switch true {
	case strings.HasSuffix(url, ".git"):
		return pullGit(dir, url, user, password)
	default:
		return errors.New("unknown repo type")
	}
}
