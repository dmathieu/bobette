package exec

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	dir, err := ioutil.TempDir("", "puller")
	assert.Nil(t, err)
	defer os.RemoveAll(dir)
	var stdout, stderr bytes.Buffer

	err = Execute(dir, []string{"echo 'hello'", "echo 'world'"}, &stdout, &stderr)
	assert.Nil(t, err)
	assert.Equal(t, "hello\nworld\n", stdout.String())
	assert.Equal(t, "", stderr.String())
}
