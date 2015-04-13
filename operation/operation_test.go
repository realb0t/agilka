package operation

import "testing"
import "github.com/realb0t/agilka/task"

func TestNewOperation(t *testing.T) {
  _ = NewOperation(task.NewTask(false))
}