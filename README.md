# gh-prompt

![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/c-bata/gh-prompt?status.svg)](https://godoc.org/github.com/c-bata/gh-prompt)


An interactive GitHub CLI featuring auto-complete. This tool provides powerful completion to GitHub's official CLI.

[![GIF animation](https://github.com/c-bata/assets/raw/master/gh-prompt/gh-prompt.gif)](#)

You can walk through issues, create pull requests, checkout pull requests locally, and more.
See https://cli.github.com/ for details.

## Installation

### Homebrew (for macOS users)

```
$ brew install c-bata/gh-prompt/gh-prompt
```

### Downloading standalone binary

Binaries are available from [github release](https://github.com/c-bata/gh-prompt/releases).

<details>
<summary>macOS (darwin) - amd64</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.0.1/gh-prompt_darwin_x86_64.zip
unzip gh-prompt_darwin_x86_64.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>

<details>
<summary>Linux - amd64</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.0.1/gh-prompt_linux_x86_64.zip
unzip gh-prompt_linux_x86_64.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>


<details>
<summary>Linux - i386</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.0.1/gh-prompt_linux_i386.zip
unzip gh-prompt_linux_i386.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>

### Building from source

```
$ git clone git@github.com:c-bata/gh-prompt.git
$ cd gh-prompt
$ make build
```

You can create multi-platform binaries via goreleaser:

```
$ goreleaser --snapshot --skip-publish --rm-dist
```

## LICENSE

This software is licensed under the MIT License (See [LICENSE](./LICENSE)).
