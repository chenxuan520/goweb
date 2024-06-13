package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type StaticPath struct {
	RelativePath string `json:"relative_path"`
	RootPath     string `json:"root_path"`
}

type Server struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type Config struct {
	Server      Server       `json:"server"`
	StaticPaths []StaticPath `json:"static_paths"`
}

var (
	GlobalConfig *Config
)

func Init() {
	configFile := "./config/config.json"
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		log.Println(err)
		configFile = "config.json"
		data, err = ioutil.ReadFile("/config/" + configFile)
		if err != nil {
			log.Println("Read config error!")
			log.Panic(err)
			return
		}
	}

	config := &Config{}

	err = json.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		log.Panic(err)
		return
	}

	GlobalConfig = config
	log.Println("Config " + configFile + " loaded.")
}

func InitWithPath(configPath string) error {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}

	config := &Config{}

	err = json.Unmarshal(data, config)

	if err != nil {
		log.Println("Unmarshal config error!")
		return err
	}

	GlobalConfig = config
	log.Println("Config " + configPath + " loaded.")
	return nil
}
