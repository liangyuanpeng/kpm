package semver

import (
	"testing"

	"gotest.tools/v3/assert"
	"kcl-lang.io/kpm/pkg/errors"
)

func TestLatestVersion(t *testing.T) {
	latest, err := LatestVersion([]string{"1.2.3", "1.4.0", "1.3.5", "1.0.0"})
	assert.Equal(t, err, nil)
	assert.Equal(t, latest, "1.4.0")

	latest, err = LatestVersion([]string{"1.2.3", "1.3.5", "1.0.0", "1.4.0-beta"})
	assert.Equal(t, err, nil)
	assert.Equal(t, latest, "1.4.0-beta")

	latest, err = LatestVersion([]string{})
	assert.Equal(t, err, errors.InvalidVersionFormat)
	assert.Equal(t, latest, "")

	latest, err = LatestVersion([]string{"invalid_version"})
	assert.Equal(t, err.Error(), "failed to parse version invalid_version\nMalformed version: invalid_version\n")
	assert.Equal(t, latest, "")

	latest, err = LatestVersion([]string{"1.2.3", "1.4.0", "1.3.5", "invalid_version"})
	assert.Equal(t, err.Error(), "failed to parse version invalid_version\nMalformed version: invalid_version\n")
	assert.Equal(t, latest, "")
}

func TestTheLatestTagWithMissingVersion(t *testing.T) {
	latest, err := LatestVersion([]string{"1.2", "1.4", "1.3", "1.0", "5", "0.1.0-beta"})
	assert.Equal(t, err, nil)
	assert.Equal(t, latest, "5")

	latest, err = LatestVersion([]string{"1.2", "1.4", "1.3", "1.0", "5.5.5"})
	assert.Equal(t, err, nil)
	assert.Equal(t, latest, "5.5.5")

	latest, err = LatestVersion([]string{"1.2", "1.4", "1.3", "1.0", "5.5"})
	assert.Equal(t, err, nil)
	assert.Equal(t, latest, "5.5")
}
