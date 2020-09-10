package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/backend/server/models"
)

// NewContact Route
func NewContact(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	contact := contactHelper.New(siteULID)
	ctx.JSON(http.StatusOK, contact)
}

// CreateContact Route
func CreateContact(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	contact, err := contactHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create contact record", err)
		return
	}
	err = contactHelper.Save(siteULID, contact)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create contact record", err)
		return
	}
	ctx.JSON(http.StatusOK, contact)
}

// RetrieveContact Route
func RetrieveContact(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	if ctx.URLParam("contactULID") == "" {
		RetrieveContacts(w, req, ctx, store)
		return
	}

	//get the contactULID from the request
	contactULID := ctx.URLParam("contactULID")
	if contactULID == "" {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid contactULID", nil)
		return
	}

	contactHelper := models.ContactHelper()
	contact, err := contactHelper.Load(siteULID, contactULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Contact record", err)
		return
	}

	ctx.JSON(http.StatusOK, contact)
}

// RetrieveContacts Route
func RetrieveContacts(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	contacts, err := contactHelper.All(siteULID)

	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve Contact records", err)
		return
	}

	ctx.JSON(http.StatusOK, contacts)
}

// PagedContacts Route
func PagedContacts(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	pageNum := ctx.URLIntParamWithDefault("pagenum", 1)
	limit := ctx.URLIntParamWithDefault("limit", 10)
	sort := ctx.URLParam("sort")
	direction := ctx.URLParam("direction")

	data, err := contactHelper.PagedBy(siteULID, pageNum, limit, sort, direction)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Unabled to get paged Contact data", err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// UpdateContact Route
func UpdateContact(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	contact, err := contactHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Contact record for update", err)
		return
	}

	// save and validate
	err = contactHelper.Save(siteULID, contact)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Contact record", err)
		return
	}

	ctx.JSON(http.StatusOK, contact)
}

// DeleteContact Route
func DeleteContact(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	contactHelper := models.ContactHelper()
	//get the contactULID from the request
	contactULID := ctx.URLParam("contactULID")
	if contactULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid ContactID for remove")
		return
	}

	isDeleted, err := contactHelper.Delete(siteULID, contactULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Contact record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
