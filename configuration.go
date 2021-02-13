package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration represents yaml file
type Configuration struct {
	Providers []struct {
		Name         string `yaml:"name"`
		Token        string `yaml:"token"`
		Repositories []struct {
			Name                      string `yaml:"name"`
			URL                       string `yaml:"url"`
			SynchronizationRepeatTime int    `yaml:"synchronization_repeat_time"`
		} `yaml:"repositories"`
	} `yaml:"providers"`
}

// Initialize parse configuration file content into configuration struct
func (c *Configuration) Initialize() {
	config, configError := ioutil.ReadFile("configuration.yaml")
	if configError != nil {
		log.Fatalf("Error occrured at configuration file: %v", configError)
	}

	config = []byte(os.ExpandEnv(string(config)))

	err := yaml.Unmarshal(config, &c)

	if err != nil {
		log.Fatalf("Error occured: %v", err)
	}
}
