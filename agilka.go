package main

import (
  //"fmt"
  //"encoding/json"
  "os"
  "fmt"
  "strings"
  //"github.com/deiwin/interact"
  "github.com/codegangsta/cli"
  "github.com/realb0t/agilka/obj"
  //"strings"
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
  obj.Init()

  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:    "new",
      Aliases: []string{"a"},
      Usage:   "add a task to the list",
      Flags:   []cli.Flag{
        cli.StringSliceFlag{
          Name: "field, f",
          Value: &cli.StringSlice{},
          Usage: "Fields defenitions",
        },
      },
      Action:  func(c *cli.Context) {
        fmt.Println("flags", parseFieldsFlags(c))
        objName := c.Args().First()
        o := obj.CreateByName(objName)
        println("added object: ", o.Marshal())
      },
    },
    {
      Name:    "complete",
      Aliases: []string{"c"},
      Usage:   "complete a task on the list",
      Action:  func(c *cli.Context) {
        println("completed task: ", c.Args().First())
      },
    },
  }

  app.Run(os.Args)
}