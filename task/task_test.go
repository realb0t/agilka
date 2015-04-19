package task

import (
  "testing"
)

const testCode = `testTicket`

func TestTaskMakeByPairs(t *testing.T) {
  pairs := []string{ 
    "code=" + testCode,
    "title=Title" +  
    "bar=baz" }
  task := NewTask(pairs)

  if false && task.Code != testCode {
    t.Error("Not correct code")
  }
}

func TestTaskMakeByJSON(t *testing.T) {
  task := NewTask(`{ 
    "code": "` + testCode + `",  
    "title": "Test ticket title",
    "desc": "",
    "author": "",
    "state": ""
  }`)

  if task.Code != testCode {
    t.Error("Not correct code")
  }
}

func TestTaskToJson(t *testing.T) {
  pairs := []string{ "code=" + testCode }
  task := NewTask(pairs)
  jsonData, _ := task.ToJSON()
  newTask := NewTask(jsonData)
  if newTask.Code != testCode {
    t.Error("Not correct code")
  }
}

func TestTaskApplyingDefaultCode(t *testing.T) {
  pairs := []string{ "title=Task title" }
  task := NewTask(pairs)
  task.ApplyDefaultCode(testCode)
  if task.Code != testCode {
    t.Error("Not correct code")
  }
}

func TestTaskDontApplyingDefaultCode(t *testing.T) {
  pairs := []string{ "title=Task title", "code=" + testCode }
  task := NewTask(pairs)
  task.ApplyDefaultCode("otherCodeByDefault")
  if task.Code != testCode {
    t.Error("Not correct code")
  }
}

func TestValidateCorrectCode(t *testing.T) {
  task := NewTask(nil)
  task.Code = "testCode123"
  _, err := task.Validate()
  if err != nil {
    t.Error("Not correct code validation")
  }
}

func TestValidateUncorrectCode(t *testing.T) {
  task := NewTask(nil)
  task.Code = "this is uncorrect_code *"
  _, err := task.Validate()
  if err == nil {
    t.Error("Not error with uncorrect code")
  }
}

func TestAvalibleState(t *testing.T) {
  task := NewTask(`{ "code": "testCode" }`)
  states := []string{ "backlog", "todo", "doing", "done" }
  for _, state := range(states) {
    task.State = state
    _, err := task.Validate()

    if err != nil {
      t.Error("Not valid state", state)
    }
  }
}

func TestWithUncorrectState(t *testing.T) {
  task := NewTask(`{ "code": "testCode" }`)
  task.State = "uncorrectState"
  _, err := task.Validate()

  if err == nil {
    t.Error("Not validate uncorrect state")
  }
}

func TestPlan(t *testing.T) {
  task := NewTask(nil)
  errTrans := task.Plan()
  if errTrans != nil {
    t.Error("Plan transition with error")
  }
  if task.State != "todo" {
    t.Error("Uncorrect TODO state")
  }
}

func TestStart(t *testing.T) {
  task := NewTask(nil)
  _ = task.Plan()
  errTrans := task.Start()
  if errTrans != nil {
    t.Error("Start transition with error")
  }
  if task.State != "doing" {
    t.Error("Uncorrect DOING state")
  }
}

func TestDone(t *testing.T) {
  task := NewTask(nil)
  _ = task.Plan()
  _ = task.Start()
  errTrans := task.Done()
  if errTrans != nil {
    t.Error("Plan transition with error")
  }
  if task.State != "done" {
    t.Error("Uncorrect DONE state")
  }
}
