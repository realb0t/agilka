package project

import (
  "os"
  "path"
  "io/ioutil"
  "github.com/realb0t/agilka/config"
)

type Project struct {
  Name string
  Path string
  Config *config.Config
  tasksPath string
  attachesPath string
  taskCount int
  attachesCount int
}

func NewProject(Name, Path string) *Project {
  tasksPath := path.Join(Path, "tasks")
  attachesPath := path.Join(Path, "attaches")
  return &Project{ Name, Path, nil,
    tasksPath, attachesPath, 0, 0 }
}

// Создает новый проект по казанному пути
func (p *Project) build() {

  // Создает папки
  var paths = []string{ p.tasksPath, 
    p.attachesPath }
  for _, path := range(paths) {
    err := os.MkdirAll(path, 0700)
    if err != nil {
      panic(err)
    }
  }

  // Создает конфигурационный файл
  p.Config = config.NewConfig(p.Name, "")
  p.Config.Save(p.Path)
}

// Загружает существующий проект
func (p *Project) load() {
  p.taskCount = len(p.objectsPaths(p.tasksPath))
  p.attachesCount = len(p.objectsPaths(p.attachesPath))
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
func (p *Project) IsExist() bool {
  _, err := os.Stat(p.tasksPath)
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

  if !p.IsExist() {
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
