package command

import "github.com/spf13/cobra"

var (
	solutionFilename  string
	testCasesFilename string
)

func init() {
	cmdExec := &cobra.Command{
		Use:   "exec [--solution-filename filename] [--test-cases-filename filename] problem_slug",
		Short: "execute problem solution on remote leetcode servers",
		Args:  cobra.ExactArgs(1),
		Run:   handleExecuteSolutionCommand,
	}

	cmdExec.Flags().StringVar(&solutionFilename, "solution-filename", "solution/solution.go", "path to solution filename")
	cmdExec.Flags().StringVar(&testCasesFilename, "testcases-filename", "solution/testcases.txt", "path to solution testcases filename")
	rootCmd.AddCommand(cmdExec)

}

func handleExecuteSolutionCommand(cmd *cobra.Command, args []string) {

}
