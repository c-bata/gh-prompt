package completer

import (
	"fmt"

	"github.com/c-bata/gh-prompt/internal/debug"
	"github.com/c-bata/go-prompt"
)

func argumentsCompleter(repo string, args []string) []prompt.Suggest {
	debug.Log(fmt.Sprintf("repo: %s, args %#v", repo, args))

	if len(args) <= 1 {
		return prompt.FilterHasPrefix(
			[]prompt.Suggest{
				{Text: "help", Description: "Help about any command"},
				{Text: "pr", Description: "Create, view, and checkout pull requests"},
				{Text: "issue", Description: "Create and view issues"},
				// Custom commands.
				{Text: "exit", Description: "Exit this program"},
			},
			args[0],
			true,
		)
	}

	switch args[0] {
	case "issue":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(
				[]prompt.Suggest{
					{Text: "create", Description: "Create a new issue"},
					{Text: "list", Description: "List and filter issues in this repository"},
					{Text: "status", Description: "Show status of relevant issues"},
					{Text: "view", Description: "View an issue in the browser"},
				},
				args[1],
				true,
			)
		}
		// TODO(c-bata): gh issue view {<number> | <url> | <branch>} [flags]
	case "pr":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(
				[]prompt.Suggest{
					{Text: "checkout", Description: "Check out a pull request in Git"},
					{Text: "create", Description: "Create a pull request"},
					{Text: "list", Description: "List and filter pull requests in this repository"},
					{Text: "status", Description: "Show status of relevant pull requests"},
					{Text: "view", Description: "View a pull request in the browser"},
				},
				args[1],
				true,
			)
		}
		// TODO(c-bata): gh pr view [{<number> | <url> | <branch>}] [flags]
	}
	return []prompt.Suggest{}
}
