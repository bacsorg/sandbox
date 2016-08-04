package sandbox

import (
	"os"

	"github.com/bunsanorg/try/scope"
)

type SystemConfig struct {
	BindMounts []string
	Symlinks   map[string]string
}

func GetSystemConfig() (config *SystemConfig, err error) {
	add := func(dir string) {
		finfo, err := os.Lstat(dir)
		if err != nil {
			if os.IsNotExist(err) {
				return // skip
			} else {
				scope.Must(err)
			}
		}
		if finfo.IsDir() {
			config.BindMounts = append(config.BindMounts, dir)
		} else if finfo.Mode()&os.ModeSymlink != 0 {
			link, err := os.Readlink(dir)
			scope.Must(err)
			config.Symlinks[dir] = link
		} else {
			return // skip
		}
	}
	scope.Try(func() {
		config = &SystemConfig{
			Symlinks: make(map[string]string),
		}
		add("/bin")
		add("/dev")
		add("/etc")
		add("/lib")
		add("/lib32")
		add("/lib64")
		add("/opt")
		add("/sbin")
		add("/usr")
	}).Catch(func(e error) {
		config = nil
		err = e
	})
	return
}
