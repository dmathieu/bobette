package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/dmathieu/bobette/exec"
	"github.com/dmathieu/bobette/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "builder",
	Short: "A docker wrapper to perform builds",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := readConfig()
		if err != nil {
			return err
		}

		workDir, err := ioutil.TempDir("", "builder")
		if err != nil {
			return err
		}
		defer os.RemoveAll(workDir)

		url := viper.GetString("repository_url")
		return handleBuild(workDir, url, os.Stdout, os.Stderr)
	},
}

func readConfig() error {
	config, err := base64.StdEncoding.DecodeString(os.Getenv("BOBETTE_CONFIG"))
	if err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(config))
	return nil
}

func handleBuild(dir, url string, stdout, stderr io.Writer) error {
	fmt.Printf("Fetching %s\n", url)
	err := repo.Pull(dir, url)
	if err != nil {
		return err
	}

	viper.SetConfigFile(path.Join(dir, "bobette.yml"))
	viper.ReadInConfig()

	fmt.Printf("Executing commands\n")
	commands := viper.GetStringSlice("commands")
	return exec.Execute(dir, commands, stdout, stderr)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
