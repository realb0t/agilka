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

func LoadConfig(filePath string) *Config {
  jsonData, err := ioutil.ReadFile(filePath)

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

func ConfigPath(basePath, projectName string) string {
  return path.Join(basePath, projectName + ".json")
}

func (c *Config) Save(dirPath string) error {
  jsonStr, err := json.MarshalIndent(c, "", "  ")
  
  if err != nil {
    return err
  }

  return ioutil.WriteFile(
    ConfigPath(dirPath, c.Name), 
    jsonStr, 0644)
}