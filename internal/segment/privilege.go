package segment

import (
	"strings"

	"git.sr.ht/~jamesponddotco/gosh/internal/color"
	"git.sr.ht/~jamesponddotco/gosh/internal/xos/xuser"
)

const (
	_user string = "$"
	_root string = "#"
)

// PrivilegeSegment represents a segment that displays the privilege level of
// the current user.
type PrivilegeSegment struct {
	// Suffix is the suffix of the segment.
	Suffix string
}

// Compile-time check to ensure PrivilegeSegment implements the Segmenter
// interface.
var _ Segmenter = (*PrivilegeSegment)(nil)

// Render renders the segment.
func (s *PrivilegeSegment) Render() string {
	var privilege string

	if xuser.IsPrivileged() {
		privilege = _root
	} else {
		privilege = _user
	}

	var textColor string

	if privilege == _root {
		textColor = color.Red
	}

	if privilege == _user {
		textColor = color.Blue
	}

	var builder strings.Builder

	builder.Grow(len(privilege) + len(s.Suffix))

	builder.WriteString(textColor)
	builder.WriteString(privilege)
	builder.WriteString(color.Reset)

	builder.WriteString(s.Suffix)

	return builder.String()
}
