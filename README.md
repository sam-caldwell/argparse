go-argparse
===========

## Description

A python-approximate golang commandline argument parser. The goal is to create a Command-line argument
parser similar to python argparse. Obviously, golang will have some differences.

## Usage

* See [example](example/main.go).

## Features

| Feature                                                        | Description                                                       |
|----------------------------------------------------------------|-------------------------------------------------------------------|
| `args.Init({parameters...string})`                             | must be run or this package will throw errors                     |
| `args.EnableDebugging()`                                       | enable the debugging feature                                      |
| `args.AddArgument(name,type,required,defaultValue,helpString)` | Add an argument (positional or optional)                          |
| `args.ParseArgs()`                                             | Parse os.Args[], validate values                                  | 
| `args.Dump()`                                                  | Return a list of strings representing the argument and its values |
| `args.Get(<name>)`                                             | Get the value for the given name.                                 |

### `Arguments.Init({parameters...string})`

> This method will initialize `Arguments` memory objects, and then it appends a
> list of descriptor strings. These descriptor strings could include copyright,
> author, description or whatever else is needed.

- Input - zero or many strings representing descriptor strings which will be used in help text.
- Output - none

### `Arguments.EnableDebugging()`

> Set the debug flag to allow debug logging facilities to stdout

- Input - None
- Output - None

### `Arguments.AddArgument(name, type, required, defaultValue, helpString)`

> Add an optional or positional argument descriptor to the `Arguments` struct

- Input:
    - `name` : string
        - `[string]`: positional argument
        - `-[letter]`: short optional argument
        - `--[string]`: long optional argument
    - `argType` : numeric argument type (`types.String`, `types.Integer`, ...)
    - `required` : bool
    - `defaultValue` : `{arbitrary value}`
    - `helpString` : string
- Output:
    - error

### `Arguments.ParseArgs()`

> Parse the arguments in `os.Args[]` into `Arguments`

- Input: None
- Output: error

### `Arguments.Dump()`

> Dump the current state of the `Arguments` object to `[]string` list

- Input: None
- Output: []string

### `Arguments.Get(<name string>)`

> Given a parameter name (string) return the arbitrary value of the parsed argument.

- Input: string
- Output: `{arbitrary value}`