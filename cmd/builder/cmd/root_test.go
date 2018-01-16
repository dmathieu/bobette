package cmd

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

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
