package command

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "pick [problem-slug | problem-url]",
		Short:   "pick a leetcode problem by slug or url with slug, default to a random one",
		Example: "leetgogo pick two-sum",
		Args:    cobra.MaximumNArgs(1),
		Run:     handlePickCommand,
	})
}

func handlePickCommand(cmd *cobra.Command, args []string) {

}
