package cmd

import (
	"github.com/hammadmajid/gotic/internal/app"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new site",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the app instance from context
		app := cmd.Context().Value(appKey).(*app.App)

		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			name = "my-site" // TODO: get name from stdin using bubbletea
		}
		app.Logger.Printf("Creating new site: %s", name)

		err := app.Project.Create(name)

		// TODO: add biolerplate code

		return err
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
