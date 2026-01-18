package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ci-supplychain-playground/app/internal/version"
)

func TestHealtz(t *testing.T) {
	r := NewRouter(RouterConfig{Version: version.Info()})
	
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr,req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	if ok, _ := body["ok"].(bool); !ok {
		t.Errorf("expected ok to be true, got %v", body["ok"])
	}
}

func TestEcho(t *testing.T) {
	r := NewRouter(RouterConfig{Version: version.Info()})
	req := httptest.NewRequest(http.MethodGet, "/echo?msg=hello", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var body map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if body["echo"] != "hello" {
		t.Fatalf("expected echo=hello, got %v", body["echo"])
	}
}