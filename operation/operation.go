package operation

import (
  "github.com/realb0t/agilka/task"
  "github.com/realb0t/agilka/project"
  "github.com/codegangsta/cli"
  "fmt"
  "strings"
  "github.com/labstack/gommon/color"
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
  tsk := task.NewTask(o.ctx.String("json"))
  tsk.ApplyPairs(o.ctx.Args())
  tsk.ApplyDefaultCode(pr.NextTaskCode())

  taskPath := pr.TaskPathByCode(tsk.Code)
  ticket := task.NewTicket(tsk, taskPath)
  if ticket.IsExist() {
    panic("Task with code " + ticket.Task.Code + " is exists")
  }

  _ = ticket.Save()
  return ticket.Task
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

// Выполнить с задачей действие
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

func (o *Operation) PrintTasks() {
  var useStates map[string]bool
  ticketsByState := map[string][]*task.Ticket{
    "backlog": []*task.Ticket{},
    "todo": []*task.Ticket{},
    "doing": []*task.Ticket{},
    "done": []*task.Ticket{},
  }

  pr := project.LoadProject(o.ctx.String("path"))
  states := o.ctx.Args()

  if len(states) == 0 {
    useStates = task.AvalibleStates()
  } else {
    useStates = make(map[string]bool)
    for _, state := range(states) {
      useStates[state] = true
    }
  }

  for _, taskPath := range(pr.TaskPaths()) {
    ticket := task.LoadTicket(taskPath)
    if useStates[ticket.Task.State] {
      ticketsByState[ticket.Task.State] = append(ticketsByState[ticket.Task.State], ticket)
    }
  }

  maxTaskCodeLen := 10
  for _, tickets := range(ticketsByState) {
    for _, ticket := range(tickets) {
      if (maxTaskCodeLen < len(ticket.Task.Code)) {
        maxTaskCodeLen = len(ticket.Task.Code)
      }
    }
  }

  formatStr := `%-` + fmt.Sprintf("%d", maxTaskCodeLen) + "s | %s\n"
  startBorder := strings.Repeat("-", maxTaskCodeLen - 3)
  endBorder := strings.Repeat("-", 20)
  fmt.Println(color.Bold(startBorder + endBorder))
  fmt.Printf(color.Bold(formatStr), "CODE:", "TITLE:")
  fmt.Println(color.Bold(startBorder + endBorder))
  for state, tickets := range(ticketsByState) {
    if len(tickets) > 0 {
      fmt.Print(color.Dim(startBorder))
      fmt.Printf(color.Bold(" %-7s "), strings.ToUpper(state))
      fmt.Println(color.Dim(strings.Repeat("-", 20 - 9)))

      for _, ticket := range(tickets) {
        fmt.Printf(formatStr, ticket.Task.Code, ticket.Task.Title)
      }
    }
  }
}
