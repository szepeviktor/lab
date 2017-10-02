package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mrEditCmd represents the mrEdit command
var mrEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit current merge request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// figure out what mr to edit
		// list mrs
		// mr matching source branch?
		branch, err := git.CurrentBranch()
		if err != nil {
			log.Fatal(err)
		}

		sourceRemote, err := gitconfig.Local("branch." + branch + ".remote")
		if err != nil {
			sourceRemote = "origin"
		}
		sourceProjectName, err := git.PathWithNameSpace(sourceRemote)
		if err != nil {
			log.Fatal(err)
		}

		title, body, err := git.Edit("MERGEREQ", msg)
		if err != nil {
			_, f, l, _ := runtime.Caller(0)
			log.Fatal(f+":"+strconv.Itoa(l)+" ", err)
		}

		if title == "" {
			log.Fatal("aborting MR due to empty MR msg")
		}

		mrURL, err := lab.MergeRequest(sourceProjectName, &gitlab.UpdateMergeRequestOptions{
			SourceBranch:    &branch,
			TargetBranch:    gitlab.String(targetBranch),
			TargetProjectID: &targetProject.ID,
			Title:           &title,
			Description:     &body,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	mrCmd.AddCommand(mrEditCmd)
}
