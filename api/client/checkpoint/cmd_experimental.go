// +build experimental

package checkpoint

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
)

// NewCheckpointCommand returns a cobra command for `checkpoint` subcommands
func NewCheckpointCommand(rootCmd *cobra.Command, dockerCli *client.DockerCli) {
	cmd := &cobra.Command{
		Use:   "checkpoint",
		Short: "Manage Container Checkpoints",
		Args:  cli.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(dockerCli.Err(), "\n"+cmd.UsageString())
		},
	}
	cmd.AddCommand(
		newCreateCommand(dockerCli),
		newListCommand(dockerCli),
		newRemoveCommand(dockerCli),
	)

	rootCmd.AddCommand(cmd)
}
