<div align="center"><p>
  <img src="./media/thoughtsync-logo.PNG" alt="ThoughtSync Logo" width="500" /> <br/>
  <a href="https://github.com/Leo-Campo/ThoughtSync/releases/latest">
     <img alt="Latest release" src="https://img.shields.io/github/v/release/lcampit/ThoughtSync?style=for-the-badge&logo=starship&labelColor=302D41&include_prerelease&sort=semver" />  
  </a>
  <a href="https://github.com/Leo-Campo/ThoughtSync/pulse">
    <img alt="Last commit" src="https://img.shields.io/github/last-commit/lcampit/ThoughtSync?style=for-the-badge&logo=starship&color=8bd5ca&logoColor=D9E0EE&labelColor=302D41"/>
  </a>
  <a href="https://github.com/Leo-Campo/ThoughtSync/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/lcampit/ThoughtSync?style=for-the-badge&logo=starship&color=ee999f&logoColor=D9E0EE&labelColor=302D41" />
  </a>
  <p><em>Edit notes at the speed of thought</em></p>
</div>

# :pencil: What is this?

ThoughtSync is a CLI tool that allows users to access
their own note collection, comfortably from the command line,
using their own `EDITOR`.

It offers git support to add and commit
changes to users note, with an option to support
git remotes and pulling/pushing changes.

ThoughtSync was born to be used with an [Obsidian](https://obsidian.md/)-type
vault in mind, but it is compatible with just
any note collection contained in a
single folder. This collection is commonly referred to
as a _vault_.

# :computer: Tech Stack

ThoughtSync is built in go using the
[cobra](https://github.com/spf13/cobra) framework and a
handful of other go libraries like
[go-editor](https://github.com/confluentinc/go-editor) and
[go-git](https://github.com/go-git/go-git).
Some other tools used are:

- [mise-en-place](https://mise.jdx.dev/), a polyglot runtime manager to handle
  tool versions, env vars and much more
- [mockery](https://github.com/vektra/mockery), mock generator for go interfaces
- [testify](https://github.com/stretchr/testify), a go toolkit to write tests
- [color](https://github.com/fatih/color) to print colored text to the
  terminal
- [gotestsum](https://github.com/gotestyourself/gotestsum), a wrapper
  for the `go test` command
- [act](https://github.com/nektos/act) for testing github actions locally
- [color](https://github.com/fatih/color) to print colored
  text to the terminal

# ✨ Features

ThoughtSync lets the user open, edit and search
through their own note collection.
Current features include:

- Editing an existing note;
- Create a new note in any subdirectory in their own vault;
- Create or open the note entry for the current day (for those that keep
  a daily journal)
- Configure note vault path, journal directory, note format and much more
- Keep track of your vault changes using git and optionally sync it
  with a remote repository

# ⚙️ Roadmap

- [ ] Ability to search through notes using popular
      tools such as [ripgrep](https://github.com/BurntSushi/ripgrep)
- [ ] Quickly group notes with a certain tag or containing a certain word
- [ ] Quickly read a note content without opening it using `cat` or similar
- [x] Preferences in a single configuration file
- [x] Quickly get and set configuration options
- [x] Git syncing with remote options
- [ ] Fuzzy find notes in your vault and open them
- [x] See the vault git status
- [ ] Open a tree view of your vault for easy navigation

# :rocket: Installation

ThoughtSync can be installed as any go product with `go install`:

```bash
go install github.com/lcampit/ThoughtSync
```

or simply by downloading the executable from the releases page

It's suggested to use an alias such as `alias ts=ThoughtSync`

# 🔥 Usage

The following is a list of available commands:

- `thoughtsync new <filename.txt>` to create a new note in your vault
- `thoughtsync today` to create and/or open the journal note of today
- `thoughtsync inbox` to quickly opne an inbox note to store thoughts and ideas,
  to be processed later
- `thoughtsync open` to open the vault directory in your editor
- `thoughtysync git` contains all git-related commands:
  - `sync` to stage, commit and optionally push your vault to
    your remote git repository.
  - `status` to see the current vault git status
  - `push` manually pushes changes to the remote vault git repository
  - `pull` manually pulls changes from the remote vault git repository
- `thoughtsync config` contains all options related to the tool configuration:
  - The base commands prints the current configuration
    options in a nicely indented view
  - `set` sets a configuration options using the given value
    and saves the configuration changes

# ⚙️ Configuration

ThoughtSync can be configured using a `thoughtsync.yaml` file
contained in `$XDG_CONFIG_HOME/thoughtsync/thoughtsync.yaml`

The configuration file is defined as follows:

- The `vault` section contains all configuration related to the note vault
  - `path` contains the full path to the vault main directory
  - `extension` contains the extension of the notes files in your vault (with dot)
- The `journal` section contains all configuration related to the journal notes:
  - `directory` contains the path, relative to `vault.path`,
    to store your journal notes in
  - `format` defines the format used to give a name to your journal
    notes, such as "2006-02-01", without extension
    - currently supported formats are `YYYY-MM-DD`, `MM-DD-YYYY`
- The `git` section contains all options related to managing your
  notes with git:
  - `enable` enables git support
  - `remote` enables git remote support and
    pushing to the remote with `thoughtsync sync`
    (has no effect if `git.enable` is false)
  - `remote-name` sets the git remote name to push to / pull from
  - `commit-message` contains the commit message used with `thoughtsync git sync`
  - `ssh` to use ssh authentication for your vault git-related commands

Here's an exhaustive configuration example:

```yaml
# Contains all vault specific options
vault:
  path: $HOME/vault # Default: $HOME/thoughtsync
  extension: ".txt" # Default: ".md"

# Contains all journal specific options
journal:
  directory: my-own-journal # Default: journal
  format: "YYYY-MM-DD" # Default: "YYYY-MM-DD"

# Contains all git specific options
git:
  enable: true # default false
  remote: true # default false
  remote-name: "remote" # default origin
  commit-message: "thoughtsync sync" # default "thoughtsync: Synced with git"
  ssh: true # default false
```

> [!NOTE]
> Some configuration options need to be the expected type, i.e.
> booleans for the enable/disable
> options. If these options provided from the configuration file are not valid,
> the default one will be silently used

# :running_man: Running the project

The CLI itself has a proper help message accessible
with the `--help` or `-h` flags. If you clone the repo,
there are some pre-built tasks defined to make
the development experience easier.  
Assuming [mise](https://mise.jdx.dev/) is installed:

- `mise run build` will build the project in a single executable `ThoughtSync`
- `mise run test` will run tests

# :high_brightness: Contributing

Contributions are welcome! Feel free to clone the project
and create pull requests or add issues.

# :date: How did it start?

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
