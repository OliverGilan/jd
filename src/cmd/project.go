package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project [<code>]",
	Short: "Changes the default project",
	Long: `Change the default project for JD. Once set, the default project can be referenced in other commands without
	specifying the 3-digit project code.
	
	You can instead change project name with -r NAME. This does not affect the project code, merely the directory name.
	
	Delete the project directory and all subdirectories with -rm. If no project <code> arg is specified, the command
	will delete the current project.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(projectCmd)
	
	projectCmd.Flags().BoolP("default", "d", false, "Use with <code> argument to change the default project")
	projectCmd.Flags().Bool("rm", false, "Deletes project and all corresponding areas, categories, and items")
	projectCmd.Flags().StringP("rename", "r", "", "Rename default project or the project specified with <code> argument")
}
