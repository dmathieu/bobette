package app

import (
	"os"
	"path/filepath"

	"github.com/dmathieu/bobette/k8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var arch string

var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "Build and ship the current folder's stack image",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := homeDir()
		k, err := k8.New(filepath.Join(home, ".kube", "config"), k8.Arch(arch))
		if err != nil {
			return err
		}

		return k.RunBuild(viper.GetString("repository"))
	},
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func init() {
	shipCmd.Flags().StringVar(&arch, "arch", "", "architecture to run the build on")
	rootCmd.AddCommand(shipCmd)
}
