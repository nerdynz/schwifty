package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// NewClient Route
func NewClient(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	client := clientHelper.New(siteULID)
	ctx.JSON(http.StatusOK, client)
}

// CreateClient Route
func CreateClient(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	client, err := clientHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create client record", err)
		return
	}
	err = clientHelper.Save(siteULID, client)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create client record", err)
		return
	}
	ctx.JSON(http.StatusOK, client)
}

// RetrieveClient Route
func RetrieveClient(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("clientULID") == "" {
		RetrieveClients(w, req, ctx, store)
		return
	}

	//get the clientULID from the request
	clientULID := ctx.URLParam("clientULID")
	if clientULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid clientULID", nil)
		return
	}

	clientHelper := models.ClientHelper()
	client, err := clientHelper.Load(siteULID, clientULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Client record", err)
		return
	}

	ctx.JSON(http.StatusOK, client)
}

// RetrieveClients Route
func RetrieveClients(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	clients, err := clientHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Client records", err)
		return
	}

	ctx.JSON(http.StatusOK, clients)
}

// PagedClients Route
func PagedClients(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := clientHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Client data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateClient Route
func UpdateClient(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	client, err := clientHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Client record for update", err)
		return
	}

	// save and validate
	err = clientHelper.Save(siteULID, client)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Client record", err)
		return
	}

	ctx.JSON(http.StatusOK, client)
}

// DeleteClient Route
func DeleteClient(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	clientHelper := models.ClientHelper()
	//get the clientULID from the request
	clientULID := ctx.URLParam("clientULID")
	if clientULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid ClientID for remove")
		return
	}

	isDeleted, err := clientHelper.Delete(siteULID, clientULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Client record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
