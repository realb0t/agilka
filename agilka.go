package main

import (
  //"fmt"
  //"encoding/json"
  "os"
  //"github.com/deiwin/interact"
  "github.com/codegangsta/cli"
  "github.com/realb0t/agilka/obj"
  //"strings"
)

func main() {
  obj.Init()

  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:    "new",
      Aliases: []string{"a"},
      Usage:   "add a task to the list",
      Action:  func(c *cli.Context) {
        objName := c.Args().First()
        objectType, objectCode := obj.NameParse(objName)
        println("added object: ", objectType, objectCode)
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