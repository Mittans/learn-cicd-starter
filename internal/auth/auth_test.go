package auth

import (
	"net/http"
	"testing"
)

func Test_GetAPIKey_Invalid(t *testing.T) {
	msg, err := GetAPIKey(nil)
	if msg != "" || err == nil {
		t.Errorf(`GetAPIKey("") = %q, %v, want "", error`, msg, err)
	}
}

func Test_GetAPIKey_Valid(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey Bearer test-token")
	msg, err := GetAPIKey(header)
	if err == nil {
		t.Errorf(`GetAPIKEY("") = %q, %v, want "" error`, msg, err)
	}
}
