package logger

import (
	"errors"
	"fmt"
	"github.com/mikevidotto/logit/internal/config"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

const logdirectory = "C:/Users/jenna/GolandProjects/logs/"

func CreateLog(conf config.Config) {
	path := logdirectory + time.Now().Local().String()[:10] + ".md"
	todaysdate := time.Now().Local().String()[:10]

	_, err := os.Stat(logdirectory)
	if errors.Is(err, os.ErrNotExist) {
		os.Mkdir(logdirectory, 0666)
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
		err = newlog.Execute(file, conf)
        if err != nil {
            log.Fatal("error executing template: ", err)
        }
	} else {
		GetLogData()
	}
}

func OpenLog() {
	path := logdirectory + time.Now().Local().String()[:10] + ".md"
	cmd := exec.Command("nvim", path)
	cmd.Stdout = os.Stdout
	cmd.Run()
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

func GetLogData() {
	//should update the config based on what exists within the log file.
	//example: user enters tasks for the day, but NOT through the -t flag
	//      this would result in having a log file with tasks that do not sync up with the configuration
	//      config.json would show: task1: "task one!" and task2: "task two!"
	//      but....
	//      the log file would show: task1: "finish logit functinliaty" and task2: "update your resume dingo"
	//      we need to find a way to sync from the log file to the config.json file... so that we can make changes wtihin the file itself.
}
