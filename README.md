
<div align="center">

[![Release](https://img.shields.io/github/release/taubyte/tau-cli.svg)](https://github.com/taubyte/tau-cli/releases)
[![License](https://img.shields.io/github/license/taubyte/tau-cli)](LICENSE)
</div>

Tau CLI is Taubyte’s command-line tool for managing projects, resources, and cloud workflows.

## Installation

### npm

```bash
npm i @taubyte/cli
```

The published package downloads the `tau` binary for your platform from [tau-cli releases](https://github.com/taubyte/tau-cli/releases) when you run the `tau` command.

### Source

Implementation lives in the [`tau`](https://github.com/taubyte/tau) repository under [`tools/tau`](https://github.com/taubyte/tau/tree/main/tools/tau).

This repository is a **distribution** layout (similar to [dream](https://github.com/taubyte/dream)): it carries the npm wrapper, release automation, and a **`tau` git submodule** pinned to the commit used for builds.

### Working from a git clone

```bash
git clone https://github.com/taubyte/tau-cli.git
cd tau-cli
git submodule update --init --recursive
```

The submodule is required if you run [GoReleaser](https://goreleaser.com/) locally or verify a release build; the npm package published to the registry does not ship the submodule (see `.npmignore`).

## Usage

Run `tau --help` after installation for current commands and options.
