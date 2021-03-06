package project

import (
  "os"
  "path"
  "io/ioutil"
  "github.com/realb0t/agilka/config"
  "strconv"
  _ "fmt"
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

func NewProject(name, projectPath string, conf *config.Config) *Project {
  tasksPath := path.Join(projectPath, "tasks")
  attachesPath := path.Join(projectPath, "attaches")
  return &Project{ name, projectPath, conf,
    tasksPath, attachesPath, 0, 0 }
}

// Загрузка проекта по пути
func LoadProject(configPath string) *Project {
  conf := config.LoadConfig(configPath)
  project := NewProject(conf.Name, configPath, conf)
  project.Load()
  return project
}

// Создает новый проект по казанному пути
func (p *Project) Build() *Project {

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
  return p
}

// Загружает данные по проекту
func (p *Project) Load() *Project {
  p.taskCount = len(p.objectsPaths(p.tasksPath))
  p.attachesCount = len(p.objectsPaths(p.attachesPath))
  if p.Config == nil {
    p.Config = config.LoadConfig(p.Path)
  }
  return p
}

func (p *Project) TaskPaths() []string {
  return p.objectsPaths(p.tasksPath)
}

// Возвращает пути к файлам из директории
func (p *Project) objectsPaths(objPath string) []string {
  infos, err := ioutil.ReadDir(objPath)
  if err != nil {
    panic(err)
  }
  paths := make([]string, 0)
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

func (p *Project) TaskPathByCode(code string) string {
  return path.Join(p.tasksPath, code + ".json")
}

// Возвращает количество задач в проекте
func (p *Project) TaskCount() int {
  return p.taskCount
}

// Возвращает следующий возможный код задачи
func (p *Project) NextTaskCode() string {
  return "task" + strconv.Itoa(p.taskCount + 1)
}

// Удаляет проект целиком с локальным репозиторием
func (p *Project) Destroy() (*Project, error) {
  var err error
  return p, err
}
