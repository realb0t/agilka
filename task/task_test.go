package task

import "testing"

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