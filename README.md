# altpkgdiff

A cli tool that output a JSON structure:

- Packages only in `p10`.
- Packages only in `sisyphus`.
- Packages with higher version in `sisyphus`.

## Project structure

```sh
altpkgdiff/
├── api/
│   └── client.go       # REST API interaction
├── cmd/
│   └── main.go         # CLI logic here
├── pkg/
│   └── diff.go         # Package diff funcs
├── go.mod              # Module
├── README.md           # Doc
└── .gitignore          # list of ignoring files
```

## Build

```sh
go build -o altpkgdiff ./cmd
```

## Use

```sh
./altpkgdiff
```
