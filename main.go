/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/Indiance/go-gh-analyzer/cmd"
import "github.com/Indiance/go-gh-analyzer/githubclient"
func main() {
	githubclient.InitClient()
	cmd.Execute()
}
