package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

var (
	Config *Conf
	configMutex sync.Mutex
)

//profile variables
type Conf struct {
	Trans struct {
		AppId string `yaml:"app_id"`
		Secret string `yaml:"secret"`
		TransApi string `yaml:"trans_api"`
	}
}


func GetConf() *Conf {
	if Config != nil {
		return Config
	}
	configMutex.Lock()
	defer configMutex.Unlock()
	if Config != nil {
		return Config
	}
    yamlFile, err := os.ReadFile("./resource/config.yaml")
    if err != nil {
        fmt.Println(err.Error())
    }
    err = yaml.Unmarshal(yamlFile, &Config)
    if err != nil {
        fmt.Println(err.Error())
    }
    return Config
}
