package config

import (
  "testing"
  _ "path"
  _ "os"
)

func TestNewConfig(t *testing.T) {
  _ = NewConfig("ProjectName", "git@project.repo/master")
}