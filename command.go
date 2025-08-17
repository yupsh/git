package command

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	yup "github.com/gloo-foo/framework"
)

type command yup.Inputs[string, flags]

func Git(parameters ...any) yup.Command {
	return command(yup.Initialize[string, flags](parameters...))
}

func (p command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Get git subcommand and arguments from positional parameters
		if len(p.Positional) == 0 {
			_, _ = fmt.Fprintf(stderr, "git: no subcommand specified\n")
			return fmt.Errorf("git requires a subcommand")
		}

		// Build git command with all positional arguments
		args := p.Positional

		// Execute git
		cmd := exec.CommandContext(ctx, "git", args...)
		cmd.Stdin = stdin
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		return cmd.Run()
	}
}
