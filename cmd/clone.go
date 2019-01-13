package main

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/urfave/cli"
)

var cloneCmd = cli.Command{
	Name:        "clone",
	Usage:       "clone git repo into GOPATH",
	ArgsUsage:   "clone [package] [repo-url]",
	Description: "clone is used to setup go project in local",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "remote",
			Value: "origin",
			Usage: "set repo remote name",
		},
		cli.StringFlag{
			Name:  "upstream",
			Usage: "add upstream repository",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		var (
			ctx        = context.Background()
			pkg        = cliCtx.Args().Get(0)
			repoURL    = cliCtx.Args().Get(1)
			remoteName = cliCtx.String("remote")
			remoteURL  = cliCtx.String("upstream")
		)

		if pkg == "" {
			return errors.New("package must be provided")
		}

		if repoURL == "" {
			return errors.New("repo url must be provided")
		}

		gopath, err := resolveGOPATH()
		if err != nil {
			return err
		}

		targetDir := filepath.Join(gopath, "src", pkg)
		repo, err := cloneRepo(ctx, targetDir, repoURL, withRemoteName(remoteName))
		if err != nil {
			return err
		}

		if remoteURL != "" {
			if err := createRemote(ctx, repo, "upstream", remoteURL); err != nil {
				return err
			}
		}
		return nil
	},
}
