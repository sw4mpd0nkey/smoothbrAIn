package httpserver

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/sw4mpd0nkey/smoothbrain/internal/httpx"
	logging "github.com/sw4mpd0nkey/smoothbrain/internal/logger"
	"github.com/sw4mpd0nkey/smoothbrain/internal/templates"
)

type middleware func(http.Handler) http.Handler

func applyMiddleware(handler http.Handler, middlewareChain ...middleware) http.Handler {
	for _, currentMiddleware := range slices.Backward(middlewareChain) {
		handler = currentMiddleware(handler)
	}
	return handler
}

func permissiveCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		if origin := request.Header.Get("Origin"); origin != "" {
			responseWriter.Header().Set("Access-Control-Allow-Origin", origin)
			responseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
			responseWriter.Header().Set("Vary", "Origin")
		}
		if request.Method == http.MethodOptions {
			responseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			responseWriter.Header().Set("Access-Control-Allow-Headers", request.Header.Get("Access-Control-Request-Headers"))
			responseWriter.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(responseWriter, request)
	})
}

func noSniff(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(responseWriter, request)
	})
}

func recoverPanics(logger *logging.Logger, renderer *templates.Renderer) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			defer func() {
				if recovered := recover(); recovered != nil {
					_ = logger.Event("unhandled_error", map[string]any{
						"method":  request.Method,
						"path":    request.URL.Path,
						"message": fmt.Sprint(recovered),
					})
					if err := httpx.RespondWithErrorPage(responseWriter, renderer, http.StatusInternalServerError, "Unhandled Error", fmt.Sprint(recovered)); err != nil {
						http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					}
				}
			}()
			next.ServeHTTP(responseWriter, request)
		})
	}
}

func LoadShedder(_ int, _ int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return next
	}
}
