/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newAreaName string;

// areaCmd represents the area command
var areaCmd = &cobra.Command{
	Use:   "area [<PRO>] [<name>]",
	Short: "Create a system area",
	Long: `Create a system area for a specific project. The command will automatically determine the project
	to create the area in by checking <PRO> argument, current working directory, default project in that order.
	If <PRO> is specified the area will be created in that system, if not but the cwd is within an existing 
	project then that system will be used, otherwise the current system will be used.
	
	Rename existing area with -r NAME.
	
	Delete existing area with -rm.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("area called")
	},
	Args: cobra.MaximumNArgs(2),
}

func init() {
	rootCmd.AddCommand(areaCmd)
	
	areaCmd.Flags().StringVarP(&newAreaName ,"rename", "r", "", "Rename area")
	areaCmd.Flags().BoolP("remove", "rm", false, "Delete an area directory and all child categories and items")
	
}
