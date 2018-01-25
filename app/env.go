package app

import (
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "View and set build environment variables",
	Long:  "",
}

var envSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a build environment variables",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envSetCmd)
}
