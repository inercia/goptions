`goptions` implements a flexible parser for command line options.

Key targets were the support for both long and short flag versions, mutually
exclusive flags, and verbs. Flags and their corresponding variables are defined
by the tags in a (possibly anonymous) struct.

![](https://circleci.com/gh/voxelbrain/goptions.png?circle-token=27cd98362d475cfa8c586565b659b2204733f25c)


# Example

```Go
package main

import (
  "github.com/voxelbrain/goptions"
  "os"
  "time"
)

func main() {
  options := struct {
    Server   string        `goptions:"-s, --server,
                                      maps='Global/Server',
                                      obligatory, description='Server to connect to'"`
    Password string        `goptions:"-p, --password,
                                      description='Don\\'t prompt for password'"`
    Timeout  time.Duration `goptions:"-t, --timeout,
                                      maps='Global/Timeout',
                                      description='Connection timeout in seconds'"`
    Help     goptions.Help `goptions:"-h, --help,
                                      description='Show this help'"`

    goptions.Verbs
    Execute struct {
      Command string   `goptions:"--command, mutexgroup='input',
                                  description='Command to exectute', obligatory"`
      Script  *os.File `goptions:"--script, mutexgroup='input',
                                  description='Script to exectute', rdonly"`
    } `goptions:"execute"`

    Delete struct {
      Path  string `goptions:"-n, --name, obligatory,
                              description='Name of the entity to be deleted'"`
      Force bool   `goptions:"-f, --force, description='Force removal'"`
    } `goptions:"delete"`

  }{ // Default values goes here
    Timeout: 10 * time.Second,
  }

  goptions.ParseAndFail(&options)
  config := struct{
      type Global struct {
          Server   string
          Timeout  int
      }
  }
  err := goptions.LoadConf(config)
}
```

```
$ go run examples/readme_example.go --help
Usage: a.out [global options] <verb> [verb options]

Global options:
        -s, --server   Server to connect to (*)
        -p, --password Don't prompt for password
        -t, --timeout  Connection timeout in seconds (default: 10s)
        -h, --help     Show this help

Verbs:
    delete:
        -n, --name     Name of the entity to be deleted (*)
        -f, --force    Force removal
    execute:
            --command  Command to exectute (*)
            --script   Script to exectute
```

## Fork changes

This fork implements some basic config file functionality, and it lets you map some command line options with config file parameter: you can use the `maps=` attribute for assigning a command line option to an (arbitrarily deep) structure element, and it will be automatically filled up when you parse the arguments.
