package entrypoint

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
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
		Message    string `json:"message"`
		Env        string `json:"env"`
		DBHost     string `json:"db_host"`
		DBUser     string `json:"db_user"`
		DBPassword string `json:"db_password"`
		DBPort     string `json:"db_port"`
	}{
		Message:    "Hello",
		Env:        os.Getenv("MY_CUSTOM_ENV"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(msg)
}
