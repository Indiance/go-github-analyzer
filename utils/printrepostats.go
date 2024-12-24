package utils

import (
	"fmt"
)

func PrintRepoStats(stats *RepoStats) {
	fmt.Println("%------------LANGUAGE COMPOSITION------------%")
	for _, langStat := range stats.LanguagePercentage {
		fmt.Printf("%s: %.2f%%\n", langStat.Language, langStat.Percentage)
	}
	fmt.Println("%------------DEVELOPMENT STATISTICS------------%")
	fmt.Printf("Number of commits: %d\n", len(stats.Commits))
	fmt.Printf("Number of issues: %d\n", len(stats.Issues))
	fmt.Printf("Number of pull requests: %d\n", len(stats.PullRequests))
	fmt.Printf("Number of branches: %d\n", len(stats.Branches))
	fmt.Println("%------------COMMUNITY STATISTICS------------%")
	fmt.Printf("Number of stars: %d\n", stats.Stars)
	fmt.Printf("Number of forks: %d\n", stats.Forks)
	fmt.Printf("Number of watchers: %d\n", stats.Watchers)
}
