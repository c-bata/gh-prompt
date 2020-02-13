package completer

import (
	"fmt"
	"os"

	"github.com/cli/cli/api"
	"github.com/cli/cli/context"
)

func BasicClient(ua string) (*api.Client, error) {
	opts := []api.ClientOption{
		api.AddHeader("User-Agent", ua),
	}
	if c, err := context.ParseDefaultConfig(); err == nil {
		opts = append(opts, api.AddHeader("Authorization", fmt.Sprintf("token %s", c.Token)))
	}
	if verbose := os.Getenv("DEBUG"); verbose != "" {
		opts = append(opts, api.VerboseLog(os.Stderr, false))
	}
	return api.NewClient(opts...), nil
}
