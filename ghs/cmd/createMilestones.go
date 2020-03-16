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

	"github.com/mangelajo/submariner-release-tools/ghs/pkg/github"
)

// createMilestonesCmd represents the createMilestones command
var createMilestonesCmd = &cobra.Command{
	Use:   "create-milestones",
	Short: "This command creates milestones on an a list of organization projects at once",
	Run: createMilestones,
}

var projects []string

func init() {
	rootCmd.AddCommand(createMilestonesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMilestonesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMilestonesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//createMilestonesCmd.PersistentFlags().StringSliceVarP(&projects,"project", "p",[]string{},)
}

func createMilestones(cmd *cobra.Command, args []string) {
	fmt.Println("createMilestones called")
	gh, err := github.NewGitHub()
	exitOnError(err, "Creating github client")
	milestone, err := gh.CreateMilestone("submariner-io", "submariner", "0.3.0")
	exitOnError(err, "Creating milestone")
	fmt.Printf("Milestone created: %v",milestone)
}

func exitOnError(err error, indication string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s: %s\n", indication, err)
		os.Exit(1)
	}
}