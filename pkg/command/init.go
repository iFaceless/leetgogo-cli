package command

import "github.com/spf13/cobra"

func init() {
	initCmd := &cobra.Command{
		Use:   "init [path to workdir]",
		Short: "initialize your workdir, code templates will be generated in your workdir",
		Args:  cobra.MaximumNArgs(1),
		Run:   handleInitCommand,
	}
	rootCmd.AddCommand(initCmd)
}

func handleInitCommand(cmd *cobra.Command, args []string) {

}
