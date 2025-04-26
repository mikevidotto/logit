package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

const (
    configpath="C:/Users/jenna/GolandProjects/logit/config.json"
    logdirectory="C:/Users/jenna/GolandProjects/logs/"
    helpText=`Usage: logit [options]

A minimal daily logging tool for developers.

Options:
  -p [project]   Set the current project name in the config.
  -t             Set tasks you are working on for today's log.
  -l             List all existing log files.
  -h             Show this help message.

When run without any flags, logit creates or opens today's log file in your editor.
`


)

type Config struct {
	Date           string `json:"date"`
	CurrentProject string `json:"currentproject"`
	Tasks          []Task
}

type Task struct {
	Name string `json:"taskname"`
}

func main() {
	config := LoadConfig()
    help := flag.Bool("h", false, "see available commands for logit.")
	project := flag.String("p", "<projectname>", "Update the current project config.")
	tasks := flag.Bool("t", false, "Update your tasks for the day.")
	list := flag.Bool("l", false, "Displays a list of your log files.")
	flag.Parse()

    if *help {
        fmt.Print(helpText)
    }
	if *project != "" && *project != "<projectname>" {
		fmt.Println("setting current project to", *project)
		config.CurrentProject = *project
		UpdateConfig(config)
	}
	if *tasks {
		tasks, err := SetTasks(2)
		if err != nil {
			log.Fatal("error setting tasks: ", err)
		}
		config.Tasks = tasks
		UpdateConfig(config)
	}
	if *list {
		fmt.Println("All Logs:")
		ShowLogs()
	}

	if len(os.Args) < 2 {
		path := logdirectory + time.Now().Local().String()[:10] + ".md"
		CreateLog(path, config)
		OpenLog(path)
	}
}

func ShowLogs() {
    direntry, err := os.ReadDir(logdirectory)
    if err != nil {
        log.Fatal("error reading directory:", logdirectory, ":", err)
    }
    for i, entry := range direntry {
        fmt.Println(i+1, entry.Name())
    }
}

func GetTemplate() []byte {
	templatepath := "C:/Users/jenna/GolandProjects/logit/template.md"
	tmplText, err := os.ReadFile(templatepath)
	if err != nil {
		log.Fatal("error reading contents of template.md", err)
	}

	return tmplText
}

func CreateLog(path string, config Config) {
	todaysdate := time.Now().Local().String()[:10]

	_, err := os.Stat("logs")
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir("logs", 0666)
	}

	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		tmplText := GetTemplate()
		newlog := template.Must(template.New(todaysdate).Parse(string(tmplText)))

		fmt.Println("Creating daily log file.")
		//os.WriteFile(path, tmplText, 0666)
		file, err := os.Create(path)
		if err != nil {
			log.Fatal("error creating file, err")
		}
		err = newlog.Execute(file, config)
	} else {
		GetLogData()
	}
}

func GetLogData() {
	//should update the config based on what exists within the log file.
	//example: user enters tasks for the day, but NOT through the -t flag
	//      this would result in having a log file with tasks that do not sync up with the configuration
	//      config.json would show: task1: "task one!" and task2: "task two!"
	//      but....
	//      the log file would show: task1: "finish logit functinliaty" and task2: "update your resume dingo"
	//      we need to find a way to sync from the log file to the config.json file... so that we can make changes wtihin the file itself.
}

func SetTasks(numtasks int) ([]Task, error) {
	in := bufio.NewReader(os.Stdin)
	var tasks []Task
	var task Task
	for i := range numtasks {
		fmt.Print("Task ", i+1, ":")
		line, err := in.ReadString('\n')
		if err != nil {
			return nil, err
		}
        if len(line) > 2 {
            task.Name = line[:len(line)-2]
        }else {
            task.Name = line + "ERROR"
        }
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func OpenLog(logpath string) {
	cmd := exec.Command("nvim", logpath)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func LoadConfig() Config {
    configData, err := os.ReadFile("C:/Users/jenna/GolandProjects/logit/config.json")
	if err != nil {
		log.Fatal("error reading from config: ", err)
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal("error unmarshaling config.json into data structure: ", err)
	}
	config.Date = time.Now().Local().String()[:10]
	config.Tasks = []Task{
		{Name: "1"},
		{Name: "2"},
	}
	return config
}

func UpdateConfig(config Config) {
	updated, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("error marshalling data structure into json: ", err)
	}
	os.WriteFile(configpath, updated, 0666)
}
