package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorskg/my-wallet/pkg/http/response"
	"github.com/victorskg/stocks-api/internal/usecases"
	"net/http"
)

type GetStockHandler struct {
	getStock usecases.GetStockUseCase
}

func NewGetStockHandler(getStock usecases.GetStockUseCase) GetStockHandler {
	return GetStockHandler{
		getStock: getStock,
	}
}

func (h GetStockHandler) GetStock(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "ticker")
	if ticker == "" {
		response.WriteResponseMessage(w, "O ticker do ativo é obrigatório.",
			http.StatusBadRequest)
		return
	}

	stock, err := h.getStock.Execute(ticker)
	if err != nil {
		responseStatus := http.StatusBadRequest

		// TODO validate not found error
		//_, isNotFoundErr := err.(domainErrors.NotFound)
		//
		//if isNotFoundErr {
		//	responseStatus = http.StatusNotFound
		//}

		response.WriteResponseMessage(w, err.Error(), responseStatus)
		return
	}

	// TODO Maybe do not return domain, instead, return a DTO
	response.WriteJSONResponse(w, stock, http.StatusOK)
}
