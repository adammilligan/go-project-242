# Project: Disk usage analyzer (Go)

[![Hexlet tests and linter status](https://github.com/adammilligan/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions)
[![Go tests](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml)

CLI utility that prints the size of a file or directory.

## Features
- optionally include hidden files and directories;
- optionally traverse directories recursively;
- optionally print sizes in a human-readable format.

## Quick start
1. Build:

```bash
make build
```

2. Run:

```bash
./bin/hexlet-path-size <path> [flags]
```

## Output
The program prints a single line:

```text
<размер>\t<путь>
```

## Flags

- `--human`, `-H` — print human-readable sizes (units: `B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB`)
- `--all`, `-a` — include hidden files and directories (names starting with `.`)
- `--recursive`, `-r` — include nested directories when calculating directory size

## Examples
```bash
./bin/hexlet-path-size data.csv
./bin/hexlet-path-size output.dat -H
./bin/hexlet-path-size project/ -a -H -r
```

## Demo
![hexlet-path-size demo](./assets/hexlet-path-size.gif)