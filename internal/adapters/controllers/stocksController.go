package stock

import (
	"github.com/go-chi/chi/v5"
	"github.com/victorskg/stocks-api/internal/adapters/controllers/handlers"
)

type StocksController struct {
	saveStocksHandler handlers.SaveStocksHandler
	getStockHandler   handlers.GetStockHandler
}

func NewStocksController(saveStocksHandler handlers.SaveStocksHandler, getStockHandler handlers.GetStockHandler) StocksController {
	return StocksController{
		saveStocksHandler: saveStocksHandler,
		getStockHandler:   getStockHandler,
	}
}

func (c StocksController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", c.saveStocksHandler.SaveStocks)
	r.Route("/{ticker}", func(r chi.Router) {
		r.Get("/", c.getStockHandler.GetStock)
	})

	return r
}
