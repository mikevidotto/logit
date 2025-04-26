package utils

import (
	"bufio"
	"fmt"
	"github.com/mikevidotto/logit/internal/config"
	"os"
)

func SetTasks(numtasks int) ([]config.Task, error) {
	in := bufio.NewReader(os.Stdin)
	var tasks []config.Task
	var task config.Task
	for i := range numtasks {
		fmt.Print("Task ", i+1, ":")
		line, err := in.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if len(line) > 2 {
			task.Name = line[:len(line)-2]
		} else {
			task.Name = line + "ERROR"
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
