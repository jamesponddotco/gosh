package segment

import (
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/gosh/internal/color"
)

// SSHSegment represents a segment that displays the current SSH session.
type SSHSegment struct {
	// The prefix to use in the segment.
	Prefix string

	// The suffix to use in the segment.
	Suffix string
}

// Compile-time check to ensure SSHSegment implements the Segmenter interface.
var _ Segmenter = (*SSHSegment)(nil)

// Render renders the segment.
func (s *SSHSegment) Render() string {
	if os.Getenv("SSH_CLIENT") == "" || os.Getenv("SSH_TTY") == "" {
		return ""
	}

	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}

	var builder strings.Builder

	builder.Grow(len(s.Prefix) + len(host) + len(s.Suffix))

	builder.WriteString(color.White)
	builder.WriteString(s.Prefix)
	builder.WriteString(color.Reset)

	builder.WriteString(color.Green)
	builder.WriteString(host)
	builder.WriteString(color.Reset)

	builder.WriteString(color.White)
	builder.WriteString(s.Suffix)
	builder.WriteString(color.Reset)

	return builder.String()
}
