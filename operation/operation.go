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

// Загрузить тикет
func (o *Operation) loadTicket() *task.Ticket {
  pr := project.LoadProject(o.ctx.String("path"))
  taskCode := o.ctx.Args().First()
  ticketPath := pr.TaskPathByCode(taskCode)
  return task.LoadTicket(ticketPath)
}

// Езменение задачи
func (o *Operation) EditTask() *task.Task {
  ticket := o.loadTicket()
  ticket.Task.ApplyPairs(o.ctx.Args().Tail())
  _ = ticket.Save()
  return ticket.Task
}

func (o *Operation) DoTask(action string) *task.Task {
  var err error
  ticket := o.loadTicket()

  switch action {
    case "plan": err = ticket.Task.Plan()
    case "start": err = ticket.Task.Start()
    case "done": err = ticket.Task.Done()
  }

  if err != nil {
    panic(err)
  }

  _ = ticket.Save()
  return ticket.Task
}
