package main

import (
	"flag"
	"fmt"
	"github.com/mikevidotto/logit/internal/config"
	"github.com/mikevidotto/logit/internal/logger"
	"github.com/mikevidotto/logit/internal/utils"
	"log"
	"os"
)

const (
	helpText = `Usage: logit [options]

A minimal daily logging tool for developers.

Options:
  -p [project]   Set the current project name in the config.
  -t             Set tasks you are working on for today's log.
  -l             List all existing log files.
  -h             Show this help message.

When run without any flags, logit creates or opens today's log file in your editor.
`
)

func main() {
	conf := config.LoadConfig()

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
		conf.CurrentProject = *project
		config.UpdateConfig(conf)
	}
	if *tasks {
		tasks, err := utils.SetTasks(2)
		if err != nil {
			log.Fatal("error setting tasks: ", err)
		}
		conf.Tasks = tasks
		config.UpdateConfig(conf)
	}
	if *list {
		fmt.Println("All Logs:")
		logger.ShowLogs()
	}

	if len(os.Args) < 2 {
		logger.CreateLog(conf)
		logger.OpenLog()
	}
}
