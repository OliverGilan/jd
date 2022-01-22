package cmd

import (
	"fmt"
	"jd/models"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var config models.Config;


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jd.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".jdconfig" (without extension).
		viper.SetConfigName(".jdconfig")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
			viper.Set("current_project", nil)
			if error := viper.SafeWriteConfig(); error != nil {
				fmt.Println(error)
			}
		} else {
			// Config file was found but another error was produced
			fmt.Fprintln(os.Stderr, "Error retreiving config file: ", err)
		}
	}
	if err := viper.Unmarshal(&config); err != nil{
		panic(err)
	}
	// path, _ := os.Getwd()
	// fmt.Println(strings.SplitAfter(path, "/"))
	// fmt.Println(config.Projects[100])
	// newProject := models.Project{Code: 200, Name: "Gabagool"}
	// config.Projects[200] = newProject
	// fmt.Println(config.Projects[200])
	// viper.Set("paths", config.Paths)
	// viper.Set("projects", config.Projects)
	// viper.WriteConfig()
}
