// +build !experimental

package checkpoint

import (
	"github.com/docker/docker/api/client"
	"github.com/spf13/cobra"
)

// NewCheckpointCommand returns a cobra command for `checkpoint` subcommands
func NewCheckpointCommand(rootCmd *cobra.Command, dockerCli *client.DockerCli) {
}
