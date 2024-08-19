# altpkgdiff

A cli tool that will output a JSON structure that shows:

- Packages only in `p10`.
- Packages only in `sisyphus`.
- Packages with higher version in `sisyphus`.

## Project structure

altpkgdiff/
├── api/
│   └── client.go       # REST API interaction
├── cmd/
│   └── main.go         # CLI logic here
├── pkg/
│   └── diff.go         # Package diff funcs
├── go.mod              # Module
├── README.md           # Doc
└── .gitignore          # ignoring files

## Build

```sh
go build -o altpkgdiff ./cmd
```

## Use

```sh
./altpkgdiff
```
