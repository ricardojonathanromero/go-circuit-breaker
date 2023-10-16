package handler

import (
	"encoding/json"
	"github.com/ricardojonathanromero/go-circuit-breaker/client/internal/request"
	"net/http"
)

func NewHandler(httpRequest request.HttpRequest) *Handler {
	return &Handler{httpRequest}
}

type Handler struct {
	httpRequest request.HttpRequest
}

func (h *Handler) Ping(w http.ResponseWriter, _ *http.Request) {
	res, err := h.httpRequest.PingToServerApp()
	data := map[string]string{
		"message": "Message",
		"data":    string(res),
	}

	byteData, _ := json.Marshal(data)
	if err != nil {
		_, _ = w.Write(byteData)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(byteData)
}
