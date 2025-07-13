// HTTP server
package http

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

func InitHTTPAssets(serverMux *http.ServeMux, clientPort int, assets embed.FS) error {

	staticFS, err := fs.Sub(assets, "frontend/build")

	if err != nil {
		return err
	}

	serverMux.Handle("GET /", FileMiddleware(staticFS, http.FileServer(http.FS(staticFS))))

	httpServer := &http.Server{
		Handler: serverMux,
		Addr:    fmt.Sprintf("127.0.0.1:%d", clientPort),
	}

	err = httpServer.ListenAndServe()

	return err

}

// If .html of the route is available it loads the .html otherwise try to load the given path
func FileMiddleware(staticFS fs.FS, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pathSplited := strings.SplitN(r.URL.Path, "/", 2)

		if len(pathSplited) != 2 {
			next.ServeHTTP(w, r)
			return
		}

		data, err := fs.ReadFile(staticFS, pathSplited[1]+".html")

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(200)
		_, err = w.Write(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
