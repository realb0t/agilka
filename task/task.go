package task

import (
  "encoding/json"
)

type Task struct {
  Code string `json:"code"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Author string `json:"author"`
  State string `json:"state"`
}

func MakeTask(fields map[string]string) Task {
  task := &Task{}
  return task
}

func MakeTaskByJSON(jsonData string) Task {
  var task Task
  err := json.Unmarshal(byte(jsonData), &task)
  return task, err
}

