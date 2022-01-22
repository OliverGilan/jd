package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "Crate new item directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("item called")
	},
}

func init() {
	rootCmd.AddCommand(itemCmd)

	itemCmd.Flags().StringP("rename", "r", "", "Rename category")
	itemCmd.Flags().Bool("rm", false, "Delete an area category and all child items")
}
