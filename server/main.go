package main

import (
	"net/http"

	"github.com/nerdynz/datastore"
	"github.com/nerdynz/rcache"

	"github.com/nerdynz/schwifty/server/server"
	"github.com/nerdynz/schwifty/server/server/models"

	"github.com/nerdynz/trove"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func main() {
	settings := trove.Load()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	log := logrus.New()
	cache := rcache.New(settings.Get("REDIS_URL"), log)
	store := datastore.New(log, settings, cache, nil)

	models.Init(store)
	attachments := negroni.NewStatic(http.Dir(store.Settings.Get("ATTACHMENTS_FOLDER")))
	attachments.Prefix = "/attachments"
	n.Use(attachments)
	adminNuxt := negroni.NewStatic(http.Dir("admin/dist/_nuxt"))
	adminNuxt.Prefix = "/admin/_nuxt"
	n.Use(adminNuxt)
	admin := negroni.NewStatic(http.Dir("./admin/dist/"))
	admin.Prefix = "/admin"
	n.Use(admin)
	public := negroni.NewStatic(http.Dir("./public/"))
	n.Use(public)

	r := server.Routes(store)
	handler := cors.AllowAll().Handler(r)
	n.UseHandler(handler)
	n.Run(":" + store.Settings.Get("PORT"))
}
