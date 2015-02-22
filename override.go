package methodman // gopkg.in/httgo/methodman.v1

import (
	"net/http"
)

func MethodOverride(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			_meth := req.FormValue("_method")
			switch _meth {
			case "PUT", "PATCH", "DELETE":
				req.Method = _meth
			}
		}

		h.ServeHTTP(w, req)
	})
}
