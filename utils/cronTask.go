package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)

func task() {
	fmt.Println("cron ...")
	now := time.Now()
	cmd := fmt.Sprintf(`
	    git add -A &&
		git commit -m "%v" &&
		git push;
	`, now)
	fmt.Println(cmd)
	err, std, stderr := ShellCmd(cmd)
	fmt.Println(err, std, stderr)
}

const shellToUse = "bash"

func ShellCmd(command string) (error, string, string) {
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
