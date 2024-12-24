package utils

import (
	"strings"
)

func DecomposeURL(url string) (owner string, repo string) {
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
