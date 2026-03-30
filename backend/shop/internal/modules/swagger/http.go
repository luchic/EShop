package swagger

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed openapi.json
var openAPISpec []byte

func Routers(mux *http.ServeMux) *http.ServeMux {
	if mux == nil {
		return mux
	}

	mux.HandleFunc("/swagger", handleSwaggerUI)
	mux.HandleFunc("/swagger/", handleSwaggerUI)
	mux.HandleFunc("/swagger/openapi.json", handleOpenAPISpec)
	return mux
}

func handleOpenAPISpec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(openAPISpec)
}

func handleSwaggerUI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Shop API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "/swagger/openapi.json",
        dom_id: "#swagger-ui"
      });
    };
  </script>
</body>
</html>`)
}
