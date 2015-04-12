package operation

import (
  "github.com/realb0t/agilka/task"
  //"io/ioutil"
  //"os"
  //"fmt"
)

type Operation struct {
  task *task.Task 
}

func NewOperation(task *task.Task) *Operation {
  return &Operation{task}
}
