package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// NewBoard Route
func NewBoard(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	board := boardHelper.New(siteULID)
	ctx.JSON(http.StatusOK, board)
}

// CreateBoard Route
func CreateBoard(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	board, err := boardHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create board record", err)
		return
	}
	err = boardHelper.Save(siteULID, board)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create board record", err)
		return
	}
	ctx.JSON(http.StatusOK, board)
}

// RetrieveBoard Route
func RetrieveBoard(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}

	if ctx.URLParam("boardULID") == "" {
		RetrieveBoards(w, req, ctx, store)
		return
	}

	//get the boardULID from the request
	boardULID, err := ctx.URLULIDParam("boardULID")
	if boardULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid boardULID", nil)
		return
	}

	boardHelper := models.BoardHelper()
	board, err := boardHelper.Load(siteULID, boardULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Board record", err)
		return
	}

	if ctx.URLBoolParam("tasks") {
		err := board.LoadTasks(siteULID)
		if err != nil {
			ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve tasks for board", err)
			return
		}
	}

	ctx.JSON(http.StatusOK, board)
}

// RetrieveBoards Route
func RetrieveBoards(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	boards, err := boardHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Board records", err)
		return
	}

	ctx.JSON(http.StatusOK, boards)
}

// PagedBoards Route
func PagedBoards(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := boardHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Board data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateBoard Route
func UpdateBoard(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	board, err := boardHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Board record for update", err)
		return
	}

	// save and validate
	err = boardHelper.Save(siteULID, board)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Board record", err)
		return
	}

	ctx.JSON(http.StatusOK, board)
}

// DeleteBoard Route
func DeleteBoard(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	boardHelper := models.BoardHelper()
	//get the boardULID from the request
	boardULID := ctx.URLParam("boardULID")
	if boardULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid BoardID for remove")
		return
	}

	isDeleted, err := boardHelper.Delete(siteULID, boardULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Board record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
