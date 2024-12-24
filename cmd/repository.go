/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Indiance/go-gh-analyzer/utils"
	"github.com/spf13/cobra"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "A command used to analyze a repository given a URL.",
	Long: `A command that can be called to analyze a repository given its URL. This is done similar to
	the root command, except a URL can be given. And thus any repository can be analyzed. Similar methods are used`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var owner, repo string
		listIssues, _ := cmd.Flags().GetBool("list-issues")
		listPRs, _ := cmd.Flags().GetBool("list-prs")
		listBranches, _ := cmd.Flags().GetBool("list-branches")
		commitHistory, _ := cmd.Flags().GetBool("commit-history")
		if len(args) == 1 {
			owner, repo = utils.DecomposeURL(args[0])
		} else if len(args) == 2 {
			owner, repo = args[0], args[1]
		}
		if owner != "" && repo != "" {
			stats, err := utils.RepoAnalyzer(owner, repo)
			if err != nil {
				log.Fatalf("Error analyzing: %v", err)
			}
			if listIssues {
				for _, issue := range stats.Issues {
					fmt.Printf("Issue: %s\n", issue.GetTitle())
					fmt.Printf("Issue State: %s\n", issue.GetState())
					fmt.Printf("Issue Author: %s\n", issue.GetUser().GetLogin())
					fmt.Println("-----")
				}
			} else if listPRs {
				for _, pr := range stats.PullRequests {
					fmt.Printf("Issue: %s\n", pr.GetTitle())
					fmt.Printf("Issue State: %s\n", pr.GetState())
					fmt.Printf("Issue Author: %s\n", pr.GetUser().GetLogin())
					fmt.Println("-----")
				}
			} else if listBranches {
				for _, branch := range stats.Branches {
					fmt.Printf("Branch Name: %s\n", branch.GetName())
					fmt.Println("-----")
				}
			} else if commitHistory {
				for _, commit := range stats.Commits {
					fmt.Printf("Author: %s\n", commit.GetAuthor().GetLogin())
					fmt.Printf("Date: %s\n", commit.Commit.GetAuthor().GetDate())
					fmt.Printf("Message: %s\n", commit.Commit.GetMessage())
					fmt.Println("-----")
				}
			} else {
				utils.PrintRepoStats(stats)
			}
		} else {
			fmt.Println("Unable to analyze repository.")
		}
	},
}

func init() {
	rootCmd.AddCommand(repositoryCmd)
}
