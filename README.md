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

set `GOOS` & `GOARCH` to your target platform

```pwsh
$env:GOOS="linux"
$env:GOARCH="amd64" 
```

run build script

```sh
go build -o altpckgdiff .\cmd\main.go
```

## Use

make file executable

```sh
chmod +x altpckgdiff
```

execute programm

```pwsh
./altpckdiff
```

save results to txt

```sh
./altpkgdiff > output.txt
```
