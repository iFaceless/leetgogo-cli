package command

import (
	"context"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "leetgogo",
	Short:   "leetgogo is a command tool to work with leetcode nicely.",
	Version: "1.0.0",
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
