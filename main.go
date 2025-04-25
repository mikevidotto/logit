package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"text/template"
	"time"
)

type Config struct {
	Date           string `json:"date"`
	CurrentProject string `json:"currentproject"`
}

func main() {
	config := LoadConfig()
	project := flag.String("p", "<projectname>", "Update the current project config.")
	flag.Parse()

	if *project != "" && *project != "<projectname>" {
		fmt.Println("setting current project to", *project)
		config.CurrentProject = *project
		UpdateConfig(config)
	}

	if len(os.Args) < 2 {

		templatepath := "C:/Users/jenna/GolandProjects/logit/template.md"
		tmplText, err := os.ReadFile(templatepath)
		if err != nil {
			fmt.Println("error reading contents of template.md", err)
		}

		path := "C:/Users/jenna/GolandProjects/logit/logs/" + time.Now().Local().String()[:10] + ".md"

		newtemplate := template.Must(template.New(time.Now().Local().String()[:10]).Parse(string(tmplText)))

		os.Mkdir("logs", 0666)
		_, err = os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Creating daily log file.")
			//os.WriteFile(path, tmplText, 0666)
			file, err := os.Create(path)
			if err != nil {
				fmt.Println("error creating file, err")
			}
			err = newtemplate.Execute(file, config)
		}

		OpenLog(path)
	}
}

func OpenLog(logpath string) {
	cmd := exec.Command("nvim", logpath)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func LoadConfig() Config {
	configData, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("error reading from config: ", err)
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("error unmarshaling config.json into data structure: ", err)
	}
    config.Date = time.Now().Local().String()[:10]
	return config
}

func UpdateConfig(config Config) {

	updated, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("error marshalling data structure into json: ", err)
	}
	os.WriteFile("./config.json", updated, 0666)
}
