package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

func createRemote(ctx context.Context, repo *git.Repository, name, url string) error {
	_, err := repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	return err
}

func cloneRepo(ctx context.Context, path string, repoURL string) (*git.Repository, error) {
	return git.PlainCloneContext(ctx, path, false, &git.CloneOptions{
		URL: repoURL,
	})
}

func resolveGOPATH() (string, error) {
	gopaths := os.Getenv("GOPATH")
	if gopaths == "" {
		return "", errors.New("missing GOPATH")
	}

	paths := filepath.SplitList(gopaths)
	if len(paths) > 1 {
		fmt.Fprintf(os.Stderr, "the GOPATH contains multiple entries, get uses the first one %s", paths[0])
	}
	return paths[0], nil
}
