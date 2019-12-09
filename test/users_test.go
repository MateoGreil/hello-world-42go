package testing

import (
	"testing"
	"net/http/httptest"
)

funct TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
}