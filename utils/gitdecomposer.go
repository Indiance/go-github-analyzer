package utils

import (
	"os/exec"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func DecomposeGit() (string, string) {
	// Execute command git remote -v to check for github repository
	gitCmd := exec.Command("git", "remote", "-v")
	var out bytes.Buffer
	gitCmd.Stdout = &out
	err := gitCmd.Run()
	// check for errors
	if err != nil {
		exitError, ok := err.(*exec.ExitError)
		// Check if a git repo even exists
		if ok && exitError.ExitCode() == 128 {
			fmt.Println("Analysis could not be conducted since there was no github repository to analyze")
			return "", ""
		}
	}
	// Obtain the output string and format it for a github url
	output := out.String()
	re := regexp.MustCompile(`(https://github\.com/[^\s]+|git@github\.com:[^\s]+)`)
	matches := re.FindAllString(output, -1)
	// return nothing if there is no github url
	if len(matches) == 0 {
		fmt.Println("No GitHub URL found")
		return "", ""
	}
	// split url into owner and repository
	url := strings.TrimSuffix(matches[0], ".git")
	var owner, repo string
	if strings.HasPrefix(url, "https://") {
		splits := strings.Split(url, "/")
		owner, repo = splits[len(splits)-2], splits[len(splits)-1]
	} else if strings.HasPrefix(url, "git@github.com:") {
		url = strings.TrimPrefix(url, "git@github.com:")
		splits := strings.Split(url, "/")
		owner, repo = splits[0], splits[1]
	}
	return owner, repo
}
