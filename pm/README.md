# cc-pm

Utility to install packages and update running processes.

## Usage

The `cc-pm` command exposes the following interface:

```
cc-pm
  --install-path|-i <Installation path>
  --package-url|-p <Package URL>`
```

## How does it work

1. `cc-pm` will start downloading the package using the provided URL
1. It will issue a `SIGTERM` to any running process that gets identified
1. Waits for all signalled processes to terminate
  * If one or more processes don't terminate within a timeout, `cc-pm` will
    issue a `SIGKILL`
1. Extracts the contents of the package in the installation destination
1. Runs the `ctl/start` script found in the installation directory

## Package format

A `cc-pm` package defines the following format:

* `bin` Directory that contains any useful binaries.
* `cfg`
  - `seed-config` Script that creates initial configuration for the package.
  - `update-config` Script that updates current configuration.
* `assets` Contains any useful non-executable files and directories.
* `ctl` Controls the package agents:
  - `start` Script that boots the agents.
  - `stop` Script that stops the agents.
    stops
* `spec.json` The specification of the package. It is thoroughly describe
  bellow.

### Package specification file

The `spec.json` file provides information for the package and defines its
dependencies:

```json
{
  "name": <The package name>,
  "description": <A useful -short- description for the defined package>,
  "agents": [
    <Binary name, the binary has to exist under bin> [...]
  ],
  "dependencies": {
    "apt": [<Aptitute package> [...]],
    "pip": [<Pindip package> [...]]
  }
}
```

The `dependencies` field of the package definition can support multiple
dependency management technologies. It is currently able to install aptitude
package and pip packages.
