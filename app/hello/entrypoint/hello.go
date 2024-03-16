package entrypoint

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

var logger = slog.Default().WithGroup("hello")

type (
	HelloHandler struct {
	}
)

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h HelloHandler) Register() {
	http.HandleFunc("/hello", h.handle)
}

func (h HelloHandler) handle(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request")
	msg, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: "Hello",
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(msg)
}
