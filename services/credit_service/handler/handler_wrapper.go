package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kumin/go-tpc/services/customer_service/entities"
)

type HandlerFn func(r *http.Request) (interface{}, error)

// handle http error
func HandlerWrapper(hfn HandlerFn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := hfn(r)
		if err != nil {
			switch err.Error() {
			case entities.ParamInvalid.Error():
				JSONError(w, err, http.StatusBadRequest)
			case entities.MethodNotAllowErr.Error():
				JSONError(w, err, http.StatusMethodNotAllowed)
			default:
				JSONError(w, err, http.StatusInternalServerError)
			}
			return
		}
		d, err := json.Marshal(obj)
		if err != nil {
			JSONError(w, err, http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		_, _ = w.Write(d)
	}
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}
