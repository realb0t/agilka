package obj

import (
  "strings"
  "encoding/json"
)

type IObject interface {
  Change(field string, value interface{}) IObject
  Destroy() IObject
  Save() IObject
  Marshal() string
}

type Task struct {
  Code string `json:"code"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Author string `json:"author"`
  State string `json:"state"`
}

var types = make(map[string]func()*Task)

func Init() {
  types["task"] = func() *Task {
    return &Task{}
  }
}

func CreateByName(objName string) IObject {
  nameParts := strings.Split(objName, ":")
  typeName, code := nameParts[0], nameParts[1]
  println("Create object with type (", typeName, ")", code)
  obj := &Task{ Code: code }
  return obj
}

func (o *Task) Change(field string, value interface{}) IObject {
  return o
} 

func (o *Task) Destroy() IObject {
  return o
}

func (o *Task) Save() IObject {
  return o
}

func (o *Task) Marshal() string {
  b, err := json.Marshal(o)
  if err != nil {
    panic(err)
  }
  return string(b)
}

func NameParse(objName string) (string, string) {
  nameParts := strings.Split(objName, ":")
  return nameParts[0], nameParts[1]
}

