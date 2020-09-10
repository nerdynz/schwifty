package server

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/flow"
	"github.com/nerdynz/router"
	"github.com/nerdynz/schwifty/backend/server/actions"
	"github.com/nerdynz/schwifty/backend/server/models"
	"github.com/nerdynz/security"
	"github.com/unrolled/render"
	"gopkg.in/mattes/migrate.v1/migrate"
)

var store *datastore.Datastore

func Routes(ds *datastore.Datastore) *bone.Mux {
	store = ds
	r := router.New(
		render.New(render.Options{
			Layout:     "application",
			Extensions: []string{".html"},
			Funcs: []template.FuncMap{
				HelperFuncs,
			},
			// prevent having to rebuild for every template reload... This is an important setting for development speed
			IsDevelopment:               store.Settings.IsDevelopment(),
			RequirePartials:             store.Settings.IsDevelopment(),
			RequireBlocks:               store.Settings.IsDevelopment(),
			RenderPartialsWithoutPrefix: true,
		}), ds, &Key{
			Store: store,
		},
	)
	// r.Mux.Handle("/admin/", http.FileServer(http.Dir("./admin/dist/")))
	if store.Settings.IsProduction() {
		r.GET("/admin/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
		r.GET("/admin/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a/:a", actions.SPA, security.NoAuth)
	}

	// r.GET("/__ua", actions.Analytics, security.NoAuth)

	// r.GET("/", actions.Home, security.NoAuth)
	// r.GET("/home", actions.RedirectHome, security.NoAuth)
	// // r.GET("/contact", actions.ContactUs, security.NoAuth)

	// r.GET("/kitchen-sink", actions.KitchenSink, security.NoAuth)

	// Scaffold routes
	// r.GET("/api/v1/views/:month/:year", actions.Views, security.NoAuth)
	// r.GET("/api/v1/sitesettings", siteSettings, security.NoAuth)
	r.GET("/api/v1/schema", Schema, security.NoAuth)
	r.PST("/api/v1/login", actions.Login, security.NoAuth)
	// r.GET("/api/v1/user", actions.UserDetails, security.Disallow)

	r.GET("/:api/v1/person/new", actions.NewPerson, security.Disallow)
	r.PST("/:api/v1/person/create", actions.CreatePerson, security.Disallow)
	r.GET("/:api/v1/person/retrieve", actions.RetrievePeople, security.Disallow)
	r.GET("/:api/v1/person/retrieve/:personULID", actions.RetrievePerson, security.Disallow)
	r.GET("/:api/v1/person/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedPeople, security.Disallow)
	r.PUT("/:api/v1/person/update/:personULID", actions.UpdatePerson, security.Disallow)
	r.DEL("/:api/v1/person/delete/:personULID", actions.DeletePerson, security.Disallow)

	r.GET("/:api/v1/message/new", actions.NewMessage, security.Disallow)
	r.PST("/:api/v1/message/create", actions.CreateMessage, security.Disallow)
	r.GET("/:api/v1/message/retrieve", actions.RetrieveMessages, security.Disallow)
	r.GET("/:api/v1/message/retrieve/:messageULID", actions.RetrieveMessage, security.Disallow)
	r.GET("/:api/v1/message/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedMessages, security.Disallow)
	r.PUT("/:api/v1/message/update/:messageULID", actions.UpdateMessage, security.Disallow)
	r.DEL("/:api/v1/message/delete/:messageULID", actions.DeleteMessage, security.Disallow)

	// r.GET("/api/v1/people/new", actions.NewPerson, security.NoAuth)
	// r.PST("/api/v1/people/create", actions.CreatePerson, security.Disallow)
	// r.GET("/api/v1/people/retrieve", actions.RetrievePeople, security.NoAuth)
	// r.GET("/api/v1/people/retrieve/:personID", actions.RetrievePerson, security.NoAuth)
	// r.PUT("/api/v1/people/update/:personID", actions.UpdatePerson, security.Disallow)

	// r.GET("/api/v1/page/new", actions.NewPage, security.NoAuth)
	// r.PST("/api/v1/page/create", actions.CreatePage, security.Disallow)
	// r.GET("/api/v1/page/retrieve", actions.RetrievePages, security.NoAuth)
	// r.GET("/api/v1/page/retrieve/:pageID", actions.RetrievePage, security.NoAuth)
	// r.GET("/api/v1/page/retrieve/byslug/:slug", actions.RetrievePageBySlug, security.NoAuth)
	// r.PUT("/api/v1/page/update/:pageID", actions.UpdatePage, security.Disallow)
	// r.DEL("/api/v1/page/delete/:pageID", actions.DeletePage, security.Disallow)
	// r.PUT("/api/v1/page/sort", actions.ChangePageSort, security.Disallow)

	// r.GET("/:api/v1/person/new", actions.NewPerson, security.Disallow)
	// r.PST("/:api/v1/person/create", actions.CreatePerson, security.Disallow)
	// r.GET("/:api/v1/person/retrieve", actions.RetrievePeople, security.Disallow)
	// r.GET("/:api/v1/person/retrieve/:personID", actions.RetrievePerson, security.Disallow)
	// r.GET("/:api/v1/person/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedPeople, security.Disallow)
	// r.PUT("/:api/v1/person/update/:personID", actions.UpdatePerson, security.Disallow)
	// r.DEL("/:api/v1/person/delete/:personID", actions.DeletePerson, security.Disallow)

	r.GET("/:api/v1/client/new", actions.NewClient, security.Disallow)
	r.PST("/:api/v1/client/create", actions.CreateClient, security.Disallow)
	r.GET("/:api/v1/client/retrieve", actions.RetrieveClients, security.Disallow)
	r.GET("/:api/v1/client/retrieve/:clientID", actions.RetrieveClient, security.Disallow)
	r.GET("/:api/v1/client/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedClients, security.Disallow)
	r.PUT("/:api/v1/client/update/:clientID", actions.UpdateClient, security.Disallow)
	r.DEL("/:api/v1/client/delete/:clientID", actions.DeleteClient, security.Disallow)

	r.GET("/:api/v1/job/new", actions.NewJob, security.Disallow)
	r.PST("/:api/v1/job/create", actions.CreateJob, security.Disallow)
	r.GET("/:api/v1/job/retrieve", actions.RetrieveJobs, security.Disallow)
	r.GET("/:api/v1/job/retrieve/:jobULID", actions.RetrieveJob, security.Disallow)
	r.GET("/:api/v1/job/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedJobs, security.Disallow)
	r.PUT("/:api/v1/job/update/:jobULID", actions.UpdateJob, security.Disallow)
	r.DEL("/:api/v1/job/delete/:jobULID", actions.DeleteJob, security.Disallow)

	r.GET("/:api/v1/board/new", actions.NewBoard, security.Disallow)
	r.PST("/:api/v1/board/create", actions.CreateBoard, security.Disallow)
	r.GET("/:api/v1/board/retrieve", actions.RetrieveBoards, security.Disallow)
	r.GET("/:api/v1/board/retrieve/:boardULID", actions.RetrieveBoard, security.Disallow)
	r.GET("/:api/v1/board/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedBoards, security.Disallow)
	r.PUT("/:api/v1/board/update/:boardULID", actions.UpdateBoard, security.Disallow)
	r.DEL("/:api/v1/board/delete/:boardULID", actions.DeleteBoard, security.Disallow)

	r.GET("/:api/v1/task/new", actions.NewTask, security.Disallow)
	r.PST("/:api/v1/task/create", actions.CreateTask, security.Disallow)
	r.GET("/:api/v1/task/retrieve", actions.RetrieveTasks, security.Disallow)
	r.GET("/:api/v1/task/retrieve/:taskULID", actions.RetrieveTask, security.Disallow)
	r.GET("/:api/v1/task/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedTasks, security.Disallow)
	r.PUT("/:api/v1/task/update/:taskULID", actions.UpdateTask, security.Disallow)
	r.DEL("/:api/v1/task/delete/:taskULID", actions.DeleteTask, security.Disallow)

	r.GET("/:api/v1/actionable/new", actions.NewActionable, security.Disallow)
	r.PST("/:api/v1/actionable/create", actions.CreateActionable, security.Disallow)
	r.GET("/:api/v1/actionable/retrieve", actions.RetrieveActionables, security.Disallow)
	r.GET("/:api/v1/actionable/retrieve/:actionableID", actions.RetrieveActionable, security.Disallow)
	r.GET("/:api/v1/actionable/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedActionables, security.Disallow)
	r.PUT("/:api/v1/actionable/update/:actionableID", actions.UpdateActionable, security.Disallow)
	r.DEL("/:api/v1/actionable/delete/:actionableID", actions.DeleteActionable, security.Disallow)

	r.GET("/:api/v1/setting/new", actions.NewSetting, security.Disallow)
	r.PST("/:api/v1/setting/create", actions.CreateSetting, security.Disallow)
	r.GET("/:api/v1/setting/retrieve", actions.RetrieveSettings, security.Disallow)
	// r.GET("/:api/v1/setting/retrieve/:settingID", actions.RetrieveSetting, security.Disallow)
	// r.GET("/:api/v1/setting/paged/:sort/:direction/limit/:limit/pagenum/:pagenum", actions.PagedSettings, security.Disallow)
	r.PUT("/:api/v1/setting/update/:settingID", actions.UpdateSetting, security.Disallow)
	r.DEL("/:api/v1/setting/delete/:settingID", actions.DeleteSetting, security.Disallow)

	// r.POST("/api/v1/upload/crop", actions.CroppedFileUpload, security.Disallow)
	// r.POST("/api/v1/upload/:quality/:type", actions.FileUpload, security.NoAuth)
	r.POST("/api/v1/upload/:type", actions.FileUpload, security.NoAuth)
	// r.GET("/:api/v1/imagemeta/retrieve/:uniqueid", actions.RetrieveImageMeta, security.Disallow)

	// r.GET("/sitemap.xml", websitemap, security.NoAuth)
	r.GET("/robots.txt", robots, security.NoAuth)

	r.GET("/migrate", doMigrate, security.NoAuth)

	// GOES LAST FOR GOOD MEASURE
	// r.GET("/:slug", actions.RenderPageBySlug, security.NoAuth)
	return r.Mux
}

func Schema(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	data := struct {
		Message    *models.Message
		Person     *models.Person
		Job        *models.Job
		Milestone  *models.Milestone
		Board      *models.Board
		Task       *models.Task
		Actionable *models.Actionable
		Setting    *models.Setting
		Client     *models.Client
	}{
		Message:    models.MessageHelper().New(""),
		Person:     models.PersonHelper().New(""),
		Job:        models.JobHelper().New(""),
		Milestone:  models.MilestoneHelper().New(""),
		Board:      models.BoardHelper().New(""),
		Task:       models.TaskHelper().New(""),
		Actionable: models.ActionableHelper().New(""),
		Setting:    models.SettingHelper().New(""),
		Client:     models.ClientHelper().New(""),
	}
	ctx.JSON(http.StatusOK, data)
}

func robots(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	robotsTxt := `User-agent: Teoma
Disallow: /
User-agent: twiceler
Disallow: /
User-agent: Gigabot
Disallow: /
User-agent: Scrubby
Disallow: /
User-agent: Nutch
Disallow: /
User-agent: baiduspider
Disallow: /
User-agent: naverbot
Disallow: /
User-agent: yeti
Disallow: /
User-agent: psbot
Disallow: /
User-agent: asterias
Disallow: /
User-agent: yahoo-blogs
Disallow: /
User-agent: YandexBot
Disallow: /
User-agent: Sosospider
Disallow: /
User-agent: *
Disallow: /admin
User-agent: *
Disallow: /df9249a6-0d56-11e8-ba89-0ed5f89f718b
User-agent: *
Disallow: /kitchen-sink
Sitemap: ` + ctx.Settings.Get("WEBSITE_BASE_URL") + `sitemap.xml`
	ctx.Renderer.Text(ctx.W, 200, robotsTxt)
}

// func websitemap(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
// 	sm := sitemap.New()
// 	pages, _ := models.PageHelper().All()

// 	for _, page := range pages {
// 		if page.ShowInNav == "Placeholder" {
// 			continue // skip placeholder
// 		}
// 		sm.Add(&sitemap.URL{
// 			Loc:        ctx.Settings.Get("WEBSITE_BASE_URL") + page.Slug + "/",
// 			LastMod:    &page.DateModified,
// 			ChangeFreq: sitemap.Weekly,
// 		})
// 	}

// 	ctx.W.Header().Set("Content-Type", "text/xml")
// 	sm.WriteTo(ctx.W)
// }

// func siteSettings(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
// 	topNav, err := models.PageHelper().LoadTopNav()
// 	if err != nil {
// 		ctx.ErrorJSON(http.StatusInternalServerError, "Failed to load top nav", err)
// 		return
// 	}
// 	data := struct {
// 		TopNav models.NavItems
// 	}{
// 		topNav,
// 	}
// 	ctx.JSON(http.StatusOK, data)
// }

func doMigrate(w http.ResponseWriter, req *http.Request, ctx *flow.Context, store *datastore.Datastore) {
	errs, ok := migrate.UpSync(ctx.Settings.Get("DATABASE_URL")+"?sslmode=disable", "./server/models/migrations")
	if !ok {
		finalError := ""
		for _, err := range errs {
			finalError += err.Error() + "\n"
		}
		ctx.ErrorText(500, "migration failed", errors.New(finalError))
	}
	ctx.Text(200, "Done")
}
