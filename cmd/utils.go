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

type cloneOpt func(*git.CloneOptions) error

func withRemoteName(name string) cloneOpt {
	return func(opt *git.CloneOptions) error {
		opt.RemoteName = name
		return nil
	}
}

func createRemote(ctx context.Context, repo *git.Repository, name, url string) error {
	_, err := repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	return err
}

// cloneRepo clones remote repo into local path.
func cloneRepo(ctx context.Context, path string, repoURL string, opts ...cloneOpt) (*git.Repository, error) {
	copt := &git.CloneOptions{
		URL: repoURL,
	}

	for _, o := range opts {
		if err := o(copt); err != nil {
			return nil, err
		}
	}
	return git.PlainCloneContext(ctx, path, false, copt)
}

// resolveGOPATH gets the first GOPATH.
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
