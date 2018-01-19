package repo

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	githttp "github.com/AaronO/go-git-http"
	gitauth "github.com/AaronO/go-git-http/auth"
	"github.com/stretchr/testify/assert"
)

func TestPull(t *testing.T) {
	t.Run("an unknown repo type", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "puller")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)

		err = Pull(dir, "https://example.com", "")
		assert.Equal(t, errors.New("unknown repo type"), err)
	})

	t.Run("a git repo type", func(t *testing.T) {
		git := githttp.New("../fixtures")
		s := httptest.NewServer(git)
		defer s.Close()

		dir, err := ioutil.TempDir("", "puller")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)

		err = Pull(dir, s.URL+"/repo.git", "")
		assert.Nil(t, err)
	})

	t.Run("a git repo with authentication", func(t *testing.T) {
		git := githttp.New("../fixtures")

		auth := gitauth.Authenticator(func(info gitauth.AuthInfo) (bool, error) {
			if info.Username == "me" && info.Password == "mypassword" {
				return true, nil
			}
			return false, nil
		})

		s := httptest.NewServer(auth(git))
		defer s.Close()

		dir, err := ioutil.TempDir("", "puller")
		assert.Nil(t, err)
		defer os.RemoveAll(dir)

		err = Pull(dir, s.URL+"/repo.git", "me:mypassword")
		assert.Nil(t, err)
	})
}
