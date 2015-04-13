package project

import (
  "os"
  "path"
)

type Project struct {
  Name string
  Path string
  taskCount int
}

func NewProject(Name, Path string) *Project {
  return &Project{Name, Path, 0}
}

// Создает структуру папок для проекта
func (p *Project) build() {
  var err error
  var paths = []string{ path.Join(p.Path, "tasks"), 
    path.Join(p.Path, "attaches") }
  for _, path := range(paths) {
    err = os.MkdirAll(path, 0700)
    if err != nil {
      panic(err)
    }
  }
}

func (p *Project) taskPaths() []string {
  var paths []string
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
