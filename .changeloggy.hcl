changelog_file            = "CHANGELOG.md"
default_version_increment = "minor"

format {
  date = "January 2, 2006"

  template = <<-TEMPLATE
## {{ .Version }} ({{ .Date }})
{{ range .Kinds }}
{{ .Heading }}:

{{ range .Entries }}- {{ .Body }} ([#{{ .PR }}](https://github.com/sreallymatt/changeloggy/pull/{{ .PR }}))
{{ end }}{{ end }}
TEMPLATE
}

kind "improvement" {
  heading  = "Improvements"
  priority = 1

  type "cli" {
    regex    = "^cli: `[^`]+` - .*$"
    example  = "cli: `config init` - add support for a new `--force` flag"
    priority = 1
  }

  type "internal" {
    regex    = "^internal: `[^`]+` - .*$"
    example  = "internal: `templatehelper` - add a new template helper to make templating even better"
    priority = 2
  }
}

kind "bug" {
  heading = "Bugs"
  priority = 2

  type "cli-bug" {
    regex    = "^cli: `[^`]+` - fix .*$"
    example  = "cli: `generate` - fix an issue that prevented generating the changelog during an active solar eclipse"
    priority = 1
  }

  type "generic-bug" {
    regex    = "^fix .*$"
    example  = "fix a panic caused by dark magic"
    priority = 2
  }
}

kind "dependencies" {
  heading  = "Dependencies"
  priority = 3

  type "deps" {
    regex    = "^`[^`]+` has been updated from `[^`]+` to `[^`]+`$"
    example  = "`github.com/hashicorp/go-azure-helpers` has been updated from `v0.41.0` to `v0.42.0`"
    priority = 1
  }
}
