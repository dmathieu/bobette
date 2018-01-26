package app

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/dmathieu/bobette/k8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var envTemplate = template.Must(template.New("notificationContent").Parse(`Env configuration:

* Name: {{.ObjectMeta.Name}}
* Namespace: {{.ObjectMeta.Namespace}}

{{if .Data}}Secrets:
{{range $k, $v := .Data}}
  * {{$k}}
{{end}}{{else}}There are no secrets to show at the moment.{{end}}
Use 'bobette env set' to set a new secret.
`))

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "View and set build environment variables",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := homeDir()
		k, err := k8.New(filepath.Join(home, ".kube", "config"), k8.Arch(arch))
		if err != nil {
			return err
		}

		config, err := k.GetSecret(viper.GetString("repository"))
		if err != nil {
			return err
		}
		return envTemplate.Execute(os.Stdout, config)
	},
}

var envSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a build environment variables",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		home := homeDir()
		k, err := k8.New(filepath.Join(home, ".kube", "config"), k8.Arch(arch))
		if err != nil {
			return err
		}

		data := strings.Split(args[0], "=")
		fmt.Fprintf(os.Stdout, "Setting %s config...", data[0])
		err = k.SetSecret(viper.GetString("repository"), data[0], []byte(data[1]))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, " done\n")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envSetCmd)
}
