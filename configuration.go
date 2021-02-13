package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

const logPathDirectory = "logs"

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

// CreateLogging is check directory then create log folder and file
func (c *Configuration) CreateLogging() {
	if _, err := os.Stat(logPathDirectory); os.IsNotExist(err) {
		fmt.Println("Log directory not exists")
		os.Mkdir(logPathDirectory, os.ModePerm)

		file, fileErr := os.OpenFile(
			filepath.Join(logPathDirectory, "logs.txt"), os.O_RDONLY|os.O_CREATE, 0644)

		if fileErr != nil {
			log.Fatalf("Error occrured at creating log file: %v", fileErr)
		}

		file.Close()
	}
}
