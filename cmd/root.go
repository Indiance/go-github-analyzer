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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-gh-analyzer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
