/*
Copyright © 2024 Indiance
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Indiance/go-gh-analyzer/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-analyzer",
	Short: "A command line interface that analyzes a github repository",
	Long: `This is a command line tool that allows you to analyze github repositories.
	If called without any commands, it will try to analyze the github repository which is associated.
	Probably done by running git remote -v and seeing if there is any github repository. Alternatively using other
	commands the cli can analyze github users, repositories, and organizations`,
	Run: func(cmd *cobra.Command, args []string) {
		owner, repo := utils.DecomposeGit()
		stats, err := utils.RepoAnalyzer(owner, repo)
		if err != nil {
			log.Fatalf("Error analyzing: %v", err)
		}
		listIssues, _ := cmd.Flags().GetBool("list-issues")
		listPRs, _ := cmd.Flags().GetBool("list-prs")
		listBranches, _ := cmd.Flags().GetBool("list-branches")
		commitHistory, _ := cmd.Flags().GetBool("commit-history")
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().Bool("list-issues", false, "List all issues in a repository")
	rootCmd.PersistentFlags().Bool("list-prs", false, "List all the pull requests in a repository")
	rootCmd.PersistentFlags().Bool("list-branches", false, "List all the branches in a repository")
	rootCmd.PersistentFlags().Bool("commit-history", false, "List the commit history of a repository")
}
