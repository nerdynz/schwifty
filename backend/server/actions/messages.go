package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/backend/server/models"
)

// NewMessage Route
func NewMessage(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	message := messageHelper.New(siteULID)
	ctx.JSON(http.StatusOK, message)
}

// CreateMessage Route
func CreateMessage(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	message, err := messageHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create message record", err)
		return
	}
	err = messageHelper.Save(siteULID, message)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create message record", err)
		return
	}
	ctx.JSON(http.StatusOK, message)
}

// RetrieveMessage Route
func RetrieveMessage(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("messageULID") == "" {
		RetrieveMessages(w, req, ctx, store)
		return
	}

	//get the messageULID from the request
	messageULID := ctx.URLParam("messageULID")
	if messageULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid messageULID", nil)
		return
	}

	messageHelper := models.MessageHelper()
	message, err := messageHelper.Load(siteULID, messageULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Message record", err)
		return
	}

	ctx.JSON(http.StatusOK, message)
}

// RetrieveMessages Route
func RetrieveMessages(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	messages, err := messageHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Message records", err)
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

// PagedMessages Route
func PagedMessages(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := messageHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Message data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateMessage Route
func UpdateMessage(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	message, err := messageHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Message record for update", err)
		return
	}

	// save and validate
	err = messageHelper.Save(siteULID, message)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Message record", err)
		return
	}

	ctx.JSON(http.StatusOK, message)
}

// DeleteMessage Route
func DeleteMessage(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	messageHelper := models.MessageHelper()
	//get the messageULID from the request
	messageULID := ctx.URLParam("messageULID")
	if messageULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid MessageID for remove")
		return
	}

	isDeleted, err := messageHelper.Delete(siteULID, messageULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Message record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
