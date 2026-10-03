package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingReturnsPlainTextPong(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	res := httptest.NewRecorder()

	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusOK)
	}
	if got := res.Header().Get("Content-Type"); got != "text/plain; charset=utf-8" {
		t.Fatalf("Content-Type = %q, want %q", got, "text/plain; charset=utf-8")
	}
	if got := res.Body.String(); got != "pong" {
		t.Fatalf("body = %q, want %q", got, "pong")
	}
}

func TestPingRejectsNonGetMethods(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/ping", nil)
	res := httptest.NewRecorder()

	NewHandler().ServeHTTP(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", res.Code, http.StatusMethodNotAllowed)
	}
}
