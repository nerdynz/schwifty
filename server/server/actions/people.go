package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// NewPerson Route
func NewPerson(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	person := personHelper.New(siteULID)
	ctx.JSON(http.StatusOK, person)
}

// CreatePerson Route
func CreatePerson(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	person, err := personHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create person record", err)
		return
	}
	err = personHelper.Save(siteULID, person)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create person record", err)
		return
	}
	ctx.JSON(http.StatusOK, person)
}

// RetrievePerson Route
func RetrievePerson(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("personULID") == "" {
		RetrievePeople(w, req, ctx, store)
		return
	}

	//get the personULID from the request
	personULID := ctx.URLParam("personULID")
	if personULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid personULID", nil)
		return
	}

	personHelper := models.PersonHelper()
	person, err := personHelper.Load(siteULID, personULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Person record", err)
		return
	}

	ctx.JSON(http.StatusOK, person)
}

// RetrievePeople Route
func RetrievePeople(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	people, err := personHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Person records", err)
		return
	}

	ctx.JSON(http.StatusOK, people)
}

// PagedPeople Route
func PagedPeople(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := personHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Person data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdatePerson Route
func UpdatePerson(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	person, err := personHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Person record for update", err)
		return
	}

	// save and validate
	err = personHelper.Save(siteULID, person)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Person record", err)
		return
	}

	ctx.JSON(http.StatusOK, person)
}

// DeletePerson Route
func DeletePerson(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	personHelper := models.PersonHelper()
	//get the personULID from the request
	personULID := ctx.URLParam("personULID")
	if personULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid PersonID for remove")
		return
	}

	isDeleted, err := personHelper.Delete(siteULID, personULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Person record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
