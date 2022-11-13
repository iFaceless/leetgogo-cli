package command

import "github.com/spf13/cobra"

var favoriteName string

func init() {
	favCmd := &cobra.Command{
		Use:   "favorite [--favorite-name <example-collection>] problem-slug [...problem-slug]",
		Short: "favorite problem by slug",
		Args:  cobra.MinimumNArgs(1),
		Run:   handleFavoriteCommand,
	}
	favCmd.Flags().StringVar(&favoriteName, "favorite-name", "Favorite", "custom favorite name")
	rootCmd.AddCommand(favCmd)

	unfavCmd := &cobra.Command{
		Use:   "unfavorite [--favorite-name <example-collection>] problem-slug [...problem-slug]",
		Short: "unfavorite problem by slug",
		Args:  cobra.MinimumNArgs(1),
		Run:   handleUnfavoriteCommand,
	}
	unfavCmd.Flags().StringVar(&favoriteName, "favorite-name", "Favorite", "custom favorite name")
	rootCmd.AddCommand(unfavCmd)
}

func handleFavoriteCommand(cmd *cobra.Command, args []string) {

}

func handleUnfavoriteCommand(cmd *cobra.Command, args []string) {

}
