package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/credentials"
	"github.com/docker/cli/cli/config/types"
	"github.com/docker/docker/pkg/homedir"
	"github.com/pkg/errors"
)

const (
	// ConfigFileName is the name of config file
	ConfigFileName = "config.json"
	configFileDir  = ".docker"
	oldConfigfile  = ".dockercfg"
	contextsDir    = "contexts"
)

var (
	initConfigDir sync.Once
	configDir     string
	homeDir       string
)

// resetHomeDir is used in testing to resets the "homeDir" package variable to
// force re-lookup of the home directory between tests.
func resetHomeDir() {
	homeDir = ""
}

func getHomeDir() string {
	if homeDir == "" {
		homeDir = homedir.Get()
	}
	return homeDir
}

func setConfigDir() {
	if configDir != "" {
		return
	}
	configDir = os.Getenv("DOCKER_CONFIG")
	if configDir == "" {
		configDir = filepath.Join(getHomeDir(), configFileDir)
	}
}

// Dir returns the directory the configuration file is stored in
func Dir() string {
	initConfigDir.Do(setConfigDir)
	return configDir
}

// ContextStoreDir returns the directory the docker contexts are stored in
func ContextStoreDir() string {
	return filepath.Join(Dir(), contextsDir)
}

// SetDir sets the directory the configuration file is stored in
func SetDir(dir string) {
	configDir = filepath.Clean(dir)
}

// Path returns the path to a file relative to the config dir
func Path(p ...string) (string, error) {
	path := filepath.Join(append([]string{Dir()}, p...)...)
	if !strings.HasPrefix(path, Dir()+string(filepath.Separator)) {
		return "", errors.Errorf("path %q is outside of root config directory %q", path, Dir())
	}
	return path, nil
}

// LegacyLoadFromReader is a convenience function that creates a ConfigFile object from
// a non-nested reader
func LegacyLoadFromReader(configData io.Reader) (*configfile.ConfigFile, error) {
	configFile := configfile.ConfigFile{
		AuthConfigs: make(map[string]types.AuthConfig),
	}
	err := configFile.LegacyLoadFromReader(configData)
	return &configFile, err
}

// LoadFromReader is a convenience function that creates a ConfigFile object from
// a reader
func LoadFromReader(configData io.Reader) (*configfile.ConfigFile, error) {
	configFile := configfile.ConfigFile{
		AuthConfigs: make(map[string]types.AuthConfig),
	}
	err := configFile.LoadFromReader(configData)
	return &configFile, err
}

// Load reads the configuration files in the given directory, and sets up
// the auth config information and returns values.
// FIXME: use the internal golang config parser
func Load(configDir string) (*configfile.ConfigFile, error) {
	if configDir == "" {
		configDir = Dir()
	}

	filename := filepath.Join(configDir, ConfigFileName)
	configFile := configfile.New(filename)

	// Try happy path first - latest config file
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		err = configFile.LoadFromReader(file)
		if err != nil {
			err = errors.Wrap(err, filename)
		}
		return configFile, err
	} else if !os.IsNotExist(err) {
		// if file is there but we can't stat it for any reason other
		// than it doesn't exist then stop
		return configFile, errors.Wrap(err, filename)
	}

	// Can't find latest config file so check for the old one
	filename = filepath.Join(getHomeDir(), oldConfigfile)
	if file, err := os.Open(filename); err == nil {
		defer file.Close()
		if err := configFile.LegacyLoadFromReader(file); err != nil {
			return configFile, errors.Wrap(err, filename)
		}
	}
	return configFile, nil
}

// LoadDefaultConfigFile attempts to load the default config file and returns
// an initialized ConfigFile struct if none is found.
func LoadDefaultConfigFile(stderr io.Writer) *configfile.ConfigFile {
	configFile, err := Load(Dir())
	if err != nil {
		fmt.Fprintf(stderr, "WARNING: Error loading config file: %v\n", err)
	}
	if !configFile.ContainsAuth() {
		configFile.CredentialsStore = credentials.DetectDefaultStore(configFile.CredentialsStore)
	}
	return configFile
}
