package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	"github.com/nerdynz/fileupload"
	flow "github.com/nerdynz/flow"
)

func FileUpload(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	filename, filePath, err := fileupload.FromRequestToFile(ctx.Req, ctx.Settings.Get("ATTACHMENTS_FOLDER"))
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Upload failed", err)
		return
	}
	data := &fileupload.File{}
	data.FileName = filename
	data.URL = ctx.Protocol + ctx.Req.Host + "/attachments/" + filename
	if ctx.URLParam("type") == "file" {
		data.URL = filePath
	}
	ctx.JSON(http.StatusOK, data)
}
