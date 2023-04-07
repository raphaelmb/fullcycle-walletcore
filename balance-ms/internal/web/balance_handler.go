package web

import (
	"net/http"

	getbalance "github.com/raphaelmb/fullcycle-balance-ms/internal/usecase/get_balance"
)

type WebBalanceHandler struct {
	GetBalanceByIDUseCase getbalance.GetBalanceByIDUseCase
}

func NewWebBalanceHandler(getBalanceByID getbalance.GetBalanceByIDUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceByIDUseCase: getBalanceByID,
	}
}

func (h *WebBalanceHandler) BalanceById(w http.ResponseWriter, r *http.Request) {
	// dto := getbalance.GetBalanceInputDTO{ID: "1"}
	// output, err := h.GetBalanceByIDUseCase.Execute(dto)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}
