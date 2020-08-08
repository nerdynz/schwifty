package actions

import (
	"net/http"

	"github.com/nerdynz/datastore"
	flow "github.com/nerdynz/flow"
	"github.com/nerdynz/schwifty/server/server/models"
)

// Login Route
func Login(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	helper := models.PersonHelper()
	person, err := helper.FromRequest("", ctx.Req)
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to read login details", err)
		return
	}

	sessionInfo, err := ctx.Padlock.LoginReturningInfo(person.Email, person.Password)
	if err != nil {
		ctx.ErrorJSON(http.StatusUnauthorized, "Failed to login. Incorrect username or password", err)
		return
	}

	ctx.JSON(http.StatusOK, sessionInfo)
}

// Logout Route
func Logout(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	ctx.Padlock.Logout()
	ctx.Redirect("/", http.StatusSeeOther)
}
