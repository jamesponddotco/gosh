// Package xos contains functions and utilities to extend Go's standard os
// module.
package xos

import (
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

// ErrLessThanZero is returned when the length given to Gettwd is less than
// zero.
const ErrLessThanZero xerrors.Error = "length must be greater than or equal to zero"

const (
	_ellipsis  string = "..."
	_separator string = string(os.PathSeparator)
)

// Gettwd returns a truncated version of the current working directory. The
// truncation retains the specified number of final path elements, providing a
// shorter representation of the path similar to the behavior of PROMPT_DIRTRIM
// in Bash.
//
// If the length parameter is less than zero, Gettwd returns an ErrLessThanZero
// error. If the length is greater than or equal to the number of elements in
// the current directory path, the full path is returned without truncation.
func Gettwd(length int) (string, error) {
	if length < 0 {
		return "", ErrLessThanZero
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if strings.HasPrefix(cwd, home) {
		cwd = "~" + cwd[len(home):]
	}

	parts := strings.Split(cwd, _separator)

	if len(parts)-1 <= length {
		return cwd, nil
	}

	var builder strings.Builder

	if strings.HasPrefix(cwd, "~") {
		builder.WriteString("~")

		if len(parts) > 1 {
			builder.WriteString(_separator)
			builder.WriteString(_ellipsis)
			builder.WriteString(_separator)
		}
	} else {
		builder.WriteString(_ellipsis)
		builder.WriteString(_separator)
	}

	builder.WriteString(xstrings.JoinWithSeparator(_separator, parts[len(parts)-length:]...))

	return builder.String(), nil
}
