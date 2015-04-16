package config

import (
  "path"
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  Name string `json:name`
  Repository string `json:repository`
}

func NewConfig(name, repository string) *Config {
  return &Config{name, repository}
}

func LoadConfig(basePath string) *Config {
  jsonData, err := ioutil.ReadFile(ConfigPath(basePath))

  if err != nil {
    panic(err)
  }

  var conf *Config
  err = json.Unmarshal(jsonData, &conf)

  if err != nil {
    panic(err)
  }

  return conf
}

func ConfigPath(basePath string) string {
  return path.Join(basePath, "Agilkafile")
}

func (c *Config) Save(dirPath string) error {
  jsonStr, err := json.MarshalIndent(c, "", "  ")
  
  if err != nil {
    return err
  }

  return ioutil.WriteFile(
    ConfigPath(dirPath), 
    jsonStr, 0644)
}