package actions

import (
	flow "github.com/nerdynz/flow"
)

func Estimate(ctx *flow.Context) {
	// siteID := ctx.SiteID()
	// id := ctx.URLIntParamWithDefault("projectID", -1)
	// job, err := models.JobHelper().Load(siteID, id)
	// if err != nil {
	// 	ctx.ErrorJSON(500, "Failed to create estimate", err)
	// 	return
	// }

	// err = job.LoadMilestones(siteID)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// client, err := models.ClientHelper().Load(siteID, job.ClientID)
	// if err != nil {
	// 	ctx.ErrorJSON(http.StatusInternalServerError, "No Client", err)
	// 	return
	// }
	// ctx.Add("Job", job)
	// ctx.Add("Client", client)

	// req := ctx.Req
	// // Create new PDF generator

	// // Add one page from an URL
	// proto := "http://"
	// if req.TLS != nil {
	// 	proto = "https://"
	// }

	// token := ctx.URLParam("authtoken")

	// u := proto + req.Host + "/quote/" + strconv.Itoa(id) + "?authtoken=" + token

	// if ctx.URLBoolParam("raw") {
	// 	ctx.Text(200, u)
	// 	return
	// }
	// b, err := pdf.Gen(u, "", false)
	// if err != nil {
	// 	ctx.ErrorHTML(500, "Failed to create quote", err)
	// 	return
	// }
	// ctx.PDF(b.Bytes())
}
