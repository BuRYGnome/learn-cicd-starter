package auth

import (
	"net/http"
	"testing"
)

func TestSomething(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	got, err := GetAPIKey(headers)

	if err != nil {
		t.Fatal(err)
	}

	if got != "my-secret-key" {
		t.Errorf("expected %q, got %q", "my-secret-key", got)
	}
}
