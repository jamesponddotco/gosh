package segment

import (
	"os/exec"
	"strings"

	"git.sr.ht/~jamesponddotco/gosh/internal/color"
	"git.sr.ht/~jamesponddotco/xstd-go/xunsafe"
)

// GitSegment represents a segment that displays the current git branch.
type GitSegment struct {
	// Prefix is the text before the git branch.
	Prefix string

	// Suffix is the text after the git branch.
	Suffix string
}

// Compile-time check to ensure that GitSegment implements the Segmenter
// interface.
var _ Segmenter = (*GitSegment)(nil)

// Render renders the segment.
func (s *GitSegment) Render() string {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return ""
	}

	branchStr := strings.TrimSpace(xunsafe.BytesToString(branch))

	var builder strings.Builder

	builder.Grow(len(s.Prefix) + len(branch) + len(s.Suffix))

	builder.WriteString(color.White)
	builder.WriteString(s.Prefix)
	builder.WriteString(color.Reset)

	builder.WriteString(color.Yellow)
	builder.WriteString(branchStr)
	builder.WriteString(color.Reset)

	builder.WriteString(color.White)
	builder.WriteString(s.Suffix)
	builder.WriteString(color.Reset)

	return builder.String()
}
