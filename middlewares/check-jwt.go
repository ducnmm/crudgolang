package middlewares

import (
	"errors"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/mauduckg/crudgolang/auth"
	"github.com/mauduckg/crudgolang/utils"
)

func CheckJwt(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := jwt.Verify(r)

		if err != nil {
			res.ERROR(w, 401, errors.New("Unauthorized"))
			return
		}

		next(w, r, ps)
	}
}



