package actions

import (
	"database/sql"
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/backend/server/models"
)

// NewSetting Route
func NewSetting(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	settingHelper := models.SettingHelper()
	setting := settingHelper.New(siteULID)
	ctx.JSON(http.StatusOK, setting)
}

// CreateSetting Route
func CreateSetting(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	settingHelper := models.SettingHelper()
	setting, err := settingHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create setting record", err)
		return
	}
	err = settingHelper.Save(siteULID, setting)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to create setting record", err)
		return
	}
	ctx.JSON(http.StatusOK, setting)
}

// RetrieveSettings Route
func RetrieveSettings(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	settingHelper := models.SettingHelper()
	setting, err := settingHelper.One(siteULID, "1=1")
	if err != nil {
		if err == sql.ErrNoRows && setting == nil {
			setting = settingHelper.New(siteULID)
		} else {
			ctx.ErrorJSON(http.StatusInternalServerError, "Failed to retrieve settings", err)
			return
		}
	}
	ctx.JSON(http.StatusOK, setting)
}

// UpdateSetting Route
func UpdateSetting(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	settingHelper := models.SettingHelper()
	setting, err := settingHelper.FromRequest(siteULID, ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load Setting record for update", err)
		return
	}

	// save and validate
	err = settingHelper.Save(siteULID, setting)
	// other type of error
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to save updated Setting record", err)
		return
	}

	ctx.JSON(http.StatusOK, setting)
}

// DeleteSetting Route
func DeleteSetting(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	siteULID, err := ctx.SiteULID()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Invalid Site", err)
		return
	}
	settingHelper := models.SettingHelper()
	//get the settingULID from the request
	settingULID := ctx.URLParam("settingULID")
	if settingULID == "" {
		ctx.JSON(http.StatusInternalServerError, "Invalid SettingID for remove")
		return
	}

	isDeleted, err := settingHelper.Delete(siteULID, settingULID)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to remove the Setting record", err)
		return
	}
	ctx.JSON(http.StatusOK, isDeleted)
}
