package main

import "github.com/foomo/terradays/cmd"

// version is being set during build
var version string

func main() {
	cmd.Execute(version)
}
