package httpapi


import(
	"net/http"
)

type RouterConfig struct {
	Version VersionInfo
}

func NewRouter(cfg RouterConfig) http.Handler {
	mux := http.NewServeMux()

	h := Handlers{Version: cfg.Version}

	mux.HandleFunc("GET /healthz", h.Healthz)
	mux.HandleFunc("GET /version", h.VersionHandler)
	mux.HandleFunc("GET /echo", h.Echo)

	return withStandardMiddleware(mux)
}

func withStandardMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}