package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/liyy7/gitwrapper"
)

func getRepoConfig() *gitwrapper.RepoConfig {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rc, err := gitwrapper.NewRepoConfig(dir)
	if err != nil {
		log.Fatal(err)
	}
	return rc
}

func main() {
	app := cli.NewApp()
	app.Name = "gw"
	app.Usage = "a simple wrapper command for git"
	app.Version = "0.0.1"
	app.Commands = getCommands()
	app.EnableBashCompletion = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() cli.Commands {
	return cli.Commands{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new refresh branch",
			Action: func(c *cli.Context) error {
				return getRepoConfig().CreateBranch(c.Args().First())
			},
		},
		{
			Name:    "commit",
			Aliases: []string{"cm"},
			Usage:   "commit changes",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "a",
					Usage: "add changes",
				},
				cli.StringFlag{
					Name:  "m",
					Usage: "give a commit message",
				},
			},
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				if err := rc.Commit(c.Bool("a"), c.String("m")); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "push",
			Aliases: []string{"p"},
			Usage:   "push to remote",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "g",
					Usage: "open compare page on github",
				},
			},
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				if err := rc.Push(); err != nil {
					return err
				}
				if c.Bool("g") {
					return rc.OpenCompareURL()
				}
				return nil
			},
		},
	}
}
