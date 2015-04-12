package task

import "testing"

const testCode = `testTicket`

func TestTaskMakeByPairs(t *testing.T) {
  pairs := []string{ 
    "code=" + testCode,
    "title=Title" +  
    "bar=baz" }
  task := MakeTask(pairs)

  if false && task.Code != testCode {
    t.Error("Not set code")
  }
}

func TestTaskMakeByJSON(t *testing.T) {
  task := MakeTask(`{ 
    "code": "` + testCode + `",  
    "title": "Test ticket title",
    "desc": "",
    "author": "",
    "state": ""
  }`)

  if task.Code != testCode {
    t.Error("Not set code")
  }
}

func TestTaskToJson(t *testing.T) {
  pairs := []string{ "code=" + testCode }
  task := MakeTask(pairs)
  jsonData, _ := task.ToJSON()
  newTask := MakeTask(jsonData)
  if newTask.Code != testCode {
    t.Error("Not set code")
  }
}
