package completer

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var globalOptions = []prompt.Suggest{
	{Text: "--repo", Description: "Select another repository using the OWNER/REPO format"},
	{Text: "-R", Description: "Select another repository using the OWNER/REPO format"},
	{Text: "--help", Description: "Show help for command"},
}

func (c *Completer) optionCompleter(args []string, word string) []prompt.Suggest {
	l := len(args)
	long := strings.HasPrefix(word, "--")
	if l <= 2 {
		return prompt.FilterHasPrefix(globalOptions, word, false)
	}

	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "issue":
		switch commandArgs[1] {
		case "create":
			suggests = []prompt.Suggest{
				{Text: "-b", Description: "Supply a body. Will prompt for one otherwise."},
				{Text: "--body", Description: "Supply a body. Will prompt for one otherwise."},
				{Text: "-t", Description: "Supply a title. Will prompt for one otherwise."},
				{Text: "--title", Description: "Supply a title. Will prompt for one otherwise."},
				{Text: "-w", Description: "Open the browser to create an issue"},
				{Text: "--web", Description: "Open the browser to create an issue"},
			}
		case "list":
			suggests = []prompt.Suggest{
				{Text: "-a", Description: "Filter by assignee"},
				{Text: "--assignee", Description: "Filter by assignee"},
				{Text: "-l", Description: "Filter by label"},
				{Text: "--label", Description: "Filter by label"},
				{Text: "-L", Description: "Maximum number of issues to fetch (default 30)"},
				{Text: "--limit", Description: "Maximum number of issues to fetch (default 30)"},
				{Text: "-s", Description: "Filter by state: {open|closed|all}"},
				{Text: "--state", Description: "Filter by state: {open|closed|all}"},
			}
		case "status":
			suggests = []prompt.Suggest{}
		case "view":
			suggests = []prompt.Suggest{
				{Text: "-p", Description: "Display preview of issue content"},
				{Text: "--preview", Description: "Display preview of issue content"},
			}
		}
	case "pr":
		switch commandArgs[1] {
		case "checkout":
			suggests = []prompt.Suggest{}
		case "create":
			suggests = []prompt.Suggest{
				{Text: "-B", Description: "The branch into which you want your code merged"},
				{Text: "--base", Description: "The branch into which you want your code merged"},
				{Text: "-b", Description: "Supply a body. Will prompt for one otherwise."},
				{Text: "--body", Description: "Supply a body. Will prompt for one otherwise."},
				{Text: "-d", Description: "Mark pull request as a draft"},
				{Text: "--draft", Description: "Mark pull request as a draft"},
				{Text: "-t", Description: "Supply a title. Will prompt for one otherwise."},
				{Text: "--title", Description: "Supply a title. Will prompt for one otherwise."},
				{Text: "-w", Description: "Open the web browser to create a pull request"},
				{Text: "--web", Description: "Open the web browser to create a pull request"},
			}
		case "list":
			suggests = []prompt.Suggest{
				{Text: "-a", Description: "Filter by assignee"},
				{Text: "--assignee", Description: "Filter by assignee"},
				{Text: "-B", Description: "Filter by base branch"},
				{Text: "--base", Description: "Filter by base branch"},
				{Text: "-l", Description: "Filter by label"},
				{Text: "--label", Description: "Filter by label"},
				{Text: "-L", Description: "Maximum number of items to fetch (default 30)"},
				{Text: "--limit", Description: "Maximum number of items to fetch (default 30)"},
				{Text: "-s", Description: "Filter by state: {open|closed|merged|all} (default 'open')"},
				{Text: "--state", Description: "Filter by state: {open|closed|merged|all} (default 'open')"},
			}
		case "status":
			suggests = []prompt.Suggest{}
		case "view":
			suggests = []prompt.Suggest{
				{Text: "-p", Description: "Display preview of pull request content"},
				{Text: "--preview", Description: "Display preview of pull request content"},
			}
		}
	default:
		suggests = []prompt.Suggest{}
	}

	suggests = append(suggests, globalOptions...)
	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

func getPreviousOption(d prompt.Document) (cmds []string, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}

	cmds, _ = excludeOptions(args)
	if strings.HasPrefix(option, "-") {
		return cmds, option, true
	}
	return nil, "", false
}

func (c *Completer) completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmds, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}

	// repository
	if option == "-R" || option == "--repo" {
		return prompt.FilterHasPrefix(
			[]prompt.Suggest{},
			d.GetWordBeforeCursor(),
			true,
		), true
	}

	switch cmds[0] {
	case "issue":
		if len(cmds) < 2 {
			return []prompt.Suggest{}, false
		}

		switch cmds[1] {
		case "create":
			switch option {
			case "-b", "--body":
				return []prompt.Suggest{}, true
			case "-t", "--title":
				return []prompt.Suggest{}, true
			}
		case "list":
			switch option {
			case "-a", "--assignee":
				return []prompt.Suggest{}, true
			case "-l", "--label":
				// TODO(c-bata): complete label
				return []prompt.Suggest{}, true
			case "-L", "--limit":
				return []prompt.Suggest{}, true
			case "-s", "--state":
				return prompt.FilterHasPrefix(
					[]prompt.Suggest{
						{Text: "open"},
						{Text: "closed"},
						{Text: "all"},
					},
					d.GetWordBeforeCursor(),
					true,
				), true
			}
		}
	case "pr":
		if len(cmds) < 2 {
			return []prompt.Suggest{}, false
		}

		switch cmds[1] {
		case "create":
			switch option {
			case "-B", "--base":
				return []prompt.Suggest{}, true
			case "-b", "--body":
				return []prompt.Suggest{}, true
			case "-t", "--title":
				return []prompt.Suggest{}, true
			}
		case "list":
			switch option {
			case "-a", "--assignee":
				return []prompt.Suggest{}, true
			case "-B", "--base":
				return []prompt.Suggest{}, true
			case "-l", "--label":
				// TODO(c-bata): complete label
				return []prompt.Suggest{}, true
			case "-L", "--limit":
				return []prompt.Suggest{}, true
			case "-s", "--state":
				return prompt.FilterHasPrefix(
					[]prompt.Suggest{
						{Text: "open"},
						{Text: "closed"},
						{Text: "merged"},
						{Text: "all"},
					},
					d.GetWordBeforeCursor(),
					true,
				), true
			}
		}
	}

	return []prompt.Suggest{}, false
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	cmd := args[0]
	filtered := make([]string, 0, l)

	var skipNextArg bool
	for i := 0; i < len(args); i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		if cmd == "logs" && args[i] == "-f" {
			continue
		}

		for _, s := range []string{
			"-b", "--body",
			"-t", "--title",
			"-a", "--assignee",
			"-l", "--label",
			"-L", "--limit",
			"-s", "--state",
			"-B", "--base",
		} {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}
		filtered = append(filtered, args[i])
	}
	return filtered, skipNextArg
}
