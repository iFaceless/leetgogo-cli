package command

import "github.com/spf13/cobra"

func init() {
	cmdSubmit := &cobra.Command{
		Use:   "exec [--solution-filename filename] problem_slug",
		Short: "submit solution code to leetcode server",
		Args:  cobra.ExactArgs(1),
		Run:   handleSubmitSolutionCommand,
	}

	cmdSubmit.Flags().StringVar(&solutionFilename, "solution-filename", "solution/solution.go", "path to solution filename")
	rootCmd.AddCommand(cmdSubmit)
}
func handleSubmitSolutionCommand(cmd *cobra.Command, args []string) {

}
