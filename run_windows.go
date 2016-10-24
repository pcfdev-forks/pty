package pty

import (
	"errors"
	"os"
	"os/exec"
)

func Start(c *exec.Cmd) (pty *os.File, err error) {
	return nil, errors.New("unsupported on windows")
}
