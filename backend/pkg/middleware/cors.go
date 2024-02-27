package middleware

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CORS(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter,
		r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")

		next(w, r, ps)
	}
}
