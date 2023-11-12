package app

import (
	"fmt"
	"os"

	pkgerrors "github.com/pkg/errors"
)

// Config holds basic information about the app
type Config struct {
	Version string // mandatory
	Env     Env    // mandatory
	Server  string

	Project   string        // mandatory
	Component ComponentType // mandatory
	Team      string

	AppName string // mandatory
}

// IsValid checks if the config is valid or not
func (c Config) IsValid() error {
	if c.Version == "" {
		return fmt.Errorf("version missing: %w", ErrInvalidAppConfig)
	}
	if !c.Env.Valid() {
		return fmt.Errorf("invalid env: %s given: %w", c.Env.String(), ErrInvalidAppConfig)
	}
	if c.Project == "" {
		return fmt.Errorf("project missing: %w", ErrInvalidAppConfig)
	}
	if !c.Component.Valid() {
		return fmt.Errorf("invalid component: %s given: %w", c.Component.String(), ErrInvalidAppConfig)
	}
	if c.AppName == "" {
		return fmt.Errorf("app name missing: %w", ErrInvalidAppConfig)
	}
	if len(c.AppName) <= len(c.Project) || c.AppName[0:len(c.Project)+1] != fmt.Sprintf("%s-", c.Project) {
		return pkgerrors.WithStack(fmt.Errorf("app name should start with project name as prefix: %w", ErrInvalidAppConfig))
	}

	return nil
}

// NewConfigFromEnv returns the Config from sensible os envvars.
// NOTE: The Config may or may not be valid and should be validated separately if needed
func NewConfigFromEnv() Config {
	return Config{
		Env:       Env(os.Getenv("ENVIRONMENT")),
		Version:   os.Getenv("VERSION"),
		Server:    os.Getenv("SERVER_NAME"),
		Project:   os.Getenv("PROJECT_NAME"),
		Component: ComponentType(os.Getenv("PROJECT_COMPONENT")),
		Team:      os.Getenv("TEAM_NAME"),
		AppName:   os.Getenv("APP_NAME"),
	}
}
