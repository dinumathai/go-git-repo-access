# Git repository access using golang

A sample program to access git repository access using golang program.


## How to run ?

1. Set the variables in [main.go](main.go)
```
	gitRepoURL := "https://github.com/dinumathai/go-git-repo-access.git"
	gitRepoUserName := "repo-writer"
	gitRepoPassword := "the-password"
	authorName := "My Name"
	authorEmail := "myemail@org.com"
	cloneDir := "./workdir/go-git-repo-access"
```
2. Run program
```
go mod tidy
go run main.go
```


## Reference
1. https://github.com/go-git/go-git/tree/master/_examples
1. https://pkg.go.dev/github.com/go-git/go-git/
