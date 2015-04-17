package operation

import (
  "github.com/realb0t/agilka/task"
  "github.com/realb0t/agilka/project"
  "github.com/codegangsta/cli"
  //"io/ioutil"
  "os"
  //"fmt"
)

type Operation struct {
  ctx *cli.Context
}

// Создание операции
func NewOperation(c *cli.Context) *Operation {
  return &Operation{c}
}

// Создание нового проекта
func (o *Operation) CreateProject() *project.Project {
  pr := project.NewProject(o.ctx.String("name"),
    o.ctx.String("path"), nil)
  if pr.IsExist() {
    panic("Project has been exist")
  }
  pr.Build()
  return pr
}

// Создание новой задачи
func (o *Operation) CreateTask() *task.Task {
  pr := project.LoadProject(o.ctx.String("path"))
  task := task.NewTask(o.ctx.String("json"))
  task.ApplyPairs(o.ctx.Args())

  if task.Code == "" {
    task.Code = pr.NextTaskCode()
  }
  taskPath := pr.TaskPathByCode(task.Code)
  _, err := os.Stat(taskPath)
  if err == nil {
    panic("Task with code " + task.Code + " is exists")
  }

  _ = task.Save(taskPath)
  return task
}

