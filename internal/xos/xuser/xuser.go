// Package xuser contains functions and utilities to extend Go's standard
// os/user module.
package xuser

import "os"

// IsPrivileged returns true if the current user has root privileges.
func IsPrivileged() bool {
	if os.Getuid() == 0 {
		return true
	}

	_, found := os.LookupEnv("SUDO_USER")

	return found
}
