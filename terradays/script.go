package terradays

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// script execution
func script(path string, args []string, stdin io.Reader) error {
	cmd := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("%s %s", path, strings.Join(args, " ")),
	)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
