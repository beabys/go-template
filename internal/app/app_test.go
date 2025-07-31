package app

import (
	"fmt"
	"testing"

	"github.com/beabys/go-template/internal/app/config"
	mockconfig "github.com/beabys/go-template/mocks/app/config"
	"github.com/beabys/go-template/pkg/logger"
	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {

	t.Run("test new should return a new app", func(t *testing.T) {
		app := New()
		var want *App
		assert.IsType(t, app, want)
	})

	t.Run("test fail Set configs", func(t *testing.T) {
		app := New()
		mockConfigs := mockconfig.NewAppConfig(t)
		mockConfigs.On("LoadConfigs").Return(fmt.Errorf("fails"))
		err := app.SetConfigs(mockConfigs)
		assert.ErrorContains(t, err, "fails")
	})

	t.Run("test Success Set configs", func(t *testing.T) {
		app := New()
		mockConfigs := mockconfig.NewAppConfig(t)
		mockConfigs.On("LoadConfigs").Return(nil)
		err := app.SetConfigs(mockConfigs)
		assert.NoError(t, err)
	})

	t.Run("test Set Logger", func(t *testing.T) {
		app := New()
		logger, _ := logger.NewZapLogger([]string{}, []string{}, zap.DebugLevel)
		app.SetLogger(logger)
		assert.Equal(t, app.GetLogger(), logger)

	})
	t.Run("test fail Setup", func(t *testing.T) {
		app := New()
		mockConfig := mockconfig.NewAppConfig(t)
		mockConfig.On("LoadConfigs").Return(fmt.Errorf("make it fail"))
		assert.ErrorContains(t, app.Setup(mockConfig), "make it fail")
	})

	t.Run("test Success Setup", func(t *testing.T) {
		app := New()
		mockConfig := mockconfig.NewAppConfig(t)
		mockConfig.On("LoadConfigs").Return(nil)
		mockConfig.On("GetConfigs").Return(&config.Config{})
		assert.NoError(t, app.Setup(mockConfig))
	})
}

func TestRecover(t *testing.T) {
	t.Run("test Recover function", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				assert.Equal(t, r, "this message means function panic and go into the recoverer")
				return
			}
			t.Errorf("The code did not panic")
		}()
		app := New()
		mockLogger, _ := logger.NewZapLogger([]string{}, []string{}, zap.DebugLevel)
		logger := &MockLogger{mockLogger}
		app.Logger = logger
		app.Recoverer(WithPanic)
	})
}

func WithPanic() {
	panic("panic")
}

type MockLogger struct {
	*logger.ZapLogger
}

func (m *MockLogger) Warn(string, ...logger.LogField) {
	// using the error function to panic during the recovery
	// in that case means recovery is going inside properly
	panic("this message means function panic and go into the recoverer")
}
