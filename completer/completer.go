package completer

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/cli/cli/api"
	"github.com/cli/cli/context"
	"github.com/cli/cli/git"
)

var ErrNotFoundRemotes = errors.New("git remotes are not found on your current directory")

type Completer struct {
	client  *api.Client
	remotes context.Remotes
	repo    *api.Repository
}

func fromURL(u *url.URL) (owner, repo string, err error) {
	parts := strings.SplitN(strings.TrimPrefix(u.Path, "/"), "/", 3)
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid path: %s", u.Path)
	}
	return parts[0], strings.TrimSuffix(parts[1], ".git"), nil
}

func NewCompleter(version string) (*Completer, error) {
	client, err := BasicClient(fmt.Sprintf("gh-prompt %s", version))
	if err != nil {
		return nil, err
	}

	gitRemotes, err := git.Remotes()
	if err != nil {
		return nil, err
	}
	if len(gitRemotes) == 0 {
		return nil, ErrNotFoundRemotes
	}
	remotes := make(context.Remotes, 0, len(gitRemotes))
	sshTranslate := git.ParseSSHConfig().Translator()
	for _, r := range gitRemotes {
		var owner, repo string
		if r.FetchURL != nil {
			owner, repo, _ = fromURL(sshTranslate(r.FetchURL))
		}
		if (owner == "" || repo == "") && r.PushURL != nil {
			owner, repo, _ = fromURL(sshTranslate(r.PushURL))
		}
		remotes = append(remotes, &context.Remote{
			Remote: r,
			Owner:  owner,
			Repo:   repo,
		})
	}

	repoContext, err := context.ResolveRemotesToRepos(remotes, client, "")
	if err != nil {
		return nil, err
	}

	baseRepo, err := repoContext.BaseRepo()
	if err != nil {
		return nil, err
	}

	return &Completer{
		client:  client,
		remotes: remotes,
		repo:    baseRepo,
	}, nil
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(commands, d.GetWordBeforeCursor(), true)
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
		return c.optionCompleter(args, w)
	}

	// Return suggestions for option
	if suggests, found := c.completeOptionArguments(d); found {
		return suggests
	}

	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return []prompt.Suggest{}
	}

	repo := checkRepoArg(d)
	return c.argumentsCompleter(repo, commandArgs)

}

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
