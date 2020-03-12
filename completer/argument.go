package completer

import (
	"fmt"

	"github.com/c-bata/gh-prompt/internal/debug"
	"github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "help", Description: "Help about any command"},
	{Text: "pr", Description: "Create, view, and checkout pull requests"},
	{Text: "repo", Description: "Create, clone, fork, and view repositories"},
	{Text: "issue", Description: "Create and view issues"},
	// Custom commands.
	{Text: "exit", Description: "Exit this program"},
}

func (c *Completer) argumentsCompleter(repo string, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(
			commands,
			args[0],
			true,
		)
	}

	switch args[0] {
	case "issue":
		debug.Log(fmt.Sprintf("here! %#v", args))
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
		if args[1] == "view" && len(args) == 3 {
			suggests := getIssueNumberSuggestions(c, repo)
			suggests = append(suggests, getIssueURLSuggestions(c, repo)...)
			return prompt.FilterHasPrefix(
				suggests,
				args[2],
				true,
			)
		}
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
		if args[1] == "view" && len(args) == 3 {
			suggests := getPullRequestsNumberSuggestions(c, repo)
			suggests = append(suggests, getPullRequestsBranchSuggestions(c, repo)...)
			// This makes 'Text' section of completion window too long.
			// suggests = append(suggests, getPullRequestsURLSuggestions(c, repo)...)
			return prompt.FilterHasPrefix(
				suggests,
				args[2],
				true,
			)
		}
		if args[1] == "checkout" && len(args) == 3 {
			suggests := getPullRequestsNumberSuggestions(c, repo)
			suggests = append(suggests, getPullRequestsBranchSuggestions(c, repo)...)
			// This makes 'Text' section of completion window too long.
			// suggests = append(suggests, getPullRequestsURLSuggestions(c, repo)...)
			return prompt.FilterHasPrefix(
				suggests,
				args[2],
				true,
			)
		}
	case "repo":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(
				[]prompt.Suggest{
					{Text: "clone", Description: "Clone a repository locally"},
					{Text: "create", Description: "Create a new repository"},
					{Text: "fork", Description: "Create a fork of a repository."},
					{Text: "view", Description: "View a repository in the browser."},
				},
				args[1],
				true,
			)
		}
	}
	return []prompt.Suggest{}
}
