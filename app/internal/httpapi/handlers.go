package httpapi

import(
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type VersionInfo interface {
	String() string
	JSON() map[string]string
}

type Handlers struct {
	Version VersionInfo
}

func (h *Handlers) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"ok":		true,
		"ts":		time.Now().UTC().Format(time.RFC3339),
		"components": "ci-supplychain-playground",
	})
}

func (h Handlers) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.Version.JSON())
}

func (h Handlers) Echo(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	msg = strings.TrimSpace(msg)
	if len(msg) > 200 {
		msg = msg[:200]
	}
	msg = strings.ReplaceAll(msg, "\n", " ")
	msg = strings.ReplaceAll(msg, "\r", " ")

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"echo": msg,
	})
}