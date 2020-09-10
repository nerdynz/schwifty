package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/backend/server/models"
)

// NewTask Route
func NewTask(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	task := taskHelper.New(siteULID)
	ctx.JSON(http.StatusOK, task)
}

// CreateTask Route
func CreateTask(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	task, err := taskHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create task record", err)
		return
	}
	err = taskHelper.Save(siteULID, task)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create task record", err)
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// RetrieveTask Route
func RetrieveTask(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("taskULID") == "" {
		RetrieveTasks(w, req, ctx, store)
		return
	}

	//get the taskULID from the request
	taskULID, err := ctx.URLULIDParam("taskULID")
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid taskULID", err)
		return
	}

	taskHelper := models.TaskHelper()
	task, err := taskHelper.Load(siteULID, taskULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Task record", err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// RetrieveTasks Route
func RetrieveTasks(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	tasks, err := taskHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Task records", err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

// PagedTasks Route
func PagedTasks(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := taskHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Task data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateTask Route
func UpdateTask(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	task, err := taskHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Task record for update", err)
		return
	}

	// save and validate
	err = taskHelper.Save(siteULID, task)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Task record", err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// DeleteTask Route
func DeleteTask(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	taskHelper := models.TaskHelper()
	//get the taskULID from the request
	taskULID := ctx.URLParam("taskULID")
	if taskULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid TaskID for remove")
		return
	}

	isDeleted, err := taskHelper.Delete(siteULID, taskULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Task record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}

// r.GET("/:api/v1/task/new", actions.NewTask, security.Disallow)
// r.PST("/:api/v1/task/create", actions.CreateTask, security.Disallow)
// r.GET("/:api/v1/task/retrieve", actions.RetrieveTasks, security.Disallow)
// r.GET("/:api/v1/task/retrieve/:taskULID", actions.RetrieveTask, security.Disallow)
// r.GET("/:api/v1/task/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedTasks, security.Disallow)
// r.PUT("/:api/v1/task/update/:taskULID", actions.UpdateTask, security.Disallow)
// r.DEL("/:api/v1/task/delete/:taskULID", actions.DeleteTask, security.Disallow)
