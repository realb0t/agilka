package task

import (
  "encoding/json"
)

// Структура задачи
type Task struct {
  Code string `json:"code"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Author string `json:"author"`
  State string `json:"state"`
}

// Создание 
func makeTaskByPairs(pairs []string) Task {
  return &Task{}
}

// Создание нового эксземляра задачи
func MakeTask(fields interface{}) Task {
  var task Task

  switch t := v.(type) {
    case string:
      err := json.Unmarshal(byte(fields), &task)
    case []string:
      task = makeTaskByPairs(fields)
    default:
      task = &Task{}
  }
  
  return task
}

// Перевод задачи в формат JSON
func (t &Task) ToJSON() []byte, err {
  return json.MarshalIndent(t, "", "  ")
}