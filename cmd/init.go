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
	"errors"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new Johnny Decimal system in the current directory.",
	Long: `Init creates the metadata for a new Johnny Decimal system and either assigns 
	a project code automatically or uses the given code.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, _ := os.Getwd()
		
		nextCode, err := getNextProjectCode(cmd.Flags().GetInt("code"))
		if err != nil{
			panic(err)
		}

		setName(nextCode, args[0])
		setPath(nextCode, cwd)

		if def, _ := cmd.Flags().GetBool("default"); def {
			viper.Set("default_project", nextCode)
		}
		if phys, _ := cmd.Flags().GetBool("physical"); phys {
			err := os.Mkdir(strconv.Itoa(nextCode) + args[0], 0755)
			if err != nil{
				panic(err)
			}
		}

		if error := viper.WriteConfig(); error != nil {
			panic(error)
		}
	},
	Args: cobra.MaximumNArgs(1), //The only arg should be one positional arg for the name of the new system
}

func setPath(code int, cwd string){
	projectPaths := viper.GetStringMapString("project_paths")
	if(projectPaths == nil){
		projectPaths = make(map[string]string, 899)
	}
	projectPaths[strconv.Itoa(code)] = cwd
	viper.Set("project_names", projectPaths)
}

func setName(code int, name string){
	projectNames := viper.GetStringMapString("project_names")
	if(projectNames == nil){
		projectNames = make(map[string]string, 899)
	}
	projectNames[strconv.Itoa(code)] = name
	viper.Set("project_names", projectNames)
}
 

func getNextProjectCode(fcode int, e error) (int, error){
	var code int;
	projectCodes := viper.GetIntSlice("projects")

	if e == nil {
		if fcode > 100 && fcode < 1000{
			if(projectCodes[fcode-100] != 0){
				return -1, errors.New("code in use")
			}
			return fcode, nil
		}else{
			return -1, errors.New("code must be a 3-digit integer")
		}
	}else{
		if projectCodes == nil {
			projectCodes = make([]int, 899)
			code = 100
			projectCodes[0] = 100
		}
		for i := 0; i < cap(projectCodes); i++{
			if(projectCodes[i] == 0){
				code = i + 100
				projectCodes[i] = code
				break
			}
		}
	}

	viper.Set("projects", projectCodes)
	return code, nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().IntP("code", "c", -1, "Set as code for new project")
	initCmd.Flags().BoolP("default", "d", false, "Set as default project after creating")
	initCmd.Flags().BoolP("physical", "p", false, "Create directory to hold new project")
}
