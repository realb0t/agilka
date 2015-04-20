package main

import (
  "os"
  "fmt"
  "strings"
  "github.com/codegangsta/cli"
  _ "github.com/realb0t/agilka/project"
  "github.com/realb0t/agilka/task"
  "github.com/realb0t/agilka/operation"
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

  currentPath := func() string {
    cPath, _ := os.Getwd()
    return cPath
  }

  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:  "init",
      Usage: "Initialize new project in current directory",
      Flags: []cli.Flag {
        cli.StringFlag{
          Name: "name",
          Value: "AgilkaProject",
          Usage: "Project name",
        },
        cli.StringFlag{
          Name: "path",
          Value: currentPath(),
          Usage: "Project current PATH",
        },
      },
      Action:  func(c *cli.Context) {
        pr := operation.NewOperation(c).CreateProject()
        fmt.Println("Create project:", pr.Name)
      },
    },
    {
      Name:  "task",
      Usage: "Operations for ticker task",
      Subcommands: []cli.Command{
        {
          Name:    "create",
          Aliases: []string{"c", "new", "add"},
          Usage:   "agilka task create [field=value] ... [fieldN=valueN] --json='{}'",
          Description: "Create new task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "json",
              Value: task.DefaultTaskJSON(),
              Usage: "JSON object for task",
            },
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action: func(c *cli.Context) {
            t := operation.NewOperation(c).CreateTask()
            fmt.Println("Created task with code", t.Code)
          },
        },
        {
          Name:    "edit",
          Aliases: []string{"e", "update", "change"},
          Usage:   "agilka task edit [taskCode] [field=value] ... [fieldN=valueN]",
          Description: "Edit exist task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action:  func(c *cli.Context) {
            t := operation.NewOperation(c).EditTask()
            println("Edit task with code", t.Code)
          },
        },
        {
          Name:    "plan",
          Usage:   "agilka task plan [taskCode]",
          Description: "Planned task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action:  func(c *cli.Context) {
            t := operation.NewOperation(c).DoTask("plan")
            println("Task", t.Code, "in todo")
          },
        },
        {
          Name:    "start",
          Usage:   "agilka task start [taskCode]",
          Description: "Start task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action:  func(c *cli.Context) {
            t := operation.NewOperation(c).DoTask("start")
            println("Task", t.Code, "in doing")
          },
        },
        {
          Name:    "done",
          Usage:   "agilka task done [taskCode]",
          Description: "Done task",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action:  func(c *cli.Context) {
            t := operation.NewOperation(c).DoTask("done")
            println("Task", t.Code, "in done")
          },
        },
        {
          Name:    "list",
          Usage:   "agilka task list [state1] ... [stateN]",
          Description: "Print all task in states",
          Flags:   []cli.Flag {
            cli.StringFlag{
              Name: "path",
              Value: currentPath(),
              Usage: "Project current PATH",
            },
          },
          Action:  func(c *cli.Context) {
            operation.NewOperation(c).PrintTasks()
          },
        },
      },
    },
  }

  app.Run(os.Args)
}