package global

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var GlobalObject *Conf

type Conf struct {
	ApiKey string `yaml:"apiKey"`
}

func LoadApiKey(currentDir string) {
	globalPath := filepath.Join(currentDir,"cfg")
	confPath := filepath.Join(globalPath,"config.yml")
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &GlobalObject); err != nil {
		panic(err)
	}
	if len(GlobalObject.ApiKey) == 0 {
		panic("API Key missing")
	}
}