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
	Run: func(cmd *cobra.Command, args []string) {
		var owner, repo string
		if len(args) == 1 {
			owner, repo = utils.DecomposeGit()
		} else {
			owner, repo = args[0], args[1]
		}
		if owner != "" && repo != "" {
			stats, err := utils.RepoAnalyzer(owner, repo)
			if err != nil {
				log.Fatalf("Error analyzing: %v", err)
			}
			fmt.Println("%------------LANGUAGE COMPOSITION------------%")
			for language, bytes := range stats.Languages {
				fmt.Println(language, "with bytes: ", bytes)
			}
			fmt.Println("%------------COMMIT HISTORY------------%")
			for _, commit := range stats.Commits {
				fmt.Printf("Author: %s\n", commit.GetAuthor().GetLogin())
				fmt.Printf("Date: %s\n", commit.Commit.GetAuthor().GetDate())
				fmt.Printf("Message: %s\n", commit.Commit.GetMessage())
				fmt.Println("-----")
			}
			fmt.Println("%------------ISSUE STATISTICS------------%")
			for _, issue := range stats.Issues {
				fmt.Printf("Issue: %s\n", issue.GetTitle())
				fmt.Printf("Issue State: %s\n", issue.GetState())
				fmt.Printf("Issue Author: %s\n", issue.GetUser().GetLogin())
				fmt.Println("-----")
			}
			fmt.Println("%------------PULL REQUEST STATISTICS------------%")
			for _, pr := range stats.PullRequests {
				fmt.Printf("Issue: %s\n", pr.GetTitle())
				fmt.Printf("Issue State: %s\n", pr.GetState())
				fmt.Printf("Issue Author: %s\n", pr.GetUser().GetLogin())
				fmt.Println("-----")
			}
			fmt.Println("%------------BRANCH STATISTICS------------%")
			for _, branch := range stats.Branches {
				fmt.Printf("Branch Name: %s\n", branch.GetName())
				fmt.Println("-----")
			}
			fmt.Println("%------------COMMUNITY STATISTICS------------%")
			fmt.Printf("Number of stars: %d\n", stats.Stars)
			fmt.Printf("Number of forks: %d\n", stats.Forks)
			fmt.Printf("Number of watchers: %d\n", stats.Watchers)
			fmt.Println("-----")
		}

	},
}

func init() {
	rootCmd.AddCommand(repositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
