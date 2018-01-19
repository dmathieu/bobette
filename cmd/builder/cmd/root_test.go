package cmd

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"

	githttp "github.com/AaronO/go-git-http"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHandleBuild(t *testing.T) {
	git := githttp.New("../../../fixtures")
	s := httptest.NewServer(git)
	defer s.Close()

	dir, err := ioutil.TempDir("", "root")
	assert.Nil(t, err)

	var stdout, stderr bytes.Buffer
	err = handleBuild(dir, s.URL+"/repo.git", &stdout, &stderr)
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "hello world\n")
	assert.Equal(t, "", stderr.String())
}

func TestReadConfig(t *testing.T) {
	t.Run("with no config set", func(t *testing.T) {
		err := readConfig()
		assert.Nil(t, err)
		assert.Equal(t, nil, viper.Get("repository_url"))
	})

	t.Run("with a config set", func(t *testing.T) {
		os.Setenv("BOBETTE_CONFIG", base64.StdEncoding.EncodeToString([]byte(`{"repository_url":"https://example.com"}`)))
		defer os.Setenv("BOBETTE_CONFIG", "")

		err := readConfig()
		assert.Nil(t, err)
		assert.Equal(t, "https://example.com", viper.Get("repository_url"))
	})
}
