package main

import (
  //"fmt"
  //"encoding/json"
  "os"
  "fmt"
  "strings"
  //"github.com/deiwin/interact"
  "github.com/codegangsta/cli"
  "github.com/realb0t/agilka/project"
)

func parseFieldsFlags(c *cli.Context) map[string]string {
  flags := c.StringSlice("field")
  fields := make(map[string]string)
  for _, fieldVal := range flags {
    f := strings.Split(fieldVal, "=")
    fields[f[0]] = f[1]
  }

  return fields
}

func main() {
  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:    "init",
      Usage:   "Initialize new project in current dirrectory",
      Flags:   []cli.Flag {
        cli.StringFlag{
          Name: "name",
          Value: "AgilkaProject",
          Usage: "Project name",
        },
        cli.StringFlag{
          Name: "path",
          Value: func() string {
            projectPath, _ := os.Getwd()
            return projectPath
          }(),
          Usage: "Project current PATH",
        },
      },
      Action:  func(c *cli.Context) {
        projectName := c.String("name")
        projectPath := c.String("path")
        fmt.Println("Create project:", projectName)
        pr := project.NewProject(projectName, projectPath)
        if pr.IsExist() {
          panic("Project has been exist")
        }
        pr.Initialize()
      },
    },
    {
      Name:      "ticket",
      Usage:     "operations for ticker ticket",
      Subcommands: []cli.Command{
        {
          Name:    "create",
          Aliases: []string{"a"},
          Usage:   "create a new ticket",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "json",
              Value: "{}",
              Usage: "JSON object for ticket",
            },
          },
          Action:  func(c *cli.Context) {
            values := c.Args()
            jsonVal := c.String("json")
            fmt.Println("added object values: ", values)
            fmt.Println("added jsonVal: ", jsonVal)
          },
        },
        {
          Name:    "edit",
          Aliases: []string{"c"},
          Usage:   "complete a task on the list",
          Action:  func(c *cli.Context) {
            objCode := c.Args().First()
            println("edit ticket by code: ", objCode)
          },
        },
      },
    },
  }

  app.Run(os.Args)
}