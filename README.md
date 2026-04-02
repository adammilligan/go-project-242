# Проект: Анализатор размера диска (Go)

[![Hexlet tests and linter status](https://github.com/adammilligan/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions)
[![Go tests](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/adammilligan/go-project-242/actions/workflows/go-tests.yml)

## Описание
Утилита для подсчёта размера файла или директории в байтах.

## Возможности
- учитывать скрытые файлы и директории;
- считать рекурсивно (включая вложенные папки);
- выводить размер в человекочитаемом формате.

## Быстрый старт
1. Собрать бинарник:

```bash
make build
```

2. Запустить:

```bash
./bin/hexlet-path-size <path> [flags]
```

## Вывод
Программа печатает одну строку в формате:

```text
<размер>\t<путь>
```

## Флаги
Флаги управляют подсчётом размера:

- `--human`, `-H` — выводить человекочитаемые размеры (автоподбор единиц: `B`, `KB`, `MB`, `GB`, `TB`, `PB`, `EB`).
- `--all`, `-a` — учитывать скрытые файлы и директории (имена начинаются с `.`).
- `--recursive`, `-r` — включать подпапки при подсчёте размера директории. Без `-r` учитываются только файлы в указанной папке.

## Примеры
```bash
./bin/hexlet-path-size data.csv
./bin/hexlet-path-size output.dat -H
./bin/hexlet-path-size project/ -a -H -r
```

## Демонстрация
![Демонстрация работы hexlet-path-size](./assets/hexlet-path-size.gif)