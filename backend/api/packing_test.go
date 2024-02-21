package api

import (
	"bytes"
	"encoding/json"
	"github.com/enkhalifapro/order-packs-calculator/internal/packing"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculate(t *testing.T) {
	router := httprouter.New()

	logger := &logrus.Entry{}

	// NOTE: since packing manager does not depend on concrete types, so no need to mock it. Otherwise it has to be mocked
	packingManager := packing.NewManager(logger)
	packingAPI := NewPackingAPI(logger, packingManager)
	router.POST("/calc", packingAPI.Calculate)

	body := make(map[string]interface{})
	body["items"] = 263
	body["packs"] = []int{31, 53, 23}
	b, _ := json.Marshal(body)
	r, _ := http.NewRequest(http.MethodPost, "/calc", bytes.NewReader(b))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if !(w.Code == http.StatusOK) {
		t.Errorf("routing failed.")
		t.FailNow()
	}
	resBody := w.Body.Bytes()
	res := packing.PackMix{}
	if err := json.Unmarshal(resBody, &packing.PackMix{}); err != nil {
		t.Errorf("unexpectd result failed.")
		t.FailNow()
	}
	assert.Equal(t, 0, res.ExtraItems)
}
