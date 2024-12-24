package utils

import (
	"bytes"
	"fmt"
	"os/exec"
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
	owner, repo = DecomposeURL(url)
	return owner, repo
}
