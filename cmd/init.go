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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new Johnny Decimal system in the current directory.",
	Long: `Init creates the metadata for a new Johnny Decimal system and either assigns 
	a project code automatically or uses the given code.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err == nil {
			project := config.GetActiveProject(cwd)

			if project != 0 {
				panic("Can't initialize a project within an existing project")
			}
		} else {
			panic(err)
		}

		var nextCode int
		if code, err := cmd.Flags().GetInt("code"); err == nil {
			if !config.IsProjectCodeAvailable(code) {
				panic("Code is already in use: " + config.Projects[code].Name)
			}
			nextCode = code
		} else {
			if nc := config.GetNextProjectCode(); nc == -1 {
				panic("All project codes taken.")
			} else {
				nextCode = nc
			}
		}

		setName(nextCode, args[0])
		setPath(nextCode, cwd)

		if def, _ := cmd.Flags().GetBool("default"); (def || config.DefaultProject == 0) {
			config.DefaultProject = nextCode;
		}
		if phys, _ := cmd.Flags().GetBool("physical"); phys {
			err := os.Mkdir(strconv.Itoa(nextCode) + " " + args[0], 0755)
			if err != nil{
				panic(err)
			}
		}

		if e := config.SaveConfig(); e != nil {
			panic(e)
		}
	},
	Args: cobra.ExactArgs(1), //The only arg should be one positional arg for the name of the new system
}

func setPath(code int, cwd string) {
	if config.Paths == nil {
		config.Paths = make(map[int]string, 899)
	}
	config.Paths[code] = cwd
}

func setName(code int, name string) {
	if config.Projects == nil {
		config.Projects = make(map[int]models.Project, 899)
	}
	config.Projects[code] = models.Project{
		Code: code,
		Name: name,
		Areas: nil,
	};
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().IntP("code", "c", -1, "Set as code for new project")
	initCmd.Flags().BoolP("default", "d", false, "Set as default project after creating")
	initCmd.Flags().BoolP("physical", "p", false, "Create directory to hold new project")
}
