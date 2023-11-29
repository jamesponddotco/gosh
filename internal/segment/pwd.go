package segment

import (
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/gosh/internal/color"
	"git.sr.ht/~jamesponddotco/gosh/internal/xos"
)

// PWDSegment represents a segment that displays the current working directory.
type PWDSegment struct {
	// Prefix is the text shown before the current working directory.
	Prefix string

	// Suffix is the text shown after the current working directory.
	Suffix string

	// Parts is the number of directory parts to show.
	Parts int
}

// Compile-time check to ensure that PWDSegment implements the Segmenter
// interface.
var _ Segmenter = (*PWDSegment)(nil)

// Render renders the segment.
func (s *PWDSegment) Render() string {
	var (
		cwd string
		err error
	)

	if s.Parts <= 0 {
		cwd, err = os.Getwd()
		if err != nil {
			return ""
		}
	} else {
		cwd, err = xos.Gettwd(s.Parts)
		if err != nil {
			return ""
		}
	}

	var builder strings.Builder

	builder.Grow(len(s.Prefix) + len(cwd) + len(s.Suffix))

	builder.WriteString(color.White)
	builder.WriteString(s.Prefix)
	builder.WriteString(color.Reset)

	builder.WriteString(color.Blue)
	builder.WriteString(cwd)
	builder.WriteString(color.Reset)

	builder.WriteString(color.White)
	builder.WriteString(s.Suffix)
	builder.WriteString(color.Reset)

	return builder.String()
}
