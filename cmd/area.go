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
	"github.com/spf13/viper"
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
		fmt.Println("area called:")
		// cwd, _ := os.Getwd()


		// if error := viper.WriteConfig(); error != nil {
		// 	panic(error)
		// }
		var c AreaConfig;
		viper.UnmarshalKey("projects.100.areas", &c)
		fmt.Println(c)
	},
	Args: cobra.MaximumNArgs(2),
}

type AreaConfig struct {
	Areas []map[int]string;
}

func (a AreaConfig) getNextAreaCode() (int, error){
	return 0, nil
}

func (a AreaConfig) isCodeAvailable(code int) (bool){
	// for _, elem := range a.Areas {
	// 	if
	// }
	return true
}

func getNextAreaCode(projectCode string, cwd string, acode int, err error) (int, error) {
	// var areaCode int
	// var projectCode int

	// projectPaths := viper.GetStringMapString("project_paths")
	// for key, elem := range projectPaths {
	// 	if strings.Contains(cwd, elem){
	// 		key, _ := strconv.Atoi(key)
	// 		projectCode = key
	// 		break
	// 	}
	// }
	// viper.GetSlice
	// areaCodes := viper.GetIntSlice(projectCode)

	// if err == nil {
	// 	if acode >= 10 && acode <= 90 && acode % 10 == 0{
	// 		if(areaCodes[acode-100] != 0){
	// 			return -1, errors.New("code in use")
	// 		}
	// 		return fcode, nil
	// 	}else{
	// 		return -1, errors.New("code must be a 3-digit integer")
	// 	}
	// }else{
	// 	if areaCodes == nil {
	// 		areaCodes = make([]int, 10)
	// 		areaCode = 10
	// 		areaCodes[0] = 10
	// 	}
	// 	for i := 0; i < cap(projectCodes); i++{
	// 		if(projectCodes[i] == 0){
	// 			code = i + 100
	// 			projectCodes[i] = code
	// 			break
	// 		}
	// 	}
	// }

	// viper.Set("projects", projectCodes)
	return 0, nil
}

func init() {
	rootCmd.AddCommand(areaCmd)
	
	areaCmd.Flags().IntP("code", "c", -1, "Set as code for new project")
	areaCmd.Flags().StringP("rename", "r", "", "Rename area")
	areaCmd.Flags().Bool("rm", false, "Delete an area and all child categories and items")
}
