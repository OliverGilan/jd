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
	"jd/models"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

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
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		projectCode := config.GetActiveProject(cwd);
		if projectCode == 0{
			panic("No active project. You must create at least one project before creating any areas")
		}
		project := config.Projects[projectCode]

		var nextCode int;
		if code, err := cmd.Flags().GetInt("code"); err == nil {
			if !project.IsAreaCodeAvailable(code) {
				panic("Code is already in use: " + project.Areas[code].Name)
			}
			nextCode = code
		} else {
			if nc := project.GetNextAreaCode(); nc == -1 {
				panic("All project codes taken.")
			} else {
				nextCode = nc
			}
		}

		dirName := strconv.Itoa(nextCode) + "-" + strconv.Itoa(nextCode + 9) + " " + args[0]

		if e := os.Chdir(config.Paths[projectCode]); e != nil {
			panic(e)
		}
		if e := os.Mkdir(dirName, 0755); e != nil {
			panic(e)
		}

		area := models.Area{
			Code: nextCode,
			Name: dirName,
			Categories: make(map[int]models.Category, 10),
		};

		project.Areas[nextCode] = area;

		if e := config.SaveConfig(); e != nil {
			panic(e)
		}
		os.Chdir(cwd)
	},
	Args: cobra.MaximumNArgs(2),
}

func init() {
	rootCmd.AddCommand(areaCmd)
	
	areaCmd.Flags().IntP("code", "c", -1, "Set as code for new project")
	areaCmd.Flags().StringP("rename", "r", "", "Rename area")
	areaCmd.Flags().Bool("rm", false, "Delete an area and all child categories and items")
}
