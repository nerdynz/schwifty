package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// NewMilestone Route
func NewMilestone(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	milestone := milestoneHelper.New(siteULID)
	ctx.JSON(http.StatusOK, milestone)
}

// CreateMilestone Route
func CreateMilestone(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	milestone, err := milestoneHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create milestone record", err)
		return
	}
	err = milestoneHelper.Save(siteULID, milestone)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create milestone record", err)
		return
	}
	ctx.JSON(http.StatusOK, milestone)
}

// RetrieveMilestone Route
func RetrieveMilestone(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("milestoneULID") == "" {
		RetrieveMilestones(w, req, ctx, store)
		return
	}

	//get the milestoneULID from the request
	milestoneULID := ctx.URLParam("milestoneULID")
	if milestoneULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid milestoneULID", nil)
		return
	}

	milestoneHelper := models.MilestoneHelper()
	milestone, err := milestoneHelper.Load(siteULID, milestoneULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Milestone record", err)
		return
	}

	ctx.JSON(http.StatusOK, milestone)
}

// RetrieveMilestones Route
func RetrieveMilestones(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	milestones, err := milestoneHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Milestone records", err)
		return
	}

	ctx.JSON(http.StatusOK, milestones)
}

// PagedMilestones Route
func PagedMilestones(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := milestoneHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Milestone data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateMilestone Route
func UpdateMilestone(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	milestone, err := milestoneHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Milestone record for update", err)
		return
	}

	// save and validate
	err = milestoneHelper.Save(siteULID, milestone)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Milestone record", err)
		return
	}

	ctx.JSON(http.StatusOK, milestone)
}

// DeleteMilestone Route
func DeleteMilestone(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	milestoneHelper := models.MilestoneHelper()
	//get the milestoneULID from the request
	milestoneULID := ctx.URLParam("milestoneULID")
	if milestoneULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid MilestoneID for remove")
		return
	}

	isDeleted, err := milestoneHelper.Delete(siteULID, milestoneULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Milestone record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
