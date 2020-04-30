package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

var ErrProcessNotFound = errors.New("process not found")


func getProcessPid(process string) (pid string, err error){
	ps := exec.Command("ps")

	grep := exec.Command("grep", process)
	grep.Stdin, _ = ps.StdoutPipe()

	ignore := exec.Command("grep", "-v", "grep")
	ignore.Stdin, _ = grep.StdoutPipe()
	out, err := ignore.StdoutPipe()
	if err != nil {
		panic(err)
	}

	err = ps.Start()
	if err != nil {
		panic(err)
	}
	defer ps.Wait()

	err = grep.Start()
	if err != nil {
		panic(err)
	}
	defer grep.Wait()

	err = ignore.Start()
	if err != nil {
		panic(err)
	}
	defer ignore.Wait()

	var lines []string
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	switch len(lines) {
	case 0:
		fmt.Printf("not found %s\n", process)
		return "", ErrProcessNotFound
	case 1:
		line := lines[0]
		pid = strings.Fields(line)[0]
		fmt.Printf("found %s process, pid %s\n", process, pid)
	default:
		line := lines[0]
		pid = strings.Fields(line)[0]
		fmt.Printf("found more than 1 %s process, pid: %s\n", process, pid)
	}

	return pid, nil
}


func killProcess(pid string) {
	cmd := exec.Command("kill", pid)

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("killed %s\n", pid)
	}
}
