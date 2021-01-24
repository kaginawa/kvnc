package kvnc


import (
	"os/exec"
	"syscall"
)

// PrepareBackgroundCommand prepares to run exec.Cmd silently.
func PrepareBackgroundCommand(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
