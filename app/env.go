package app

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/dmathieu/bobette/k8"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var envTemplate = template.Must(template.New("notificationContent").Parse(`Env configuration:

* Name: {{.ObjectMeta.Name}}
* Namespace: {{.ObjectMeta.Namespace}}

{{if .Data}}Secrets:
{{range $k, $v := .Data}}
  * {{$k}} - {{$v}}
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
		return nil
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envSetCmd)
}
