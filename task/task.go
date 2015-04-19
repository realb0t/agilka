// Модель задачи. Базовые
package task

import (
  //"github.com/realb0t/agilka/project"
  "encoding/json"
  "reflect"
  "strings"
  "strconv"
  "io/ioutil"
  valid "github.com/asaskevich/govalidator"
)

// Структура задачи
type Task struct {
  Code string `json:"code" valid:"required,alphanum"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Author string `json:"author" valid:"alphanum"`
  State string `json:"state" valid:"task_state"`
}

func AvalibleStates() map[string]bool {
  return map[string]bool{
    "backlog": true, "todo": true,
    "doing": true, "done": true,
  }
}

// Объект дефолтной задачи
func DefaultTask() *Task {
  return &Task{ "", "", "", "", "backlog" }
}

// JSON от дефолтной задачи
func DefaultTaskJSON() string {
  jsonData, _ := DefaultTask().ToJSON()
  return string(jsonData)
}

// Создание нового эксземляра задачи
//
// Поддерживает следующие варианты:
// - fields string || []byte = JSON
// - fields []string = [ "field1=value1", "field2=value2" ]
func NewTask(fields interface{}) *Task {
  task := DefaultTask()

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
      task.ApplyPairs(f)
    default:
  }
  
  return task
}

// Применение пар ["ключ=значение"] к задаче
func (task *Task) ApplyPairs(pairs []string) {
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
}

// Запланировать задачу
func (t *Task) Plan() error {
  t.State = "todo"
  return nil
}

// Начать делать задачу
func (t *Task) Start() error {
  t.State = "doing"
  return nil
}

// Завершить задачу
func (t *Task) Done() error {
  t.State = "done"
  return nil
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

func (t *Task) validate() (bool, error) {
  valid.TagMap["task_state"] = valid.Validator(func(state string) bool {
    return AvalibleStates()[state]
  })

  return valid.ValidateStruct(t)
}

// Сохранить задачу для указанного проекта
func (t *Task) Save(taskPath string) error {
  _, err := t.validate()

  if err != nil {
    panic(err)
  }

  jsonStr, err := t.ToJSON()

  if err != nil {
    panic(err)
  }

  return ioutil.WriteFile(taskPath, jsonStr, 0644)
}

// Структура тикета (задача как файл)
type Ticket struct {
  Task *Task
  path string
}

func NewTicket(task *Task, path string) *Ticket {
  return &Ticket{task, path}
}

func LoadTicket(path string) *Ticket {
  jsonData, err := ioutil.ReadFile(path)

  if err != nil {
    panic(err)
  }

  return NewTicket(NewTask(jsonData), path)
}

func (t *Ticket) Save() error {
  return t.Task.Save(t.path)
}
