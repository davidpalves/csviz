package pager

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Pager(data string) {
	cmd := exec.Command(os.ExpandEnv("$PAGER"), "-S")

	// Feed it with the string you want to display.
	cmd.Stdin = strings.NewReader(data)

	// This is crucial - otherwise it will write to a null device.
	cmd.Stdout = os.Stdout

	// Fork off a process and wait for it to terminate.
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
