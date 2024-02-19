package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HealthAPI manages the health of the service
type HealthAPI struct {
	logger *logrus.Entry
}

func NewHealthAPI(logger *logrus.Entry) *HealthAPI {
	return &HealthAPI{
		logger: logger,
	}
}

func (h *HealthAPI) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func (h *HealthAPI) Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("OK"))
}
