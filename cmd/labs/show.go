package labs

import (
	"fmt"

	"github.com/databricks/cli/cmd/labs/project"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/spf13/cobra"
)

func newShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show NAME",
		Args:  cobra.ExactArgs(1),
		Short: "Shows information about the project",
		Annotations: map[string]string{
			"template": cmdio.Heredoc(`
			Name: {{.name}}
			Description: {{.description}}
			Python: {{.is_python}}

			Folders:
			 - lib: {{.lib_dir}}
			 - cache: {{.cache_dir}}
			 - config: {{.config_dir}}

			`),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			installed, err := project.Installed(ctx)
			if err != nil {
				return err
			}
			if len(installed) == 0 {
				return fmt.Errorf("no projects found")
			}
			name := args[0]
			for _, v := range installed {
				isDev := name == "." && v.IsDeveloperMode(ctx)
				isMatch := name == v.Name
				if !(isDev || isMatch) {
					continue
				}
				return cmdio.Render(ctx, map[string]any{
					"name":        v.Name,
					"description": v.Description,
					"cache_dir":   v.CacheDir(ctx),
					"config_dir":  v.ConfigDir(ctx),
					"lib_dir":     v.EffectiveLibDir(ctx),
					"is_python":   v.IsPythonProject(ctx),
				})
			}
			return nil
		},
	}
}
