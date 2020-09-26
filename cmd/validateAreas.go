package cmd

import (
	"github.com/brianseitel/mudder/internal/world"
	"github.com/spf13/cobra"
)

// loadAreasCmd represents the loadAreas command
var loadAreasCmd = &cobra.Command{
	Use:   "areas",
	Short: "Load the areas into memory",
	Long:  `Load all the areas into memory`,
	Run: func(cmd *cobra.Command, args []string) {
		world.Load()
	},
}

func init() {
	rootCmd.AddCommand(loadAreasCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadAreasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadAreasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
