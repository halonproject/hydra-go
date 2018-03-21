package hydra

import "testing"

func TestNewError(t *testing.T) {
	err := newError(HYDRA_RESPONSE_TIMEOUT_ERROR)

	if str := err.String(); str != "" {
		t.Error("Expected err reason to be \"\" but got", str)
	}

	if code := err.Code(); code != HYDRA_RESPONSE_TIMEOUT_ERROR {
		t.Error("Expected err code to be 1 but got", code)
	}

	if errorStr := err.Error(); errorStr != "error code 1: " {
		t.Error("Expected error string to be \"error code 1: \" but got", errorStr)
	}
}
