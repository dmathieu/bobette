package repo

import (
	"errors"
	"io"
	"strings"
)

// Pull fetches the code for a repository into the specified directory
func Pull(dir, url, auth string, stdout, stderr io.Writer) error {
	var user, password string
	if auth != "" {
		a := strings.Split(auth, ":")
		user, password = a[0], a[1]
	}

	switch true {
	case strings.HasSuffix(url, ".git"):
		return pullGit(dir, url, user, password, stdout, stderr)
	default:
		return errors.New("unknown repo type")
	}
}
