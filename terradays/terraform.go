package terradays

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Terraform hold the path to the command line tool
var Terraform string

// terraform execution
func terraform(command string, args []string, stdin io.Reader) error {
	cmd := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("%s %s %s", Terraform, command, strings.Join(args, " ")),
	)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
