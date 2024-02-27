package api

import (
	"encoding/json"
	"github.com/enkhalifapro/order-packs-calculator/internal/packing"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// PackingManager is an interface that holds the packing functionalities
type PackingManager interface {
	CalculatePacks(items int, packs []int) packing.PackMix
}

// PackCalcAPI manages the order packs calculations
type PackCalcAPI struct {
	manager PackingManager
	logger  *logrus.Entry
}

// NewPackingAPI returns a new instance of PackCalcAPI
func NewPackingAPI(logger *logrus.Entry, manager PackingManager) *PackCalcAPI {
	return &PackCalcAPI{
		manager: manager,
		logger:  logger,
	}
}

// Calculate returns the order packs calculations
func (p *PackCalcAPI) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error invalid payload", http.StatusBadRequest)
		p.logger.WithError(err)
		return
	}
	payload := struct {
		Items int   `json:"items"`
		Packs []int `json:"packs"`
	}{}
	if err := json.Unmarshal(reqBody, &payload); err != nil {
		http.Error(w, "Error try again ", http.StatusInternalServerError)
		p.logger.WithError(err)
		return
	}
	res := p.manager.CalculatePacks(payload.Items, payload.Packs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
