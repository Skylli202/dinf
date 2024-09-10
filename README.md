# Directory INFormation

`dinf` stands for directory information.
It is a (very) small CLI application that:

- Count the files in the current working directory.
  - (opt.) Count the files recursively.
  - (opt.) Output only the file count instead of a human friendly sentence.
- Size of the current working directory
  - (opt.) Size the files recursively.
  - (opt.) Output only the file count instead of a human friendly sentence

... many more features to come

## Goals

The goals of this (very) small project is to practice Golang development:
it's pattern, it's standard library.
While also practicing on developing tests at various levels.
Here both the internals and the top-level CLI is tested.

## Tests

`/cmd` tests have the responsibility to tests if the command has the expected
flags, that the shorthand version are properly translated.
They also work as integration tests.

`/internals` tests mostly tests the formatting of a function, as well as
the proper behavior for a given set of option.

`/internals/dir` tests the pure logic of the commands.
