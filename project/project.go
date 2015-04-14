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
  attachesCount int
}

func NewProject(Name, Path string) *Project {
  tasksPath := path.Join(Path, "tasks")
  attachesPath := path.Join(Path, "attaches")
  return &Project{ Name, Path,
    tasksPath, attachesPath, 0, 0 }
}

// Создает новый проект по казанному пути
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

// Загружает существующий проект
func (p *Project) load() {
  p.taskCount = len(p.objectsPaths(p.TasksPath))
  p.attachesCount = len(p.objectsPaths(p.AttachesPath))
}

// Возвращает пути к файлам из директории
func (p *Project) objectsPaths(objPath string) []string {
  infos, err := ioutil.ReadDir(objPath)
  if err != nil {
    panic(err)
  }
  paths := make([]string, len(infos))
  for _, info := range(infos) {
    aPath := path.Join(objPath, info.Name())
    paths = append(paths, aPath)
  }
  return paths
}

// Проверяет наличие папки проекта
func (p *Project) isExist() bool {
  _, err := os.Stat(p.Path)
  if err != nil {
    return !os.IsNotExist(err)
  } else {
    return true
  }
}


// Производит создание проекта
// на пустой директории
//
// И загружает созданный или ранее
// существующий проект
func (p *Project) Initialize() (error) {
  var err error

  if !p.isExist() {
    p.build()
  }

  p.load()

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
