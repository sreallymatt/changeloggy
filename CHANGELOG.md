## v0.3.0 (September 21, 2026)

Improvements:

- cli: `add` - entries can now be added in `hcl`, `yml`, or `md` formats ([#34](https://github.com/sreallymatt/changeloggy/pull/34))
- cli: `check` - checks for entries in `hcl`, `yml`, or `md` formats ([#34](https://github.com/sreallymatt/changeloggy/pull/34))
- cli: `generate` - parses entries in `hcl`, `yml`, or `md` formats to generate the changelog ([#34](https://github.com/sreallymatt/changeloggy/pull/34))

Bugs:

- cli: `types` - fix example output for the `--table` flag, preserving the raw string by wrapping it backticks ([#36](https://github.com/sreallymatt/changeloggy/pull/36))

## v0.2.0 (September 18, 2026)

Improvements:

- cli: `add` - entry validation now prints more detail on error ([#28](https://github.com/sreallymatt/changeloggy/pull/28))
- cli: `types` - add support for the `--table` flag ([#11](https://github.com/sreallymatt/changeloggy/pull/11))

Bugs:

- cli: `add` - fix new entry file write to no longer include an initial empty line ([#29](https://github.com/sreallymatt/changeloggy/pull/29))

## v0.1.0 (September 8, 2026)

Initial release

Features:

* `changeloggy add` - create or append to a changelog entry file for a pull request
* `changeloggy check` - validate a single entry file (`--pr`) or all entries in the entries directory
* `changeloggy generate` - render all entries into a changelog file, with automatic semver version bumping
* `changeloggy types` - list all configured kinds and entry types
* `changeloggy config init` - scaffold a default `.changeloggy.hcl` configuration file
* `changeloggy config validate` - validate an existing configuration file
* Entry format validation via per-type regex patterns
* Configurable changelog template via Go `text/template`
* Entry archiving or automatic cleanup after generation (`archive_entries`)
