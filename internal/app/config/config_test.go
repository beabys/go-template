package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"gitlab.com/beabys/quetzal"

	"gotest.tools/assert"
)

func TestConfig(t *testing.T) {
	t.Run("test new should return a new config", func(t *testing.T) {
		config := New()
		var want = &Config{}
		var err error
		if !reflect.DeepEqual(config, want) {
			err = fmt.Errorf("Config not equals, New() = %v, want %v", config, want)
		}
		assert.NilError(t, err)
	})

	t.Run("test Get Config should return config", func(t *testing.T) {
		new := New()
		config := new.GetConfigs()
		var err error
		if !reflect.DeepEqual(new, config) {
			err = fmt.Errorf("Config not equals, New() = %v, want %v", new, config)
		}
		assert.NilError(t, err)
	})

	t.Run("test fail Load configs", func(t *testing.T) {
		os.Setenv("CONFIG_FILE", "/path/no/exist.yaml")
		config := New()
		assert.ErrorContains(t, config.LoadConfigs(), "Fail to load configs")
		os.Unsetenv("CONFIG_FILE")
	})

	t.Run("test Success Loading configs", func(t *testing.T) {
		config := New()
		file, testPath, err := createTestConfigFile("./../../../test/")
		assert.NilError(t, err)
		os.Setenv("CONFIG_FILE", file)
		assert.NilError(t, config.LoadConfigs())
		os.Unsetenv("CONFIG_FILE")
		os.RemoveAll(testPath)
	})
}

type MockConfig struct {
	App    MockApplicationConfig `mapstructure:"application"`
	Logger MockLoggerConfig      `mapstructure:"logger"`
}
type MockApplicationConfig struct {
	Port int `mapstructure:"port"`
}
type MockLoggerConfig struct {
	LogOutput   string `mapstructure:"log_output_to"`
	ErrorOutput string `mapstructure:"log_errors_to"`
	Level       string `mapstructure:"log_level"`
}
type MockConfigFail struct {
	*Config
}

func (m *MockConfigFail) LoadConfigs() error {
	return errors.New("Fail loading configs")
}

func createTestConfigFile(path string) (string, string, error) {
	testpath := path + quetzal.RandomString(8)
	config := testpath + "/env.local.json"
	err := os.MkdirAll(testpath, os.ModePerm)
	if err != nil {
		return "", testpath, err
	}
	data := `{"application": {"port": 8080}}`
	createFile(config, data)
	return config, testpath, nil
}

func createFile(path, data string) {
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		fmt.Printf("error creating file")
	}
}
