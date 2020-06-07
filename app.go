package main

import (
	"fmt"

	"github.com/urfave/cli"
	"gitlab.com/tbacompany/gitector/printer"
	"gitlab.com/tbacompany/gitector/reader"
	"gitlab.com/tbacompany/gitector/rules"
	"gitlab.com/tbacompany/gitector/utils"
)

// CreateNewApp creates a new CLI Applciation
func CreateNewApp() *cli.App {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: ".gitector",
			Usage: "Load configuration from `FILE`",
		},
		cli.StringFlag{
			Name:  "directory, d",
			Value: ".",
			Usage: "Directory to compare default .",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "console",
			Usage: "Use certain output console, markdown, html. Default console",
		},
		cli.IntFlag{
			Name:  "limit, l",
			Usage: "Limit check to `LIMIT` last commits",
			Value: 0,
		},
		cli.BoolFlag{
			Name:  "direct-input, di",
			Usage: "Use direct input ",
		},
	}

	app.Action = func(c *cli.Context) error {
		// gitScope is defined same way as git diff does it branch..branch or commit..commit
		gitScope := "master.."
		if c.NArg() > 0 {
			gitScope = c.Args().Get(0)
			if gitScope == "init" {
				utils.InitConfig(".gitector.json")
				return nil
			}
		}

		initParams := reader.InitParams{
			Scope:            gitScope,
			ConfigFilePath:   c.String("config"),
			Limit:            c.Int("limit"),
			Directory:        c.String("directory"),
			DirectInput:      gitScope,
			UsingDirectInput: c.Bool("direct-input"),
			Output:           c.String("output"),
		}
		commits := reader.ReadGitCommits(initParams)
		fmt.Printf("Running on directory: %#v comparing branches: %#v \n", initParams.Directory, gitScope)
		errors := rules.Rules(commits, initParams.Directory)
		fmt.Println(generateErrorOutput(initParams.Output, errors))
		if len(errors) > 0 {
			return cli.NewExitError("", 1)
		}
		return nil
	}

	return app
}

func generateErrorOutput(format string, errors []rules.GitError) string {
	switch format {
	case "html":
		return printer.GenerateHTML(errors)
	case "markdown":
		return printer.GenerateMarkdown(errors)
	case "json":
		return printer.GenerateJSON(errors)
	default:
		return printer.GenerateConsoleErrors(errors)
	}
}
