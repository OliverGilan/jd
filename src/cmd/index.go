package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index [<PRO>]",
	Short: "Display system index",
	Long: `Display the system index. If no project code is provided the default system will be displayed.
	
	To display all system indices use the -a flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("index called")
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)

	indexCmd.Flags().BoolP("all", "a", false, "Display all system indices")

}
