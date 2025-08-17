package git

import (
	"context"
	"fmt"
	"io"

	execCmd "github.com/yupsh/exec"
	execOpt "github.com/yupsh/exec/opt"
	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/git/opt"
)

// Flags represents the configuration options for git commands
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Git creates a new git command with the given parameters
func Git(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, input io.Reader, output, stderr io.Writer) error {
	if len(c.Positional) == 0 {
		return fmt.Errorf("git: no command specified")
	}

	// Build git command arguments
	args := make([]any, 0, len(c.Positional)+10)
	args = append(args, "git")

	// Add positional arguments
	for _, arg := range c.Positional {
		args = append(args, arg)
	}

	// Convert git flags to exec options
	execOpts := c.buildExecOptions()
	args = append(args, execOpts...)

	// Create and execute the underlying exec command
	execCommand := execCmd.Exec(args...)
	return execCommand.Execute(ctx, input, output, stderr)
}

func (c command) buildExecOptions() []any {
	var opts []any

	// Add working directory if repository is specified and looks like a path
	if c.Flags.Repository != "" {
		opts = append(opts, execOpt.WorkingDir(string(c.Flags.Repository)))
	}

	// Add quiet/verbose flags to exec
	if bool(c.Flags.Quiet) {
		opts = append(opts, execOpt.Quiet)
	}

	// Add environment variables for author if specified
	if c.Flags.Author != "" {
		opts = append(opts, execOpt.EnvVar(fmt.Sprintf("GIT_AUTHOR_NAME=%s", c.Flags.Author)))
	}
	if c.Flags.Email != "" {
		opts = append(opts, execOpt.EnvVar(fmt.Sprintf("GIT_AUTHOR_EMAIL=%s", c.Flags.Email)))
	}

	// Always inherit environment
	opts = append(opts, execOpt.InheritEnv)

	return opts
}

func (c command) String() string {
	return fmt.Sprintf("git %v", c.Positional)
}

// Wrapper commands for common git operations

// Status creates a git status command
func Status(parameters ...any) yup.Command {
	args := append([]any{"status"}, parameters...)
	return Git(args...)
}

// Commit creates a git commit command
func Commit(message string, parameters ...any) yup.Command {
	args := make([]any, 0, 4+len(parameters))
	args = append(args, "commit", "-m", message)
	args = append(args, parameters...)
	return Git(args...)
}

// CommitAll creates a git commit -a command
func CommitAll(message string, parameters ...any) yup.Command {
	args := make([]any, 0, 5+len(parameters))
	args = append(args, "commit", "-a", "-m", message)
	args = append(args, parameters...)
	return Git(args...)
}

// Push creates a git push command
func Push(parameters ...any) yup.Command {
	args := append([]any{"push"}, parameters...)
	return Git(args...)
}

// PushOrigin creates a git push origin command
func PushOrigin(branch string, parameters ...any) yup.Command {
	args := make([]any, 0, 3+len(parameters))
	args = append(args, "push", "origin", branch)
	args = append(args, parameters...)
	return Git(args...)
}

// Pull creates a git pull command
func Pull(parameters ...any) yup.Command {
	args := append([]any{"pull"}, parameters...)
	return Git(args...)
}

// PullOrigin creates a git pull origin command
func PullOrigin(branch string, parameters ...any) yup.Command {
	args := make([]any, 0, 3+len(parameters))
	args = append(args, "pull", "origin", branch)
	args = append(args, parameters...)
	return Git(args...)
}

// Fetch creates a git fetch command
func Fetch(parameters ...any) yup.Command {
	args := append([]any{"fetch"}, parameters...)
	return Git(args...)
}

// FetchOrigin creates a git fetch origin command
func FetchOrigin(parameters ...any) yup.Command {
	args := append([]any{"fetch", "origin"}, parameters...)
	return Git(args...)
}

// Add creates a git add command
func Add(parameters ...any) yup.Command {
	args := append([]any{"add"}, parameters...)
	return Git(args...)
}

// AddAll creates a git add -A command
func AddAll(parameters ...any) yup.Command {
	args := append([]any{"add", "-A"}, parameters...)
	return Git(args...)
}

// Log creates a git log command
func Log(parameters ...any) yup.Command {
	args := append([]any{"log"}, parameters...)
	return Git(args...)
}

// LogOneline creates a git log --oneline command
func LogOneline(parameters ...any) yup.Command {
	args := append([]any{"log", "--oneline"}, parameters...)
	return Git(args...)
}

// Diff creates a git diff command
func Diff(parameters ...any) yup.Command {
	args := append([]any{"diff"}, parameters...)
	return Git(args...)
}

// Checkout creates a git checkout command
func Checkout(branch string, parameters ...any) yup.Command {
	args := make([]any, 0, 2+len(parameters))
	args = append(args, "checkout", branch)
	args = append(args, parameters...)
	return Git(args...)
}

// CheckoutNew creates a git checkout -b command
func CheckoutNew(branch string, parameters ...any) yup.Command {
	args := make([]any, 0, 3+len(parameters))
	args = append(args, "checkout", "-b", branch)
	args = append(args, parameters...)
	return Git(args...)
}
