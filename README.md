# Terradays

Terradays is a small utility around terraform that tries to help to workaround some issues that are currently not yet fixed. For now, it will read in a `terradays.yaml` file containing a list of actions that it will then apply in sequence.

## Installation

Download and install precombiled [releases](https://github.com/foomo/terradays/releases).

On Mac OS X you can also install the latest release via Homebrew:

```bash
brew install foomo/terradays/terradays
```

## Usage

```
Available Commands:
  apply       builds or changes infrastructure
  help        Help about any command
  version     prints out version number

Flags:
      --debug   Debug output
  -h, --help    help for terradays
```

## Configuration

```yaml
# Configuration version
version: "1"

# List of actions to run in sequence
actions:
  # Module action
  - type: module
    name: cluster       # Name of the module to run
  # Script action
  - type: script
    path: script.sh     # Path to the script to execute
    args: []            # Script arguements
    stdin: ""           # Script stdin
```