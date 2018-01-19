package cmd

import (
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
		workDir, err := ioutil.TempDir("", "builder")
		if err != nil {
			return err
		}
		defer os.RemoveAll(workDir)

		return handleBuild(
			workDir,
			os.Getenv("REPO_URL"),
			os.Getenv("REPO_AUTH"),
			os.Stdout,
			os.Stderr,
		)
	},
}

func handleBuild(dir, url, auth string, stdout, stderr io.Writer) error {
	fmt.Printf("Fetching %s\n", url)
	err := repo.Pull(dir, url, auth)
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
