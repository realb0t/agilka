package project

type Project struct {
  Name string
  Path string
  taskCount int
}

func NewProject(Name, Path string) *Project {
  return &Project{Name, Path, 0}
}

// Создает структуру папок для проекта
func (p *Project) build() (*Project, error) {
  var err error
  return p, err
}

func (p *Project) taskPaths() []string {
  var paths []string
  return paths
}

func (p *Project) isExist() bool {
  return false
}

// Производит инициализацию проекта
func (p *Project) Initialize() (*Project, error) {
  var err error

  if !p.isExist() {
    _, err = p.build()
  } else {
    p.taskCount = len(p.taskPaths())
  }

  return p, err
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
