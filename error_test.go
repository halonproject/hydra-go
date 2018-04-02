package hydra

import "testing"

func TestNewError(t *testing.T) {
	err := newError(HYDRA_RESPONSE_TIMEOUT_ERROR, "timed out when reading message")

	if str := err.String(); str == "" {
		t.Error("Error string should not be empty")
	}

	if code := err.Code(); code != HYDRA_RESPONSE_TIMEOUT_ERROR {
		t.Errorf("Expected err code to be %d but got %d", HYDRA_RESPONSE_TIMEOUT_ERROR, code)
	}
}
