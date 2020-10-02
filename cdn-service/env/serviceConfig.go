package env

import (
	"io/ioutil"
	"log"

	"go.kicksware.com/api/service-common/config"
	"gopkg.in/yaml.v2"

)

type ServiceConfig struct {
	Common   config.CommonConfig    `yaml:"commonConfig"`
	Security config.SecurityConfig  `yaml:"securityConfig"`
	Auth     config.AuthConfig      `yaml:"authConfig"`
	Files    config.DataStoreConfig `yaml:"filesConfig"`
	Mongo    config.DataStoreConfig `yaml:"mongoConfig"`
	Redis    config.DataStoreConfig `yaml:"redisConfig"`
}

func ReadServiceConfig(filename string) (sc ServiceConfig, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = yaml.Unmarshal(file, &sc)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
