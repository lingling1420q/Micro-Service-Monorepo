package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

func ShellCmd(command string) (string, string, error) {
	fmt.Println(command)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func task() {
	fmt.Println("golang cron task running ...")
	qpushCmd := fmt.Sprintf(quickPush, time.Now())
	touchCmd := fmt.Sprintf("touch %v", strings.Replace(fmt.Sprintf("%v", time.Now()), " ", "", -1))
	_err, _std, _stderr := ShellCmd(touchCmd)
	fmt.Println(_err, _std, _stderr)
	err, std, stderr := ShellCmd(qpushCmd)
	fmt.Println(err, std, stderr)
}

func main() {
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, task)
	c.Start()
	select {}
}
