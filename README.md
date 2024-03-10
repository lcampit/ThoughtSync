
<div align="center"><p>
  <a href="https://github.com/Leo-Campo/ThoughtSync/releases/latest">
     <img alt="Latest release" src="https://img.shields.io/github/v/release/Leo-Campo/ThoughtSync?style=for-the-badge&logo=starship&labelColor=302D41&include_prerelease&sort=semver" />  
  </a>
  <a href="https://github.com/Leo-Campo/ThoughtSync/pulse">
    <img alt="Last commit" src="https://img.shields.io/github/last-commit/Leo-Campo/ThoughtSync?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
  </a>
  <a href="https://github.com/Leo-Campo/ThoughtSync/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/Leo-Campo/ThoughtSync?style=for-the-badge&logo=starship&color=ee999f&logoColor=D9E0EE&labelColor=302D41" />
  </a>
</div>

# Thought Sync

## :pencil: Introduction

ThoughtSync is a CLI tool that allows users to access
their own note collection, comfortably from the command line,
using their own `EDITOR`.

It was born to be used with an [Obsidian](https://obsidian.md/)-type
vault in mind, but it is compatible with just
any note collection contained in a
single vault-like folder.

## :computer: Tech Stack

ThoughtSync is built in go using the
[cobra](https://github.com/spf13/cobra) library.
Some other tools used are:

- [mise-en-place](https://mise.jdx.dev/), a polyglot runtime manager to handle
  tool versions, env vars and much more
- [mockery](https://github.com/vektra/mockery), mock generator for go interfaces
- [testify](https://github.com/stretchr/testify), a go toolkit to write tests
- [gotestsum](https://github.com/gotestyourself/gotestsum), a wrapper
  for the `go test` command

## ✨ Features

ThoughtSync lets the user open, edit and search
through their own note collection.
Current features include:

- Editing an existing note;
- Create a new note in any subdirectory in their own vault;
- Create or open the note entry for the current day (for those that keep
  a daily journal)

### ⚙️ Roadmap

- [ ] Ability to search through notes using popular
      tools such as [ripgrep](https://github.com/BurntSushi/ripgrep)
- [ ] Quickly group notes with a certain tag or containing a certain word
- [ ] Quickly read a note content without opening it using `cat` or similar
- [ ] Preferences in a single configuration file
  - [ ] Configure directories (folders) in your vault and dinamycally
        create commands for them
- [ ] Fuzzy find notes in your vault and open them

## :rocket: Installation

ThoughtSync can be installed as any go product with `go install`:

```bash
go install github.com/Leo-Campo/ThoughtSync
```
or simply by downloading the executable from the releases page

### 🔥 Usage

All commands inherit the `--vault` flag which
points to the main directory of your note
vault, i.e. the directory containing all your notes.
If you don't want to add this flag to each and every command,
you can set the `THOUGHT_SYNC_NOTE_VAULT` environment variable which will
be automatically read by the CLI.

Once that is done, simply run:

- `thoughtsync new <filename.txt>` to create a new note in your vault
- `thoughtsync today` to create and/or open the journal note of today

### :running_man: Running the project

The CLI itself has a proper help message accessible
with the `--help` or `-h` flags. If you clone the repo,
there are some pre-built tasks defined to make
the development experience easier.  
Assuming [mise](https://mise.jdx.dev/) is installed:

- `mise run build` will build the project in a single executable `ThoughtSync`
- `mise run test` will run tests

## :high_brightness: Contributing

Contributions are welcome! Feel free to clone the project
and create pull requests or add issues.

## :date: How it started

I've been caressing the idea of building my own second brain
with notes since I had a job. Being able to write down stuff
helps me keep a clear head, so I started building my own second
brain using [Notion](https://www.notion.so/).
I've always wanted to be fluent with the command line, but I could not
find proper ways to access my notes in Notion through the CLI, at least
until I found out [this](https://www.youtube.com/watch?v=zIGJ8NTHF4k)
video that showed me how to combine Neovim
with Note taking in [Obsidian](https://obsidian.md/). I made the switch,
but I noticed how the video author used bash scripts to access its own
vault. Surely there is some CLI tool I can use for that, right? no? Time to
create my own then!
