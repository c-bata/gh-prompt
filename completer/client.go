package completer

import (
	"fmt"
	"sync"
	"time"

	"github.com/c-bata/gh-prompt/internal/debug"
	"github.com/c-bata/go-prompt"
	"github.com/cli/cli/api"
	"github.com/cli/cli/context"
)

const (
	thresholdFetchInterval = 60 * time.Second
	issueLimits            = 100
	pullRequestsLimits     = 100
)

var (
	issueCache       *sync.Map
	pullRequestCache *sync.Map
	lastFetchedAt    *sync.Map
	repoCache        *sync.Map
)

func init() {
	issueCache = new(sync.Map)
	pullRequestCache = new(sync.Map)
	lastFetchedAt = new(sync.Map)
	repoCache = new(sync.Map)
}

// client

func currentRepo(ctx *Completer, argRepo string) *api.Repository {
	if argRepo == "" {
		return ctx.repo
	}

	// check cache
	if cached, ok := repoCache.Load(argRepo); ok {
		return cached.(*api.Repository)
	}

	// load repo
	repoCtx, err := context.ResolveRemotesToRepos(ctx.remotes, ctx.client, argRepo)
	if err != nil {
		debug.Log(fmt.Sprintf("err: %s", err))
		return ctx.repo
	}

	repoOverride, err := repoCtx.BaseRepo()
	if err != nil {
		debug.Log(fmt.Sprintf("err: %s", err))
		return ctx.repo
	}
	// cache repo
	repoCache.Store(argRepo, repoOverride)
	return repoOverride
}

// expire

func shouldFetch(key string) bool {
	v, ok := lastFetchedAt.Load(key)
	if !ok {
		return true
	}
	t, ok := v.(time.Time)
	if !ok {
		return true
	}
	return time.Since(t) > thresholdFetchInterval
}

func updateLastFetchedAt(key string) {
	lastFetchedAt.Store(key, time.Now())
}

// issues

func fetchIssuesIfExpired(key string, client *api.Client, repo *api.Repository) {
	if !shouldFetch(key) {
		return
	}
	updateLastFetchedAt(key)

	debug.Log("Call a request to fetch issues.")
	issues, err := api.IssueList(client, repo, "open", nil, "", issueLimits)
	if err != nil {
		debug.Log(err.Error())
	}
	issueCache.Store(key, issues)
	debug.Log("Success to fetch issues")
}

func getIssueCache(key string) (issues []api.Issue, ok bool) {
	v, ok := issueCache.Load(key)
	if !ok {
		return nil, false
	}
	issues, ok = v.([]api.Issue)
	if !ok {
		return nil, ok
	}
	return issues, true
}

func getIssueNumberSuggestions(ctx *Completer, argRepo string) []prompt.Suggest {
	repo := currentRepo(ctx, argRepo)
	cacheKey := fmt.Sprintf("get_issues:%s:%s", repo.RepoOwner(), repo.RepoName())
	go fetchIssuesIfExpired(cacheKey, ctx.client, repo)

	issues, ok := getIssueCache(cacheKey)
	if !ok {
		return []prompt.Suggest{}
	}

	s := make([]prompt.Suggest, len(issues))
	for i := range issues {
		s[i] = prompt.Suggest{
			Text:        fmt.Sprintf("%d", issues[i].Number),
			Description: issues[i].Title,
		}
	}
	return s
}

func getIssueURLSuggestions(ctx *Completer, argRepo string) []prompt.Suggest {
	repo := currentRepo(ctx, argRepo)
	cacheKey := fmt.Sprintf("get_issues:%s:%s", repo.RepoOwner(), repo.RepoName())
	go fetchIssuesIfExpired(cacheKey, ctx.client, repo)

	issues, ok := getIssueCache(cacheKey)
	if !ok {
		return []prompt.Suggest{}
	}

	s := make([]prompt.Suggest, len(issues))
	for i := range issues {
		s[i] = prompt.Suggest{
			Text:        issues[i].URL,
			Description: issues[i].Title,
		}
	}
	return s
}

// pull requests

func fetchPullRequestsIfExpired(key string, client *api.Client, repo *api.Repository) {
	if !shouldFetch(key) {
		return
	}
	params := map[string]interface{}{
		"owner": repo.RepoOwner(),
		"repo":  repo.RepoName(),
		"state": []string{"OPEN"},
	}
	updateLastFetchedAt(key)

	debug.Log("Call a request to fetch pull requests.")
	pulls, err := api.PullRequestList(client, params, pullRequestsLimits)
	if err != nil {
		debug.Log(err.Error())
	}
	pullRequestCache.Store(key, pulls)
	debug.Log("Success to fetch pull requests")
}

func getPullRequestCache(key string) (pulls []api.PullRequest, ok bool) {
	v, ok := pullRequestCache.Load(key)
	if !ok {
		return nil, false
	}
	pulls, ok = v.([]api.PullRequest)
	if !ok {
		return nil, ok
	}
	return pulls, true
}

func getPullRequestsNumberSuggestions(ctx *Completer, argRepo string) []prompt.Suggest {
	repo := currentRepo(ctx, argRepo)
	cacheKey := fmt.Sprintf("get_pulls:%s:%s", repo.RepoOwner(), repo.RepoName())
	go fetchPullRequestsIfExpired(cacheKey, ctx.client, repo)

	pulls, ok := getPullRequestCache(cacheKey)
	if !ok {
		return []prompt.Suggest{}
	}

	s := make([]prompt.Suggest, len(pulls))
	for i := range pulls {
		s[i] = prompt.Suggest{
			Text:        fmt.Sprintf("%d", pulls[i].Number),
			Description: pulls[i].Title,
		}
	}
	return s
}

func getPullRequestsBranchSuggestions(ctx *Completer, argRepo string) []prompt.Suggest {
	repo := currentRepo(ctx, argRepo)
	cacheKey := fmt.Sprintf("get_pulls:%s:%s", repo.RepoOwner(), repo.RepoName())
	go fetchPullRequestsIfExpired(cacheKey, ctx.client, repo)

	pulls, ok := getPullRequestCache(cacheKey)
	if !ok {
		return []prompt.Suggest{}
	}

	s := make([]prompt.Suggest, len(pulls))
	for i := range pulls {
		s[i] = prompt.Suggest{
			Text:        pulls[i].BaseRefName,
			Description: pulls[i].Title,
		}
	}
	return s
}

func getPullRequestsURLSuggestions(ctx *Completer, argRepo string) []prompt.Suggest {
	repo := currentRepo(ctx, argRepo)
	cacheKey := fmt.Sprintf("get_pulls:%s:%s", repo.RepoOwner(), repo.RepoName())
	go fetchPullRequestsIfExpired(cacheKey, ctx.client, repo)

	pulls, ok := getPullRequestCache(cacheKey)
	if !ok {
		return []prompt.Suggest{}
	}

	s := make([]prompt.Suggest, len(pulls))
	for i := range pulls {
		s[i] = prompt.Suggest{
			Text:        pulls[i].URL,
			Description: pulls[i].Title,
		}
	}
	return s
}
