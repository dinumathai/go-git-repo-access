package main

import (
	"fmt"
	"os"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func main() {
	gitRepoURL := "https://github.com/dinumathai/go-git-repo-access.git"
	gitRepoUserName := "repo-writer"
	gitRepoPassword := "the-password"
	authorName := "My Name"
	authorEmail := "myemail@org.com"
	cloneDir := "./workdir/go-git-repo-access"

	err := os.RemoveAll(cloneDir)
	if err != nil {
		fmt.Printf("Delete dir failed. Error = %v\n", err)
	} else {
		fmt.Println("Deleted the dir")
	}
	repo, err := git.PlainClone(cloneDir, false,
		&git.CloneOptions{
			URL:      gitRepoURL,
			Progress: os.Stdout,
			Auth: &http.BasicAuth{
				Username: gitRepoUserName,
				Password: gitRepoPassword,
			},
			Depth: 1,
		})

	if err != nil {
		fmt.Printf("Clone Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfull cloned")
	}

	err = os.Mkdir(cloneDir+"/test", 0755)
	if err != nil {
		fmt.Printf("Create dir failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Created the directory - test")
	}

	d1 := []byte("hello\ngo\n")
	err = os.WriteFile(cloneDir+"/test/test.txt", d1, 0644)
	if err != nil {
		fmt.Printf("Create file failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Created the file - test/test.txt")
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Printf("Getting work Tree failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Getting work Tree success !!")
	}
	err = w.AddWithOptions(&git.AddOptions{
		All: true,
	})
	if err != nil {
		fmt.Printf("Git Add failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Git Add success !!")
	}

	status, err := w.Status()
	if err != nil {
		fmt.Printf("Git Status failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Git Status success !!")
	}
	fmt.Println(status)

	_, err = w.Commit("fix: Test message", &git.CommitOptions{
		Author: &object.Signature{
			Name:  authorName,
			Email: authorEmail,
			When:  time.Now(),
		},
	})
	if err != nil {
		fmt.Printf("Git Commit failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Git Commit success !!")
	}

	err = repo.Push(&git.PushOptions{
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: gitRepoUserName,
			Password: gitRepoPassword,
		},
	})
	if err != nil {
		fmt.Printf("Push failed. Error = %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Push Success !! The End !!")
	}
}
