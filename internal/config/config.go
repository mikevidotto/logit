package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const configpath = "C:/Users/jenna/GolandProjects/logit/config.json"

type Config struct {
	Date           string `json:"date"`
	CurrentProject string `json:"currentproject"`
	Tasks          []Task
}

type Task struct {
	Name string `json:"taskname"`
}

func LoadConfig() Config {
	configData, err := os.ReadFile("C:/Users/jenna/GolandProjects/logit/config.json")
	if err != nil {
		log.Fatal("error reading from config: ", err)
	}

	var conf Config
	err = json.Unmarshal(configData, &conf)
	if err != nil {
		log.Fatal("error unmarshaling config.json into data structure: ", err)
	}
	conf.Date = time.Now().Local().String()[:10]
	conf.Tasks = []Task{
		{Name: "1"},
		{Name: "2"},
	}
	return conf
}

func UpdateConfig(config Config) {
	updated, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("error marshalling data structure into json: ", err)
	}
	os.WriteFile(configpath, updated, 0666)
}
