package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/backend/server/models"
)

// NewJob Route
func NewJob(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	job := jobHelper.New(siteULID)
	ctx.JSON(http.StatusOK, job)
}

// CreateJob Route
func CreateJob(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	job, err := jobHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create job record", err)
		return
	}
	err = jobHelper.Save(siteULID, job)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create job record", err)
		return
	}
	ctx.JSON(http.StatusOK, job)
}

// RetrieveJob Route
func RetrieveJob(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}

	if ctx.URLParam("jobULID") == "" {
		RetrieveJobs(w, req, ctx, store)
		return
	}

	//get the jobULID from the request
	jobULID, err := ctx.URLULIDParam("jobULID")
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid jobULID", err)
		return
	}

	jobHelper := models.JobHelper()
	job, err := jobHelper.Load(siteULID, jobULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Job record", err)
		return
	}

	if ctx.URLBoolParam("milestones") {
		err := job.LoadMilestones(siteULID)
		if err != nil {
			ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve milestones for job", err)
			return
		}
	}
	ctx.JSON(http.StatusOK, job)
}

// RetrieveJobs Route
func RetrieveJobs(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	jobs, err := jobHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Job records", err)
		return
	}

	ctx.JSON(http.StatusOK, jobs)
}

// PagedJobs Route
func PagedJobs(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := jobHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Job data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateJob Route
func UpdateJob(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	job, err := jobHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Job record for update", err)
		return
	}

	// save and validate
	err = jobHelper.Save(siteULID, job)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Job record", err)
		return
	}

	if ctx.URLBoolParam("milestones") {
		err := job.SaveMilestones(siteULID)
		if err != nil {
			ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save job milestones", err)
			return
		}
	}
	ctx.JSON(http.StatusOK, job)
}

// DeleteJob Route
func DeleteJob(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	jobHelper := models.JobHelper()
	//get the jobULID from the request
	jobULID, err := ctx.URLULIDParam("jobULID")
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid JobID for remove", err)
		return
	}

	isDeleted, err := jobHelper.Purge(siteULID, jobULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Job record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
