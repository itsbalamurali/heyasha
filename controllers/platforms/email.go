package platforms

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func EmailBot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewDecoder(r.Body).Decode(update)
}
