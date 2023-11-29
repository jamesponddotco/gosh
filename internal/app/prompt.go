package app

import (
	"fmt"
	"strings"

	"git.sr.ht/~jamesponddotco/gosh/internal/segment"
	"github.com/urfave/cli/v2"
)

const (
	_commonPrefix string = "["
	_commonSuffix string = "]"
	_gitPrefix    string = "("
	_gitSuffix    string = ")"
)

// PromptAction is the main action for the application.
func PromptAction(ctx *cli.Context) error {
	var (
		builder     strings.Builder
		pathLength  = ctx.Int("path-length")
		sshSegment  = ctx.Bool("ssh-segment")
		gitSegment  = ctx.Bool("git-segment")
		privSegment = ctx.Bool("privilege-segment")
	)

	if sshSegment {
		ssh := segment.SSHSegment{
			Prefix: _commonPrefix,
			Suffix: _commonSuffix,
		}

		builder.WriteString(ssh.Render())
	}

	cwd := segment.PWDSegment{
		Prefix: _commonPrefix,
		Suffix: _commonSuffix,
		Parts:  pathLength,
	}

	builder.WriteString(cwd.Render())

	if gitSegment {
		git := segment.GitSegment{
			Prefix: _gitPrefix,
			Suffix: _gitSuffix,
		}

		builder.WriteString(git.Render())
	}

	if privSegment {
		priv := segment.PrivilegeSegment{
			Suffix: " ",
		}

		builder.WriteString(priv.Render())
	}

	fmt.Fprint(ctx.App.Writer, builder.String())

	return nil
}
