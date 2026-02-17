package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 3 {
		fmt.Println("missing arguments")
		return 1
	}
	timesStr := os.Args[1]
	times, err := strconv.Atoi(timesStr)
	if err != nil {
		fmt.Println("could not parse times argument: ", timesStr)
		return 1
	}
	cmdName := os.Args[2]
	cmdArgs := os.Args[3:]

	successes := 0
	defer func() { fmt.Println("repeated successfully", successes, "/", times, "times") }()
	for range times {
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		e := cmd.Run()
		if e != nil {
			if exerr := new(exec.ExitError); errors.As(e, &exerr) {
				return exerr.ExitCode()
			}
		}
		successes++
	}

	return 0
}
