package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

const (
	shellToUse = "bash"
	quickPush  = `
		git add -A &&
		git commit -m "%v" &&
		git push;
	`
)

func task() {
	fmt.Println("golang cron task running ...")
	cmd := fmt.Sprintf(quickPush, time.Now())
	err, std, stderr := ShellCmd(cmd)
	fmt.Println(err, std, stderr)
}

func ShellCmd(command string) (error, string, string) {
	fmt.Println(command)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, task)
	c.Start()
	select {}
}
