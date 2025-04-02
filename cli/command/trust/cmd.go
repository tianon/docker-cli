package trust

import (
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

// NewTrustCommand returns a cobra command for `trust` subcommands
func NewTrustCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trust",
		Short: "Manage trust on Docker images",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCli.Err()),

		Deprecated: command.ContentTrustDeprecationWarning,
	}
	cmd.AddCommand(
		newRevokeCommand(dockerCli),
		newSignCommand(dockerCli),
		newTrustKeyCommand(dockerCli),
		newTrustSignerCommand(dockerCli),
		newInspectCommand(dockerCli),
	)
	// apply the deprecation notice to all subcommands too (so they *also* print out the notice on usage)
	// unfortunately, this also marks them all as "hidden" also, so "docker trust --help" is now empty, which isn't ideal, but maybe it's fine for this intermediate state?
	children := cmd.Commands()
	for len(children) > 0 {
		child := children[0]
		children = children[1:]
		child.Deprecated = cmd.Deprecated
		children = append(children, child.Commands()...)
	}
	return cmd
}
