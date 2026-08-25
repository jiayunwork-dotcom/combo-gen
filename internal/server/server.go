package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"combo-gen/internal/comb"
	"combo-gen/internal/count"
	"combo-gen/internal/perm"
	"combo-gen/internal/product"
)

type Config struct {
	Addr string
}

func New(cfg Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/api/permutations", handlePermutations)
	mux.HandleFunc("/api/combinations", handleCombinations)
	mux.HandleFunc("/api/product", handleProduct)
	mux.HandleFunc("/api/count", handleCount)
	return mux
}

func ListenAndServe(cfg Config) error {
	mux := New(cfg)
	return http.ListenAndServe(cfg.Addr, mux)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type permRequest struct {
	Items []string `json:"items"`
	K     int      `json:"k"`
}

func handlePermutations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req permRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if len(req.Items) == 0 {
		httpError(w, http.StatusBadRequest, "items array is empty")
		return
	}
	var (
		results [][]string
		err     error
	)
	if req.K > 0 {
		results, err = perm.PermutationsK(req.Items, req.K)
	} else {
		results = perm.Permutations(req.Items)
	}
	if err != nil {
		httpError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"count":   len(results),
		"results": results,
	})
}

type combRequest struct {
	Items      []string `json:"items"`
	K          int      `json:"k"`
	Repetition bool     `json:"repetition"`
}

func handleCombinations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req combRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if len(req.Items) == 0 {
		httpError(w, http.StatusBadRequest, "items array is empty")
		return
	}
	var (
		results [][]string
		err     error
	)
	if req.Repetition {
		results, err = comb.CombinationsWithRepetition(req.Items, req.K)
	} else {
		results, err = comb.Combinations(req.Items, req.K)
	}
	if err != nil {
		httpError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"count":   len(results),
		"results": results,
	})
}

type productRequest struct {
	Sets  [][]string `json:"sets"`
	Limit int        `json:"limit"`
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req productRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if len(req.Sets) == 0 {
		httpError(w, http.StatusBadRequest, "sets array is empty")
		return
	}
	var (
		results [][]string
		err     error
	)
	if req.Limit > 0 {
		results, err = product.CartesianProductLimit(req.Limit, req.Sets...)
	} else {
		results, err = product.CartesianProduct(req.Sets...)
	}
	if err != nil {
		httpError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"count":   len(results),
		"results": results,
	})
}

type countRequest struct {
	N int `json:"n"`
	K int `json:"k"`
}

func handleCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req countRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.N < 0 || req.K < 0 {
		httpError(w, http.StatusBadRequest, "n and k must be non-negative")
		return
	}
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"permutations":          count.PermutationCount(req.N, req.K),
		"combinations":          count.CombinationCount(req.N, req.K),
		"combinations_with_rep": count.CombinationWithRepCount(req.N, req.K),
		"factorial_n":           count.Factorial(req.N),
	})
}

func httpError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

func ParsePort(addr string) int {
	parts := strings.Split(addr, ":")
	if len(parts) < 2 {
		return 0
	}
	p, _ := strconv.Atoi(parts[len(parts)-1])
	return p
}

func FormatAddr(addr string) string {
	port := ParsePort(addr)
	if port == 0 {
		return addr
	}
	return fmt.Sprintf("http://0.0.0.0:%d", port)
}
