package cmd

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

		url := viper.Get("repository_url").(string)
		fmt.Printf("Fetching %s\n", url)
		err = repo.Pull(workDir, url)
		if err != nil {
			return err
		}

		viper.SetConfigFile(path.Join(workDir, "bobette.yml"))
		viper.ReadInConfig()

		fmt.Printf("Executing commands\n")
		commands := viper.Get("commands").([]string)
		err = exec.Execute(workDir, commands, os.Stdout, os.Stderr)
		if err != nil {
			return err
		}

		return nil
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
