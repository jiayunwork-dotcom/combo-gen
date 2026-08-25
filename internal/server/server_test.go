package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestPermutationsEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := permRequest{Items: []string{"a", "b", "c"}, K: 2}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/permutations", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 6 {
		t.Errorf("expected 6, got %d", resp.Count)
	}
}

func TestPermutationsEndpoint_Full(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := permRequest{Items: []string{"x", "y"}}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/permutations", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 2 {
		t.Errorf("expected 2, got %d", resp.Count)
	}
}

func TestPermutationsEndpoint_Empty(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	body := []byte(`{"items":[]}`)
	req := httptest.NewRequest(http.MethodPost, "/api/permutations", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestCombinationsEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := combRequest{Items: []string{"a", "b", "c", "d"}, K: 2}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/combinations", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 6 {
		t.Errorf("expected 6, got %d", resp.Count)
	}
}

func TestCombinationsEndpoint_WithRep(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := combRequest{Items: []string{"a", "b"}, K: 2, Repetition: true}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/combinations", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 3 {
		t.Errorf("expected 3, got %d", resp.Count)
	}
}

func TestProductEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := productRequest{Sets: [][]string{{"a", "b"}, {"1", "2"}}}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/product", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 4 {
		t.Errorf("expected 4, got %d", resp.Count)
	}
}

func TestProductEndpoint_Limit(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := productRequest{Sets: [][]string{{"a", "b", "c"}, {"1", "2", "3"}}, Limit: 3}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/product", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Count int `json:"count"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Count != 3 {
		t.Errorf("expected 3, got %d", resp.Count)
	}
}

func TestCountEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := countRequest{N: 5, K: 3}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/count", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp struct {
		Permutations int `json:"permutations"`
		Combinations int `json:"combinations"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Permutations != 60 {
		t.Errorf("expected perm=60, got %d", resp.Permutations)
	}
	if resp.Combinations != 10 {
		t.Errorf("expected comb=10, got %d", resp.Combinations)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	endpoints := []string{"/api/permutations", "/api/combinations", "/api/product", "/api/count"}
	for _, ep := range endpoints {
		req := httptest.NewRequest(http.MethodGet, ep, nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s: expected 405, got %d", ep, rec.Code)
		}
	}
}

func TestParsePort(t *testing.T) {
	if p := ParsePort(":8080"); p != 8080 {
		t.Errorf("expected 8080, got %d", p)
	}
}
