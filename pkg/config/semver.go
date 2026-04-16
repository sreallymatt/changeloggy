package config

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/Masterminds/semver/v3"
)

// headingVersionPattern matches a semver on any markdown heading line.
// tODO: configurable?
var headingVersionPattern = regexp.MustCompile(`^#+\s.*?v?(\d+\.\d+\.\d+)`)

func (c *Config) NextVersion() (string, error) {
	latest, err := c.latestVersion()
	if err != nil {
		return "", err
	}

	bumped, err := bumpVersion(latest, c.DefaultVersionIncrement)
	if err != nil {
		return "", err
	}

	// TODO: option to omit `v` prefix
	return "v" + bumped.String(), nil
}

func (c *Config) latestVersion() (*semver.Version, error) {
	cl := c.ChangelogFilePath()

	if _, err := os.Stat(cl); os.IsNotExist(err) {
		v, _ := semver.NewVersion("0.0.0")
		return v, nil
	}

	f, err := os.Open(cl)
	if err != nil {
		return nil, fmt.Errorf("opening changelog (%s): %w", cl, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if m := headingVersionPattern.FindStringSubmatch(line); len(m) >= 2 {
			v, err := semver.NewVersion(m[1])
			if err == nil {
				return v, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading changelog (%s): %w", cl, err)
	}

	v, _ := semver.NewVersion("0.0.0")
	return v, nil
}

func bumpVersion(v *semver.Version, increment string) (semver.Version, error) {
	switch increment {
	case "major":
		return v.IncMajor(), nil
	case "minor":
		return v.IncMinor(), nil
	case "patch":
		return v.IncPatch(), nil
	default:
		return semver.Version{}, fmt.Errorf("invalid `default_version_increment` (`%s`): must be `major`, `minor`, or `patch`", increment)
	}
}
