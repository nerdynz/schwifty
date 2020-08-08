package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// NewActionable Route
func NewActionable(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	actionable := actionableHelper.New(siteULID)
	ctx.JSON(http.StatusOK, actionable)
}

// CreateActionable Route
func CreateActionable(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	actionable, err := actionableHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create actionable record", err)
		return
	}
	err = actionableHelper.Save(siteULID, actionable)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create actionable record", err)
		return
	}
	ctx.JSON(http.StatusOK, actionable)
}

// RetrieveActionable Route
func RetrieveActionable(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("actionableULID") == "" {
		RetrieveActionables(w, req, ctx, store)
		return
	}

	//get the actionableULID from the request
	actionableULID := ctx.URLParam("actionableULID")
	if actionableULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid actionableULID", nil)
		return
	}

	actionableHelper := models.ActionableHelper()
	actionable, err := actionableHelper.Load(siteULID, actionableULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Actionable record", err)
		return
	}

	ctx.JSON(http.StatusOK, actionable)
}

// RetrieveActionables Route
func RetrieveActionables(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	actionables, err := actionableHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Actionable records", err)
		return
	}

	ctx.JSON(http.StatusOK, actionables)
}

// PagedActionables Route
func PagedActionables(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := actionableHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Actionable data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateActionable Route
func UpdateActionable(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	actionable, err := actionableHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Actionable record for update", err)
		return
	}

	// save and validate
	err = actionableHelper.Save(siteULID, actionable)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Actionable record", err)
		return
	}

	ctx.JSON(http.StatusOK, actionable)
}

// DeleteActionable Route
func DeleteActionable(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	actionableHelper := models.ActionableHelper()
	//get the actionableULID from the request
	actionableULID := ctx.URLParam("actionableULID")
	if actionableULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid ActionableID for remove")
		return
	}

	isDeleted, err := actionableHelper.Delete(siteULID, actionableULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Actionable record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
