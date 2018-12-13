package terradays

import "github.com/mitchellh/mapstructure"

// Actions list
type Actions []Action

// ActionType returns the type for the given index
func (a Actions) ActionType(index int) (actionType ActionType, err error) {
	var baseAction *BaseAction
	if baseAction, err = a.BaseAction(index); err == nil {
		actionType = baseAction.Type
	}
	return
}

// BaseAction returns the typed action for the given index
func (a Actions) BaseAction(index int) (action *BaseAction, err error) {
	err = mapstructure.Decode(a[index], &action)
	return
}

// ModuleAction returns the typed action for the given index
func (a Actions) ModuleAction(index int) (action *ModuleAction, err error) {
	err = mapstructure.Decode(a[index], &action)
	return
}

// ScriptAction returns the typed action for the given index
func (a Actions) ScriptAction(index int) (action *ScriptAction, err error) {
	err = mapstructure.Decode(a[index], &action)
	return
}
