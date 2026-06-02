# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repository is

A personal worktree for the book **"Learn Go with Tests"** (https://github.com/quii/learn-go-with-tests), a test-driven introduction to Go. The point is the *process*, not the destination: each topic is learned by writing a failing test first, watching it fail, then writing the minimum code to pass, then refactoring.

- Module: `learn-go-with-tests`
- Go: 1.26.3

The repo starts empty. Code accrues one package per chapter/concept as the book is worked through (e.g. `hello/`, `integers/`, `iteration/`, `arrays/`, ...). Each package is self-contained: an implementation file plus its `_test.go`.

## Commands

```bash
go test ./...                 # run every package's tests
go test ./hello/              # run one package
go test -run TestHello        # run tests matching a name (regex)
go test -v ./hello/           # verbose: show each subtest
go test -cover ./...          # coverage summary
go test -bench=. ./...        # run benchmarks (book uses these in the iteration chapter)
go vet ./...                  # static checks
gofmt -w .                    # format (or rely on goimports/GoLand on save)
```

Single-test loop while learning: `go test -run TestName -v ./<pkg>/`.

## Working conventions (TDD discipline the book teaches)

Follow the book's red-green-refactor cycle deliberately — do not skip ahead to a finished implementation:

1. **Write the test first**, with the assertion you wish were true.
2. **Run it and confirm it fails** for the expected reason (and that the failure message is meaningful). A test that has never failed isn't trusted.
3. **Write the minimum code** to make it compile and pass — even if that means returning a hardcoded value at first.
4. **Refactor** with the test as a safety net.

Notes that match the book's style:
- Test helpers take `t testing.T` (or `t *testing.T`) and call `t.Helper()` so failures report the caller's line.
- Prefer table-driven tests (`[]struct{...}` + `t.Run(name, ...)`) once a function has several cases.
- Keep example functions (`ExampleXxx`) where the book introduces them — they double as documentation and are run by `go test`.