package httpserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sw4mpd0nkey/smoothbrain/internal/handlers"
	"github.com/sw4mpd0nkey/smoothbrain/internal/httpx"
	logging "github.com/sw4mpd0nkey/smoothbrain/internal/logger"
	"github.com/sw4mpd0nkey/smoothbrain/internal/templates"
)

type Application struct {
	Handler    http.Handler
	publicRoot *os.Root
}

type Options struct {
	AppOrigin               string
	MaxPublicProductResults int
	MaxRequestBodyBytes     int64
	MaxUploadBytes          int64
	PawPalAPIKey            string
	AcornFulfillmentDelay   time.Duration
	DataDirectory           string
	FixtureDirectory        string
	TemplateDirectory       string
	PublicDirectory         string
}

func New(database *sql.DB, logger *logging.Logger, options Options) (*Application, error) {
	if options.DataDirectory == "" {
		return nil, fmt.Errorf("data directory is required")
	}
	if options.FixtureDirectory == "" {
		options.FixtureDirectory = filepath.Join(options.DataDirectory, "fixtures")
	}
	if options.MaxUploadBytes <= 0 {
		return nil, fmt.Errorf("maximum upload size must be positive")
	}
	renderer, err := templates.Load(options.TemplateDirectory)
	if err != nil {
		return nil, err
	}
	publicRoot, err := os.OpenRoot(options.PublicDirectory)
	if err != nil {
		return nil, fmt.Errorf("open public directory: %w", err)
	}

	wordhandler := handlers.NewWordHandler(renderer, logger)

	dynamicMux := http.NewServeMux()
	dynamicMux.HandleFunc("GET /word-tokenizer", wordhandler.Page)
	dynamicHandler := permissiveCORS(dynamicMux)

	mainMux := http.NewServeMux()
	mainMux.HandleFunc("GET /health", func(responseWriter http.ResponseWriter, _ *http.Request) {
		httpx.RespondWithJSON(responseWriter, http.StatusOK, map[string]any{"ok": true, "app": "bearly-secure"})
	})
	staticHandler := newStaticHandler(publicRoot)
	mainMux.Handle("GET /reset.css", staticHandler)
	mainMux.Handle("GET /styles.css", staticHandler)
	mainMux.Handle("GET /passkey.js", staticHandler)
	mainMux.Handle("GET /vendor/simplewebauthn/index.umd.min.js", staticHandler)
	mainMux.Handle("GET /shipping-widget.css", staticHandler)
	mainMux.Handle("GET /shipping-widget.html", staticHandler)
	mainMux.Handle("GET /shipping-widget.js", staticHandler)
	mainMux.Handle("GET /product-photos/{filename}", staticHandler)
	mainMux.Handle("/", dynamicHandler)

	handler := applyMiddleware(
		mainMux,
		noSniff,
		recoverPanics(logger, renderer),
	)

	fmt.Println(renderer)
	return &Application{Handler: handler, publicRoot: publicRoot}, nil
}

func newStaticHandler(publicRoot *os.Root) http.Handler {
	fileServer := http.FileServerFS(publicRoot.FS())
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		relativePath := request.URL.Path[1:]
		fileInfo, err := publicRoot.Stat(relativePath)
		if err != nil || fileInfo.IsDir() {
			http.NotFound(responseWriter, request)
			return
		}
		fileServer.ServeHTTP(responseWriter, request)
	})
}

func (application *Application) Close() error {
	return application.publicRoot.Close()
}
