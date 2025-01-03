package utils

import (
	"context"
	"fmt"

	"github.com/Indiance/go-gh-analyzer/githubclient"
	"github.com/google/go-github/v62/github"
)

type LanguageStats struct {
	Language   string
	Percentage float64
}

type RepoStats struct {
	Languages          map[string]int
	LanguagePercentage []LanguageStats
	Commits            []*github.RepositoryCommit
	Issues             []*github.Issue
	PullRequests       []*github.Issue
	Branches           []*github.Branch
	Stars              int
	Forks              int
	Watchers           int
}

func RepoAnalyzer(owner, repo string) (*RepoStats, error) {
	client := githubclient.GetClient()
	var languages map[string]int
	var commits []*github.RepositoryCommit
	var issues []*github.Issue
	var branches []*github.Branch
	var repoDetails *github.Repository
	var langErr, commitErr, issueErr, branchErr, repoErr error

	languages, _, langErr = client.Repositories.ListLanguages(context.Background(), owner, repo)

	commits, _, commitErr = client.Repositories.ListCommits(context.Background(), owner, repo, nil)

	issues, _, issueErr = client.Issues.ListByRepo(context.Background(), owner, repo, &github.IssueListByRepoOptions{State: "all"})

	branches, _, branchErr = client.Repositories.ListBranches(context.Background(), owner, repo, nil)

	repoDetails, _, repoErr = client.Repositories.Get(context.Background(), owner, repo)

	// Check for any errors after all goroutines have completed
	if langErr != nil || commitErr != nil || issueErr != nil || branchErr != nil || repoErr != nil {
		return nil, fmt.Errorf("error(s) occurred: %v %v %v %v %v", langErr, commitErr, issueErr, branchErr, repoErr)
	}

	// Calculate total number of bytes for all languages
	var totalBytes int
	for _, bytes := range languages {
		totalBytes += bytes
	}

	// Calculate the percentage for each language
	var languagePercents []LanguageStats
	for lang, bytes := range languages {
		percent := (float64(bytes) / float64(totalBytes)) * 100
		languagePercents = append(languagePercents, LanguageStats{
			Language:   lang,
			Percentage: percent,
		})
	}

	// Process the fetched issues to separate Issues and Pull Requests
	var onlyIssues, onlyPRs []*github.Issue
	for _, issue := range issues {
		if issue.PullRequestLinks == nil {
			onlyIssues = append(onlyIssues, issue)
		} else {
			onlyPRs = append(onlyPRs, issue)
		}
	}

	// Construct and return the RepoStats struct
	stats := &RepoStats{
		Languages:          languages,
		LanguagePercentage: languagePercents,
		Commits:            commits,
		Issues:             onlyIssues,
		PullRequests:       onlyPRs,
		Branches:           branches,
		Stars:              repoDetails.GetStargazersCount(),
		Forks:              repoDetails.GetForksCount(),
		Watchers:           repoDetails.GetWatchersCount(),
	}
	return stats, nil
}
