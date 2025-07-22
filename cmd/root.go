package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hammadmajid/gotic/internal/app"
	"github.com/spf13/cobra"
)

type contextKey string

const appKey contextKey = "app"

var rootCmd = &cobra.Command{
	Use:   "gotic",
	Short: "Gotic is a static site generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		return err
	},
}

func Execute(app *app.App) {
	ctx := context.WithValue(context.Background(), appKey, app)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
