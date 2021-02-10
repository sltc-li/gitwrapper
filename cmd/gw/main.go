package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/urfave/cli"

	"github.com/li-go/gitwrapper"
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
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "fish",
			Usage: "generate fish completion",
		},
	}
	app.Commands = getCommands()
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		if c.Bool("fish") {
			completion, err := c.App.ToFishCompletion()
			if err != nil {
				return err
			}
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			completionFile := path.Join(home, ".config", "fish", "completions", "gw.fish")
			fmt.Printf("Installing to %s\n", completionFile)
			if err := ioutil.WriteFile(completionFile, []byte(completion), 0644); err != nil {
				return err
			}
			fmt.Println("Done!")
			return nil
		}
		return cli.ShowAppHelp(c)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCommands() cli.Commands {
	return cli.Commands{
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update current branch",
			Action: func(c *cli.Context) error {
				return getRepoConfig().Update()
			},
		},
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
			Aliases: []string{"ci"},
			Usage:   "commit changes",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "A",
					Usage: "add changes (include un-tracked files)",
				},
				cli.StringFlag{
					Name:  "am",
					Usage: "add changes and give a commit message",
				},
			},
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				if err := rc.Commit(c.Bool("A"), c.String("am")); err != nil {
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
				cli.BoolFlag{
					Name:  "f",
					Usage: "force push",
				},
			},
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				if err := rc.Push(c.Bool("f")); err != nil {
					return err
				}
				if c.Bool("g") {
					return rc.OpenCompare()
				}
				return nil
			},
		},
		{
			Name:    "github",
			Aliases: []string{"g"},
			Usage:   "open github",
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				return rc.OpenRepo()
			},
		},
		{
			Name:    "release",
			Aliases: []string{"r"},
			Usage:   "add release tag and push tags",
			Action: func(c *cli.Context) error {
				rc := getRepoConfig()
				if err := rc.AddReleaseTag(); err != nil {
					return err
				}
				return rc.PushTags()
			},
		},
	}
}
