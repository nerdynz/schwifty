package actions

import (
	"io/ioutil"
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
)

// SPA Route
func SPA(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	ctx.W.Header().Add("content-type", "text/html")
	file, err := ioutil.ReadFile("admin/dist/index.html")
	if err != nil {
		ctx.ErrorHTML(500, "Failed to load SPA", err)
		return
	}
	ctx.Renderer.Data(ctx.W, 200, file)
}
