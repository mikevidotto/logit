package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
    templatepath := "C:/Users/jenna/GolandProjects/logit/template.md"
    template, err := os.ReadFile(templatepath)
	if err != nil {
		log.Fatal("error reading contents of template.md", err)
	}
	path := "./logs/" + time.Now().Local().String()[:10] + ".md"
	fmt.Println(path)
	os.Mkdir("logs", 0666)
	os.WriteFile(path, template, 0666)

	//use os.exec to open the log file using neovim.
	logpath := "C:/Users/jenna/GolandProjects/logit/logs/" + time.Now().Local().String()[:10] + ".md"
	//workingdir := os.Getwd
    fmt.Println(logpath)
	cmd := exec.Command("nvim", logpath)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
