package project

import (
  "testing"
  "path"
  "os"
)

func TestNewProject(t *testing.T) {
  _ = NewProject("TestProject", "./testProject")
}

func TestInitialize(t *testing.T) {
  file, _      := os.Getwd()
  workPath     := path.Join(file, "..", "testProject")
  project      := NewProject("TestProject", workPath)
  tasksPath    := path.Join(workPath, "tasks")
  attachesPath := path.Join(workPath, "attaches")
  
  _ = project.Initialize()
  
  if _, err := os.Stat(tasksPath); os.IsNotExist(err) {
    t.Error("Not create Tasks dir")
  }
  
  if _, err := os.Stat(attachesPath); os.IsNotExist(err) {
    t.Error("Not create Attaches dir")
  }

  _ = os.RemoveAll(workPath)
}
