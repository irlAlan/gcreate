package net

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5" // with go modules disabled
	"github.com/irlalan/gcreate/internal/config"
	"github.com/irlalan/gcreate/internal/dir"
	"time"
	// "gopkg.in/src-d/go-git.v4/storage/memory"
)

func HandlePkgs(pkgs []config.Pkg) {
	fmt.Println("Handling packages")
	for _, pkg := range pkgs {
		fmt.Println("repo using goroutine")
		go CloneGitRepo(pkg)
	}
	time.Sleep(time.Second * 1)
}

func CloneGitRepo(pkg config.Pkg) {
	fmt.Println("Cloning repo: ", pkg.Name)
	_, err := git.PlainClone(dir.GetHomeConfigDir()+"/pkgs/"+pkg.Name+"_"+pkg.Tag,
		false, &git.CloneOptions{
			URL:          pkg.GitUrl,
			SingleBranch: true,
			Progress:     os.Stdout,
		})
	if err != nil {
		fmt.Println("Cannot clone: ", err)
	}
}
