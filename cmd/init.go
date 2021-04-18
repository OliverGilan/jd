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

var code string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new Johnny Decimal system in the current directory.",
	Long: `Init creates the metadata for a new Johnny Decimal system and either assigns 
	a project code automatically or uses the given code.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
	Args: cobra.MaximumNArgs(1), //The only arg should be one positional arg for the name of the new system
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&code, "code", "c", "", "project code to use")
}
