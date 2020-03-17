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
	Run:   createMilestones,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(createMilestonesCmd)
}

func createMilestones(cmd *cobra.Command, args []string) {

	milestoneTitle := args[0]
	fmt.Printf("Creating milestone %s for repos: %v in org: %s\n", milestoneTitle, reposParameter, ownerParameter)
	gh, err := github.NewGitHub()
	exitOnError(err, "Creating github client")

	for _, repo := range reposParameter {
		_, _, err := gh.CreateMilestone(ownerParameter, repo, milestoneTitle)

		if err != nil {
			if github.IsAlreadyExistError(err) {
				fmt.Printf(" ✓ exists for project: %q\n", repo)
			} else {
				fmt.Fprintf(os.Stderr, " ✘ failed for %q: %s\n", repo, err)
				os.Exit(1)
			}
		} else {
			fmt.Printf(" ✔ created for project: %q\n", repo)
		}
	}
}
