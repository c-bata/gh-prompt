# gh-prompt

An interactive GitHub CLI based on https://cli.github.com/

## Installation

### Downloading standalone binary

Binaries are available from [github release](https://github.com/c-bata/gh-prompt/releases).

<details>
<summary>macOS (darwin) - amd64</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.1.0/gh-prompt_v0.1.0_darwin_amd64.zip
unzip gh-prompt_v0.1.0_darwin_amd64.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>

<details>
<summary>Linux - amd64</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.1.0/gh-prompt_v0.1.0_linux_amd64.zip
unzip gh-prompt_v0.1.0_linux_amd64.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>


<details>
<summary>Linux - i386</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.1.0/gh-prompt_v0.1.0_linux_386.zip
unzip gh-prompt_v0.1.0_linux_386.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>

<details>
<summary>Linux - arm64</summary>

```
wget https://github.com/c-bata/gh-prompt/releases/download/v0.1.0/gh-prompt_v0.1.0_linux_arm64.zip
unzip gh-prompt_v0.1.0_linux_arm64.zip
chmod +x gh-prompt
sudo mv ./gh-prompt /usr/local/bin/gh-prompt
```

</details>


### Goal

Hopefully support following commands enough.

* [x] issue
    * [x] create
    * [x] list
    * [x] status
    * [x] view
* [x] pr
    * [x] checkout
    * [x] create
    * [x] list
    * [x] status
    * [x] view

## LICENSE

This software is licensed under the MIT License (See [LICENSE](./LICENSE)).
