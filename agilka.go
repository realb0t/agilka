package main

import (
  "os"
  "fmt"
  "strings"
  "github.com/codegangsta/cli"
  "github.com/realb0t/agilka/project"
  "github.com/realb0t/agilka/task"
  "encoding/json"
  _ "github.com/deiwin/interact"
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
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("Error:", err)
    }
  }()

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
        pr := project.NewProject(projectName, projectPath, nil)
        if pr.IsExist() {
          panic("Project has been exist")
        }

        fmt.Println("Create project:", projectName)
        pr.Build()
        pr.Load()
      },
    },
    {
      Name:      "task",
      Usage:     "operations for ticker task",
      Subcommands: []cli.Command{
        {
          Name:    "create",
          Aliases: []string{"a"},
          Usage:   "create a new task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "json",
              Value: func() string {
                defaultTask := task.DefaultTask()
                json, _ := json.MarshalIndent(defaultTask, "", "  ")
                return string(json)
              }(),
              Usage: "JSON object for task",
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
            pairs := c.Args()
            jsonVal := c.String("json")
            projectPath := c.String("path")

            pr := project.LoadProject(projectPath)

            task := task.NewTask(jsonVal)
            task.ApplyPairs(pairs)
            task.Save(pr)

            fmt.Println("Created task with code", task.Code)
          },
        },
        {
          Name:    "edit",
          Aliases: []string{"c"},
          Usage:   "complete a task on the list",
          Action:  func(c *cli.Context) {
            objCode := c.Args().First()
            println("edit task by code: ", objCode)
          },
        },
      },
    },
  }

  app.Run(os.Args)
}