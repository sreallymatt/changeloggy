# changeloggy

Test from fork

A CLI tool for managing changelog entries with optional entry format validation.

## Installation

```sh
go install github.com/sreallymatt/changeloggy@latest
```

## Usage

Initialize a config file in your repository:

```sh
changeloggy config init
```

Add a changelog entry for a pull request:

```sh
changeloggy add --pr 42 --type example "Feat: an example entry"
```

Validate entries:

```sh
changeloggy check --pr 42 # validate a single PR entry
changeloggy check          # validate all entries
```

Generate the changelog:

```sh
changeloggy generate --version 1.2.0
changeloggy generate # auto-bumps version based on `default_version_increment`
```

## Configuration

Run `changeloggy config init` to create a `.changeloggy.hcl` config file. Key options:

| Option                      | Description                                                                                            |
|-----------------------------|--------------------------------------------------------------------------------------------------------|
| `changelog_file`            | Path to the changelog file                                                                             |
| `default_version_increment` | `major`, `minor`, or `patch`                                                                           |
| `archive_entries`           | If `true`, moves entries to `archive_path` after generation. If `false` or unset, entries are deleted. |
| `entries_path`              | Directory where entry files are stored (default: `.changelog/`)                                        |
