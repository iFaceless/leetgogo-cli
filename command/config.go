package command

import "github.com/spf13/cobra"

func init() {
	configCmd := &cobra.Command{
		Use:     "config key [value]",
		Short:   "manage leetgogo config",
		Example: "config leetcode.session <SESSION>\nconfig leetcode.csrftoken <TOKEN>",
		Args:    cobra.MinimumNArgs(1),
		Run:     handleConfigCommand,
	}

	rootCmd.AddCommand(configCmd)
}

func handleConfigCommand(cmd *cobra.Command, args []string) {

}
