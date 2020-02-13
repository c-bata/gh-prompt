package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

var (
	Version  = "unset"
	Revision = "unset"
)

func completerFunc(d prompt.Document) []prompt.Suggest {
	return nil
}

func executorFunc(in string) {
	fmt.Println(in)
}

func main() {
	fmt.Printf("gh-prompt %s (rev-%s)\n", Version, Revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		executorFunc,
		completerFunc,
		prompt.OptionTitle("gh-prompt: interactive GitHub CLI"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
