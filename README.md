# Проект: Анализатор размера диска (Go)

[![Hexlet tests and linter status](https://github.com/adammilligan/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions)
[![Go tests](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml)

## Что это за программа
Утилита для подсчёта размера файла или директории в байтах, с возможностью:
- учитывать скрытые файлы и директории;
- считать рекурсивно (включая вложенные папки);
- выводить размер в человекочитаемом формате.

Пример работы утилиты:

![Демонстрация работы hexlet-path-size](./assets/hexlet-path-size.gif)

## Как запустить
Собрать бинарник:

```bash
make build
```

Запустить:

```bash
./bin/hexlet-path-size <path> [flags]
```

## Формат вывода
Программа печатает одну строку в формате:

```text
<размер>\t<путь>
```

Табуляция (`\t`) используется как разделитель.

## Флаги
Флаги влияют на то, как считается размер:

- `--human`, `-H` — human-readable размеры (автоподбор единиц: `B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB`).
- `--all`, `-a` — учитывать скрытые файлы и директории (имена начинаются с `.`).
- `--recursive`, `-r` — рекурсивно обходить директории (учитываются все вложенные файлы/папки).

Примеры:

```bash
./bin/hexlet-path-size data.csv
./bin/hexlet-path-size output.dat -H
./bin/hexlet-path-size project/ -a -H -r
```