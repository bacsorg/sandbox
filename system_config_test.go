package sandbox

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSystemConfig(t *testing.T) {
	config, err := GetSystemConfig()
	require.NoError(t, err)
	for _, bind := range config.BindMounts {
		stat, err := os.Lstat(bind)
		if assert.NoError(t, err) {
			assert.True(t, stat.IsDir())
		}
	}
	for from, to := range config.Symlinks {
		link, err := os.Readlink(from)
		if assert.NoError(t, err) {
			assert.Equal(t, link, to)
		}
	}
}
