package handlers

import (
	"net/http"
	"time"

	"github.com/sw4mpd0nkey/smoothbrain/internal/httpx"
	logging "github.com/sw4mpd0nkey/smoothbrain/internal/logger"
	"github.com/sw4mpd0nkey/smoothbrain/internal/templates"
)

type pageView struct {
	templates.Page
	Error string
}
type WordHandler struct {
	renderer *templates.Renderer
	logger   *logging.Logger
}

func NewWordHandler(renderer *templates.Renderer, logger *logging.Logger) *WordHandler {
	return &WordHandler{renderer: renderer, logger: logger}
}

func (handler *WordHandler) Page(responseWriter http.ResponseWriter, request *http.Request) {
	if err := handler.renderPage(responseWriter, http.StatusOK, ""); err != nil {
		handler.internalError(responseWriter, request, err)
	}
}

func (handler *WordHandler) EncodeText(responseWriter http.ResponseWriter, request *http.Request) {
	//text := request.Form.
}

func (handler *WordHandler) renderPage(responseWriter http.ResponseWriter, statusCode int, errorMessage string) error {
	return handler.renderer.Render(responseWriter, statusCode, "word", pageView{
		Title: "Word Tokenizer Example",
		Error: errorMessage,
	})
}
func (handler *WordHandler) errorPage(responseWriter http.ResponseWriter, statusCode int, heading, message string) {
	if err := httpx.RespondWithErrorPage(responseWriter, handler.renderer, statusCode, heading, message); err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (handler *WordHandler) internalError(responseWriter http.ResponseWriter, request *http.Request, err error) {
	_ = handler.logger.Event("unhandled_error", map[string]any{"method": request.Method, "path": request.URL.Path, "message": err.Error()})
	handler.errorPage(responseWriter, http.StatusInternalServerError, "Unhandled Error", err.Error())
}

func formatTimestamp(timestamp time.Time) string {
	return timestamp.UTC().Format("2006-01-02T15:04:05.000Z")
}
