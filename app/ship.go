package app

import (
	"os"
	"path/filepath"

	"github.com/dmathieu/bobette/k8"
	"github.com/spf13/cobra"
)

// github.com/dmathieu/bobette/cmd/bobette/shipCmd represents the github.com/dmathieu/bobette/cmd/bobette/ship command
var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "Build and ship the current folder's stack image",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := homeDir()
		k, err := k8.New(filepath.Join(home, ".kube", "config"))
		if err != nil {
			return err
		}

		return k.RunBuild()
	},
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
