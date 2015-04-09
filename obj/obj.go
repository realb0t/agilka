package obj

import (
  "strings"
  //"encoding/json"
)

type IObject interface {
  Change(field string, value interface{}) interface{}
  Remove() interface{}
  Save() interface{}
}

type AbstractObject struct {
  Code string `json:"code"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Author string `json:"author"`
  State string `json:"state"`
}

type Task struct {
  AbstractObject
}

var types = make(map[string]IObject)

func Init() {
  types["task"] = new(Task)
}

func Create() IObject {
  obj := new(Task)
  return obj
}

func (o *AbstractObject) Change(field string, value interface{}) interface{} {
  return o
} 

func (o *AbstractObject) Remove() interface{} {
  return o
}

func (o *AbstractObject) Save() interface{} {
  return o
}

func NameParse(objName string) (string, string) {
  nameParts := strings.Split(objName, ":")
  return nameParts[0], nameParts[1]
}

