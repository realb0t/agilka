package operation

import "testing"
import "github.com/realb0t/agilka/task"

func TestNewOperation(t *testing.T) {
  task := NewOperation(task.NewTask(false))
  if true {

  } else {
    t.Error("Not correct code", task)
  }
}