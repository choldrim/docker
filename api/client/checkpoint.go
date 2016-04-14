package client

import (
	"fmt"

	"golang.org/x/net/context"

	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/engine-api/types"
)


// CmdCheckpoint is the parent subcommand for all checkpoint commands
//
// Usage: docker checkpoint <COMMAND> [OPTIONS]
func (cli *DockerCli) CmdCheckpoint(args ...string) error {
	cmd := Cli.Subcmd("checkpoint", []string{"COMMAND [OPTIONS]"}, checkpointUsage(), false)
	cmd.Require(flag.Min, 1)
	err := cmd.ParseFlags(args, true)
	cmd.Usage()
	return err
}

// CmdCheckpointCreate checkpoints the process running in a container
//
// Usage: docker checkpoint create CONTAINER CHECKPOINT
func (cli *DockerCli) CmdCheckpointCreate(args ...string) error {
	cmd := Cli.Subcmd("checkpoint create", []string{"CONTAINER CHECKPOINT"}, "Create a checkpoint from a running container", true)
	cmd.Require(flag.Exact, 2)

	flLeaveRunning := cmd.Bool([]string{"-leave-running"}, false, "leave the container running after checkpoint")

	cmd.ParseFlags(args, true)

	options := types.CheckpointCreateOptions{
		CheckpointID:   cmd.Arg(1),
		Exit:           !*flLeaveRunning,
	}

	err := cli.client.CheckpointCreate(context.Background(), cmd.Arg(0), options)
	if err != nil {
		return err
	}

	return nil
}


// CmdCheckpointDelete deletes a container's checkpoint
//
// Usage: docker checkpoint delete <CONTAINER> <CHECKPOINT>
func (cli *DockerCli) CmdCheckpointDelete(args ...string) error {
	cmd := Cli.Subcmd("checkpoint delete", []string{"CONTAINER CHECKPOINT"}, "Delete a container's checkpoint", false)
	cmd.Require(flag.Exact, 2)
	if err := cmd.ParseFlags(args, true); err != nil {
		return err
	}

	return cli.client.CheckpointDelete(context.Background(), cmd.Arg(0), cmd.Arg(1))
}


func checkpointUsage() string {
	checkpointCommands := [][]string{
		{"create", "Create a checkpoint from a running container"},
		{"delete", "Delete an existing checkpoint"},
		{"list", "List all checkpoints for a container"},
	}

	help := "Commands:\n"

	for _, cmd := range checkpointCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd[0], cmd[1])
	}

	help += fmt.Sprintf("\nRun 'docker checkpoint COMMAND --help' for more information on a command.")
	return help
}
