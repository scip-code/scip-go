# scip-go

SCIP indexer for Golang.

# Quick Start

## Installation

This will build and install the latest version of `scip-go`

    go install github.com/scip-code/scip-go/cmd/scip-go@latest

You can confirm it's been installed by running:

    scip-go --version

## Indexing a Go project

### Standard `go.mod` project

From the root of your project, you can run:

    scip-go

If `scip-go` is unable to determine some project information, you may need to
provide some command-line arguments.

    scip-go --module-name=NAME --module-version=VERSION

If this doesn't solve the problem, check the rest of the available flags in:

    scip-go --help

### Other build systems

The other build systems Buck/Bazel/Please/etc are supported via [Go Packages
Driver Protocol].

Usage:

    GOPACKAGESDRIVER=your_driver scip-go

Note: Due to the current protocol design cross-repo navigation will not work.

### Common Problems:

- Unable to navigate to Go standard library.
  - To solve this, you may want to use the `--go-version=go1.X.Y` flag when
    indexing and then also index the go versions manually.
  - To index the Go standard library, you'll want to check out
    https://github.com/golang/go, checkout the tag you want, navigate to the
    `src/` directory, and then run:
    - `$ scip-go --go-version=go1.X.Y` and then upload that index to your local
      sourcegraph instance.
    - After you've done this, you should be able to navigate to the standard
      library.

(NOTE: Projects without a `go.mod` may experience challenges indexing. See next
section for details)

## Indexing without shelling to `go` binary

`scip-go` by default uses a few different `go` commands from the command line to
gain information about the project and module. To avoid running `go` directly
(perhaps you have some other build system), you will need to supply the
following args.

    scip-go --module-name="<my modules name here>"

NOTE: The rest of this isn't properly implemented yet. It's on the todo list for
scip-go.

## Indexing in CI

    # Install scip-go
    go install github.com/scip-code/scip-go/cmd/scip-go@latest

    # Run scip-go
    scip-go

    # Upload index with any necessary tokens (shown here using GitHub workflow syntax)
    src code-intel upload -github-token='${{ secrets.GITHUB_TOKEN }}' -no-progress

## Docker

To build a self-contained Docker image with the indexer, use:

``` bash
nix build .#docker
docker load < result
```

# Development

## Run tests

Run all tests:

``` bash
go test ./...
```

Update snapshot tests:

``` bash
go test ./internal/index -update-snapshots
```

## Cutting releases

Every 4 weeks on Monday, the [scheduled-release workflow] opens a PR that
bumps the patch version in `internal/index/version.txt` and updates
`flake.lock`, so the Docker image picks up nixpkgs security fixes even
without code changes. To open one now, run it via **Actions →
scheduled-release → Run workflow**. For other releases, land a PR that
updates `internal/index/version.txt`.

When a version change lands on `main`, the [release workflow] tags it and
publishes the binaries and Docker images.

# Contributing

Contributors should follow the [Sourcegraph Community Code of Conduct].

  [Go Packages Driver Protocol]: https://pkg.go.dev/golang.org/x/tools/go/packages#hdr-The_driver_protocol
  [release workflow]: https://github.com/scip-code/scip-go/actions/workflows/release.yml
  [scheduled-release workflow]: https://github.com/scip-code/scip-go/actions/workflows/scheduled-release.yml
  [Sourcegraph Community Code of Conduct]: https://handbook.sourcegraph.com/company-info-and-process/community/code_of_conduct/
