package envalid_test

import (
	"testing"

	. "github.com/rhuancaetano/enceladus/pkg/envalid"
)

type OSMock struct {
	envs map[string]string
}

func (os *OSMock) LookupEnv(key string) (string, bool) {
	if value, hasEnv := os.envs[key]; hasEnv {
		return value, true
	}
	return "", false
}

func TestGetEnvironmentVariablesIntegerEnvType(t *testing.T) {
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "invalid input for GetEnvironmentVariables\n" {
			t.Errorf("should panic if envType is not a pointer to a struct")
		}
	}()

	GetEnvironmentVariables(1, os.LookupEnv)
}

func TestGetEnvironmentVariablesStringEnvType(t *testing.T) {
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "invalid input for GetEnvironmentVariables\n" {
			t.Errorf("should panic if envType is not a pointer to a struct")
		}
	}()

	GetEnvironmentVariables("test", os.LookupEnv)
}

func TestGetEnvironmentVariablesStructEnvType(t *testing.T) {
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "invalid input for GetEnvironmentVariables\n" {
			t.Errorf("should panic if envType is not a pointer to a struct")
		}
	}()

	type S struct{}

	GetEnvironmentVariables(S{}, os.LookupEnv)
}

func TestGetEnvironmentVariablesDefaultTag(t *testing.T) {
	type Environment struct {
		VAR string `env:"default"`
	}
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "'default' tag should have a value\n" {
			t.Errorf("should panic if default tag does not have a value")
		}
	}()

	GetEnvironmentVariables(&Environment{}, os.LookupEnv)
}

func TestGetEnvironmentVariablesRequiredTag(t *testing.T) {
	type Environment struct {
		VAR string `env:"required"`
	}
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "environment property VAR is required\n" {
			t.Errorf("should panic if it has required tag but no env with the same name")
		}
	}()

	GetEnvironmentVariables(&Environment{}, os.LookupEnv)
}

func TestGetEnvironmentVariablesUnadressableAttribute(t *testing.T) {
	type Temp struct{}
	type Environment struct {
		VAR *Temp
	}
	os := OSMock{
		envs: map[string]string{},
	}

	defer func() {
		if err, ok := recover().(string); !ok || err != "type '*envalid_test.Temp' is not supported by this package\n" {
			t.Errorf("struct properties must have a supported type")
		}
	}()

	GetEnvironmentVariables(&Environment{}, os.LookupEnv)
}
