package configloader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"

	"github.com/justclimber/fda/common/config"
)

const appConfigFileName = "app.toml"

var ErrAppConfigNotFound = errors.New("app config file not found")

type ConfigLoader struct{}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (c *ConfigLoader) Load() (config.Config, error) {
	var appCfg config.Config
	dir, err := findAppConfigDirPath()
	if err != nil {
		return config.Config{}, err
	}

	fileData, err := ioutil.ReadFile(filepath.Join(dir, appConfigFileName))

	if err != nil {
		return config.Config{}, fmt.Errorf("can't read config file: %w", err)
	}

	if _, err = toml.Decode(string(fileData), &appCfg); err != nil {
		return config.Config{}, fmt.Errorf("can't decode config file: %w", err)
	}
	appCfg.SetBaseDir(dir)
	return appCfg, nil
}

func findAppConfigDirPath() (string, error) {
	_, from, _, _ := runtime.Caller(1)
	dir := filepath.Dir(from)
	gopath := filepath.Clean(os.Getenv("GOPATH"))
	for dir != "/" && dir != gopath {
		appTomlFile := filepath.Join(dir, appConfigFileName)
		if _, err := os.Stat(appTomlFile); os.IsNotExist(err) {
			dir = filepath.Dir(dir)
			continue
		}
		return dir, nil
	}
	return "", ErrAppConfigNotFound
}
