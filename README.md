# Find (and) Replace
## A command line utility written in Go
![Unit Testing](https://github.com/joshuaschlichting/findreplace/actions/workflows/go.yml/badge.svg)

The application's purpose is a bit self explanatory given the nomenclature. It accepts 3 command line arguments, a filepath (or directory), a phrase to find, and a phrase to replace it with. (Entering a directory will cause a recursive find and replace function on all text files within the directory and children directories recursively.)

## Usage
```
./findreplace \
    --path /tmp/someDir \
    --find findthistext \
    --replace replacewiththis
```
