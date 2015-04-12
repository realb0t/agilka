package task

import (
  "encoding/json"
  "reflect"
  "strings"
  "strconv"
  //"fmt"
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
func makeTaskByPairs(pairs []string) *Task {
  task := &Task{}

  for _, spair := range(pairs) {
    apair := strings.Split(spair, "=")
    name, val := apair[0], apair[1]
    fieldName := strings.Title(name)
    field := reflect.ValueOf(task).Elem().FieldByName(fieldName)
    if field.IsValid() && field.CanSet() {
      switch field.Kind() {
        case reflect.Int:
          ival, _ := strconv.ParseInt(val, 0, 64)
          field.SetInt(ival)
        case reflect.Float32:
          fval, _ := strconv.ParseFloat(val, 32)
          field.SetFloat(fval)
        case reflect.Float64:
          fval, _ := strconv.ParseFloat(val, 64)
          field.SetFloat(fval)
        case reflect.String:
          field.SetString(val)
        default:
          panic("Not support type for: " + spair)
      }
    }
  }

  return task
}

// Создание нового эксземляра задачи
//
// Поддерживает следующие варианты:
// - fields string || []byte = JSON
// - fields []string = [ "field1=value1", "field2=value2" ]
func NewTask(fields interface{}) *Task {
  var task *Task

  switch f := fields.(type) {
    case []byte:
      err := json.Unmarshal(f, &task)
      if err != nil {
        panic(err)
      }
    case string:
      err := json.Unmarshal([]byte(f), &task)
      if err != nil {
        panic(err)
      }
    case []string:
      task = makeTaskByPairs(f)
    default:
      task = &Task{}
  }
  
  return task
}

// Перевод задачи в формат JSON
func (t *Task) ToJSON() ([]byte, error) {
  return json.MarshalIndent(t, "", "  ")
}

// Назначает поле code по-умолчанию, если это требуется
func (t *Task) ApplyDefaultCode(code string) *Task {
  if t.Code == "" {
    t.Code = code
  }

  return t
}