# Smash - Secure Deletion Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/bwoff11/smash)](https://goreportcard.com/report/github.com/bwoff11/smash)
[![Go Reference](https://pkg.go.dev/badge/github.com/bwoff11/smash.svg)](https://pkg.go.dev/github.com/bwoff11/smash)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bwoff11/smash)

Smash is a command-line tool for securely deleting files and directories from your file system. It overwrites files with different patterns before removing them to ensure that the data cannot be easily recovered.

## Installation

```bash
go install github.com/bwoff11/smash@latest
```

## Usage

Basic usage to delete a file or directory:

```bash
smash /path/to/file
smash /path/to/directory
```

The tool will ask for a confirmation before deleting the file or directory. To bypass the confirmation, use the `-s` or `--silent` flag:

```bash
smash -s /path/to/file
```

To change the number of overwriting loops, use the `-c` or `--count` flag:

```bash
smash -c 5 /path/to/file
```

### Flags

- `-s`, `--silent`: Run silently without asking for confirmation
- `-c`, `--count [NUMBER]`: Number of overwriting loops, default is 5

## License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for details.