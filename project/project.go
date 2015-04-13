package project

import (
  "os"
  "path"
  "io/ioutil"
)

type Project struct {
  Name string
  Path string
  TasksPath string
  AttachesPath string
  taskCount int
}

func NewProject(Name, Path string) *Project {
  tasksPath := path.Join(Path, "tasks")
  attachesPath := path.Join(Path, "attaches")
  return &Project{ Name, Path,
    tasksPath, attachesPath, 0 }
}

// Создает структуру папок для проекта
func (p *Project) build() {
  var err error
  var paths = []string{ p.TasksPath, 
    p.AttachesPath }
  for _, path := range(paths) {
    err = os.MkdirAll(path, 0700)
    if err != nil {
      panic(err)
    }
  }
}

func (p *Project) taskPaths() []string {
  infos, err := ioutil.ReadDir(p.TasksPath)
  if err != nil {
    panic(err)
  }
  paths := make([]string, len(infos))
  for _, info := range(infos) {
    aPath := path.Join(p.TasksPath, info.Name())
    paths = append(paths, aPath)
  }
  return paths
}

func (p *Project) isExist() bool {
  _, err := os.Stat(p.Path)
  if err != nil {
    return !os.IsNotExist(err)
  } else {
    return true
  }
}

// Производит инициализацию проекта
func (p *Project) Initialize() (error) {
  var err error

  if !p.isExist() {
    p.build()
  } else {
    p.taskCount = len(p.taskPaths())
  }

  return err
}

// Возвращает количество задач в проекте
func (p *Project) TaskCount() int {
  return p.taskCount
}

// Возвращает следующий возможный код задачи
func (p *Project) NextTaskCode() string {
  return "task0001"
}

// Удаляет проект целиком с локальным репозиторием
func (p *Project) Destroy() (*Project, error) {
  var err error
  return p, err
}
