package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category [<code>] [<name>]",
	Short: "Create an area category",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("rename")
		fmt.Println("category called:", name)
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)

	categoryCmd.Flags().StringP("rename", "r", "", "Rename category")
	categoryCmd.Flags().Bool("rm", false, "Delete a category and all child items")
}
