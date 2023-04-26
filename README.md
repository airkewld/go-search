# go-search
Simple app that accepts input and opens your preferred search engine all from the CLI

## Build
Run `go run ./cicd/ci.go` to build ``linux,darwin, and windows`` binaries (Docker must be running). The binaries will be stored under `./build`.
```
build
├── darwin
│   └── go-search
├── linux
│   └── go-search
└── windows
    └── go-search.exe
```

## Usage
You can simply run ``./build/<your OS>/go-search`` to get started.
```
❯ ./build/darwin/go-search
AMA: what is today's date
```
You may also be prompted to enter your preferred search engine(currently only google and duckduckgo are supported). If you'd like to avoid entering it every time, just set `FAV_SE=duck` as an environment variable.

I mainly use this in my neovim config to start a search from within neovim.

[nvim mapping](https://github.com/airkewld/dotfiles/blob/main/nvim/after/plugin/mappings.lua#L25)

[ansible role](https://github.com/airkewld/dotfiles/blob/main/roles/go-search/tasks/main.yml)
