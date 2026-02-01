package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("failed to write response", "error", err)
	}

}
