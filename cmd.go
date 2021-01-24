// +build !windows

package kvnc

import "os/exec"

// PrepareBackgroundCommand prepares to run exec.Cmd silently.
func PrepareBackgroundCommand(_ *exec.Cmd) {
	// no-op
}
