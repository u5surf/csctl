# csctl

[![Build Status](https://travis-ci.org/containership/csctl.svg?branch=master)](https://travis-ci.org/containership/csctl)
[![codecov](https://codecov.io/gh/containership/csctl/branch/master/graph/badge.svg)](https://codecov.io/gh/containership/csctl)

`csctl` is a command line interface for Containership Cloud.

This repository also contains the go client for interacting with Containership Cloud.
For more info, please refer to [the client documentation](cloud/README.md)

**Warning**: This project is currently under active development and is subject to breaking changes without notice.

## Installing

Currently the only way to install `csctl` is through go:

```
go get -u github.com/containership/csctl
```

Alternatively, you can clone this repository and install via `make`:

```
make install
```

More installation methods will be added in the future as the project matures.
There are no official releases yet.

## Usage

Please use `csctl -h` for now to discover usage.

More documentation will be added as the project matures.

### Configuration

`csctl` defaults to a config file located at `~/.containership/csctl.yaml`.
You may also choose to manually specify a config file using `--config`.

The only config option to set today is `token`.
Because the CLI does not yet support authorization, you must manually specify a `token`.
Your user token can be obtained by using the developer tools of your favorite browser to find the `JWT` for an API request.
Authorization is coming soon, but feel free to reach out with any questions in the meantime.

## Contributing

Thank you for your interest in this project and for your interest in contributing!
Feel free to open issues for feature requests, bugs, or even just questions - we love feedback and want to hear from you.

PRs are also always welcome!
However, if the feature you're considering adding is fairly large in scope, please consider opening an issue for discussion first.
