package repo

import (
	"errors"
	"strings"
)

// Pull fetches the code for a repository into the specified directory
func Pull(dir, url string) error {
	switch true {
	case strings.HasSuffix(url, ".git"):
		return pullGit(dir, url)
	default:
		return errors.New("unknown repo type")
	}
}
