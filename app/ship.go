package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

// github.com/dmathieu/bobette/cmd/bobette/shipCmd represents the github.com/dmathieu/bobette/cmd/bobette/ship command
var shipCmd = &cobra.Command{
	Use:   "ship",
	Short: "Build and ship the current folder's stack image",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shipping")
	},
}
