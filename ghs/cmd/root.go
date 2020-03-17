/*
Copyright © 2020 Miguel Ángel Ajo <majopela@redhat.com>

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
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghs",
	Short: "GHS is Github tool for Submariner",
	Long: `This is a tool to help us create milestones over multiple
org projects at once, or move backlog cards from one milestone to the next,
etc.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var reposParameter []string
var ownerParameter string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ghs.yaml)")

	rootCmd.PersistentFlags().StringSliceVarP(&reposParameter, "repos", "r", []string{},
		"The repository names you want to be handled")
	rootCmd.PersistentFlags().StringVarP(&ownerParameter, "owner", "o", "",
		"The repository owner/organization where the repositories live")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ghs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ghs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func exitOnError(err error, indication string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s: %s\n", indication, err)
		os.Exit(1)
	}
}
