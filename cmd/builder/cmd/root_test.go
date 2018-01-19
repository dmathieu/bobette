package cmd

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	githttp "github.com/AaronO/go-git-http"
	"github.com/stretchr/testify/assert"
)

func TestHandleBuild(t *testing.T) {
	git := githttp.New("../../../fixtures")
	s := httptest.NewServer(git)
	defer s.Close()

	dir, err := ioutil.TempDir("", "root")
	assert.Nil(t, err)

	var stdout, stderr bytes.Buffer
	err = handleBuild(dir, s.URL+"/repo.git", "", &stdout, &stderr)
	assert.Nil(t, err)
	assert.Contains(t, stdout.String(), "hello world\n")
	assert.Equal(t, "", stderr.String())
}
