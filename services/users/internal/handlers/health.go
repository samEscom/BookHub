package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Handlers struct {
	logger *zap.Logger
}

func NewHandlers(logger *zap.Logger) *Handlers {
	return &Handlers{logger: logger}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"status": "ok"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
