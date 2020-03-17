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
	"image/color"
	"os"

	"github.com/spf13/cobra"

	"github.com/mangelajo/submariner-release-tools/ghs/pkg/github"
)

// createLabelsCmd represents the createLabels command
var createLabelsCmd = &cobra.Command{
	Use:   "create-labels",
	Short: "This command creates labels on an a list of organization projects at once",
	Long:  "two arguments are expected: <label name> and <color in hex RGB rrggbb>",
	Run:   createLabels,
	Args:  cobra.MinimumNArgs(2),
}

var labelDescription string

func init() {
	rootCmd.AddCommand(createLabelsCmd)
	createLabelsCmd.Flags().StringVarP(&labelDescription, "description", "d", "",
		"Description for the label being created")
}

func createLabels(cmd *cobra.Command, args []string) {

	labelName := args[0]
	labelColor := args[1]

	colorVal, err := ParseHexColor(labelColor)

	exitOnError(err, "you need to provide a valid label color as 2nd argument")

	labelColor = fmt.Sprintf("%02x%02x%02x", colorVal.R, colorVal.G, colorVal.B)

	fmt.Printf("Creating label %s for repos: %v in org: %s\n", labelName, reposParameter, ownerParameter)
	gh, err := github.NewGitHub()
	exitOnError(err, "Creating github client")

	for _, repo := range reposParameter {
		_, _, err := gh.CreateLabel(ownerParameter, repo, labelName, labelColor, labelDescription)

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

// source: https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 6:
		_, err = fmt.Sscanf(s, "%02x%02x%02x", &c.R, &c.G, &c.B)
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 3:
		_, err = fmt.Sscanf(s, "%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")
	}
	return
}
