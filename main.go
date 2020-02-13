package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/gh-prompt/completer"
	"github.com/c-bata/gh-prompt/internal/debug"
	"github.com/c-bata/go-prompt"
	gpc "github.com/c-bata/go-prompt/completer"
)

var (
	Version  = "unset"
	Revision = "unset"
)

func executorFunc(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("/bin/sh", "-c", "gh "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func main() {
	fmt.Printf("gh-prompt %s (rev-%s)\n", Version, Revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")

	debug.Log("gh-prompt started")
	defer debug.Teardown()

	p := prompt.New(
		executorFunc,
		completer.Completer,
		prompt.OptionTitle("gh-prompt: interactive GitHub CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(gpc.FilePathCompletionSeparator),
	)
	p.Run()
}
