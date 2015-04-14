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
  err = json.Unmarshal(jsonData, conf)

  if err != nil {
    panic(err)
  }

  return conf
}

func (c *Config) Save(dirPath string) error {
  jsonStr, err := json.MarshalIndent(c, "", "  ")
  
  if err != nil {
    panic(err)
  }

  configPath := path.Join(dirPath, c.Name + ".json")

  return ioutil.WriteFile(configPath, jsonStr, 0644)
}