package config

import (
	"testing"
)

// Should use default value when environment variable does not exist.
func Test_getEnv_default(t *testing.T) {

	res := getEnv("ENV_VARIABLE_NOT_DEFINED", "default-value")

	if res != "default-value" {
		t.Errorf("Failed testing getEnv()  Expected: %s  Got: %s", "default-value", res)
	}
}

// Should use default value when environment variable does not exist.
func Test_getEnvAsInt_default(t *testing.T) {

	res := getEnvAsInt("ENV_VARIABLE_NOT_DEFINED", 99)

	if res != 99 {
		t.Errorf("Failed testing getEnvAsIInt() Expected: %d  Got: %d", 99, res)
	}
}
