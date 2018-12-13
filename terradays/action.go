package terradays

import (
	"io"
	"strings"
)

// Action is the generic data interface
type Action interface{}

// ActionType definition
type ActionType string

const (
	// ActionTypeModule type
	ActionTypeModule ActionType = "module"
	// ActionTypeScript type
	ActionTypeScript ActionType = "script"
)

// BaseAction ...
type BaseAction struct {
	Type ActionType `yaml:"type"`
}

// ModuleAction ...
type ModuleAction struct {
	BaseAction `mapstructure:",squash"`
	Name       string `yaml:"name"`
}

// ScriptAction ...
type ScriptAction struct {
	BaseAction `mapstructure:",squash"`
	Args       []string `yaml:"args"`
	Path       string   `yaml:"path"`
	Stdin      string   `yaml:"stdin"`
}

// StdinReader returns a io.Reader for Stdin
func (a *ScriptAction) StdinReader() io.Reader {
	return strings.NewReader(a.Stdin)
}
