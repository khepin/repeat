# repeat

Run a command multiple times.

## Installation

```
go install github.com/khepin/repeat@latest
```

## Usage

```
repeat <N> <command> [args...]
```

- `N` — number of times to run the command
- `command` — the command to execute
- `args` — optional arguments passed to the command

If the command fails, `repeat` stops immediately and exits with the same exit code.

On completion, it prints how many runs succeeded:

```
repeated successfully 3 / 3 times
```

## Examples

```sh
# Run tests 10 times to check for flakiness
repeat 10 go test ./...

# Echo hello 3 times
repeat 3 echo hello
```
