package terradays

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Plan ...
type Plan struct {
	Version string  `yaml:"version"`
	Actions Actions `yaml:"actions"`
}

// Load ...
func (p *Plan) Load(dirname string) error {
	logrus.WithField("path", dirname+"/terradays.yaml").Info("loading terradays plan")
	bytes, err := ioutil.ReadFile(dirname + "/terradays.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bytes, p)
}

// Apply runs the plan
func (p *Plan) Apply(dirname, tfargs string) error {
	logrus.Info("applying terradays plan")
	logrus.WithFields(logrus.Fields{
		"dirname": dirname,
		"tfargs":  tfargs,
	}).Debug("applying terradays plan")

	// iterate actions
	for index := range p.Actions {
		actionType, err := p.Actions.ActionType(index)
		if err != nil {
			return err
		}

		switch actionType {
		case ActionTypeModule:
			action, err := p.Actions.ModuleAction(index)
			if err != nil {
				return err
			}
			if err := p.applyModuleAction(action, dirname, tfargs); err != nil {
				return err
			}
		case ActionTypeScript:
			action, err := p.Actions.ScriptAction(index)
			if err != nil {
				return err
			}
			if err := p.applyScriptAction(action, dirname); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unkown action type: %s", actionType)
		}
	}

	return nil
}

// applyModuleAction ...
func (p *Plan) applyModuleAction(action *ModuleAction, dirname, args string) error {
	logrus.WithField("name", action.Name).Info("applying module")
	logrus.WithField("action", action).Debug("applying module")
	return terraform("apply", []string{"-target=module." + action.Name, args, dirname}, os.Stdin)
}

// applyScriptAction ...
func (p *Plan) applyScriptAction(action *ScriptAction, dirname string) error {
	logrus.WithField("path", action.Path).Info("applying script")
	logrus.WithField("action", action).Debug("applying script")
	return script(action.Path, action.Args, action.StdinReader())
}
