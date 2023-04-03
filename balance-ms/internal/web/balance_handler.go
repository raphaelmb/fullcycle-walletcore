package web

import (
	"net/http"
)

type WebBalanceHandler struct {
}

func NewWebBalanceHandler() *WebBalanceHandler {
	return &WebBalanceHandler{}
}

func (h *WebBalanceHandler) BalanceById(w http.ResponseWriter, r *http.Request) {
}
