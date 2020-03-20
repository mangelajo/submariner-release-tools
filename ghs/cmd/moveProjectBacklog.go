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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// moveProjectBacklog represents the createMilestones command
var moveProjectBacklogCmd = &cobra.Command{
	Use:   "move-project-backlog",
	Short: "This command moves the backlog column from a project to another project (sprints)",
	Run:   moveProjectBacklog,
	Args:  cobra.MinimumNArgs(2),
}

var backlogColumn string

func init() {
	rootCmd.AddCommand(moveProjectBacklogCmd)
	moveProjectBacklogCmd.Flags().StringVarP(&backlogColumn, "backlog-column", "b", "Backlog",
		"The backlog column that needs to be moved from one project sprint to the next")
}

func moveProjectBacklog(cmd *cobra.Command, args []string) {

	fmt.Printf("Moving project backlog ..\n"+
		"from: %s\n"+
		"to: %s\n"+
		"org: %s\n", args[0], args[1], ownerParameter)

	gh := githubClient()

	projects, err := gh.GetProjectIDs(ownerParameter, args)
	exitOnError(err, "Looking up for the projects")
	srcColID, err := gh.GetProjectColumn(projects[0], backlogColumn)
	exitOnError(err, "Could not find column in source project")
	dstColID, err := gh.GetProjectColumn(projects[1], backlogColumn)
	exitOnError(err, "Could not find column in source project")

	srcCol, _, err := gh.GetColumnCards(srcColID)

	for _, card := range srcCol {
		data, _ := json.MarshalIndent(card, "", "\t")
		fmt.Println(string(data))
		_, _, err := gh.MoveCardToAnotherProjectColumn(card, dstColID)
		exitOnError(err, "moving card")
	}

}
