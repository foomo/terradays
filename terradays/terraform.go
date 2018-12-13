package terradays

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// terraform execution
func terraform(command string, args []string, stdin io.Reader) error {
	cmd := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("terraform %s %s", command, strings.Join(args, " ")),
	)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
