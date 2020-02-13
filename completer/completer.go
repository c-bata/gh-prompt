package completer

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func parseArgs(t string) []string {
	splits := strings.Split(t, " ")
	args := make([]string, 0, len(splits))

	for i := range splits {
		if i != len(splits)-1 && splits[i] == "" {
			continue
		}
		args = append(args, splits[i])
	}
	return args
}

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := parseArgs(d.TextBeforeCursor())
	w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, w)
	}

	// Return suggestions for option
	if suggests, found := completeOptionArguments(d); found {
		return suggests
	}

	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return []prompt.Suggest{}
	}

	repo := checkRepoArg(d)
	return argumentsCompleter(repo, commandArgs)
}

func checkRepoArg(d prompt.Document) string {
	args := strings.Split(d.Text, " ")
	var found bool
	for i := 0; i < len(args); i++ {
		if found {
			return args[i]
		}
		if args[i] == "--repo" || args[i] == "-R" {
			found = true
			continue
		}
	}
	return ""
}
