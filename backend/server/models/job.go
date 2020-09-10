package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"strings"
	"sort"
	runner "github.com/nerdynz/dat/sqlx-runner"
)

var jobHelperGlobal *jobHelper

// Job Record
type Job struct {
  JobULID string `db:"job_ulid" json:"JobULID"`
  ClientULID string `db:"client_ulid" json:"ClientULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  Name string `db:"name" json:"Name"`
  QuoteNotes string `db:"quote_notes" json:"QuoteNotes"`
  Notes string `db:"notes" json:"Notes"`
  Status string `db:"status" json:"Status"`
  
	Milestones  Milestones  `json:"Milestones"`
	
}

type Jobs []*Job

func (h *jobHelper) beforeSave(record *Job) (err error) {
	if record.DateCreated.IsZero() {
		record.DateCreated = time.Now()
	}
	record.DateModified = time.Now()
	

	validationErr := h.validate(record)
	if validationErr != nil {
		return validationErr
	}
	return err
}

func (h *jobHelper) afterSave(record *Job) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type jobHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func JobHelper() *jobHelper {
	if jobHelperGlobal == nil {
		jobHelperGlobal = newJobHelper(modelDB, modelCache)
	}
	return jobHelperGlobal
}

func newJobHelper(db *runner.DB, cache Cache) *jobHelper {
	helper := &jobHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"job_ulid", "job_ulid", "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "name", "quote_notes", "notes", "status"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *jobHelper) New(siteULID string) *Job {
	record := &Job{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *jobHelper) FromRequest(siteULID string, req *http.Request) (*Job, error) {
	record := h.New(siteULID)
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		// working with json
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(record)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Disabled - bring in h.structDecoder from gorilla")
		// // working with form values
		// err := req.ParseForm()
		// if err != nil {
		// 	return nil, err
		// }

		// err = h.structDecoder.Decode(record, req.PostForm)
		// if err != nil {
		// 	return nil, err
		// }
	}
	if record.SiteULID != siteULID {
		return nil, errors.New("*Job update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *jobHelper) Load(siteULID string, ulid string) (*Job, error) {
	record, err := h.One(siteULID, "job_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *jobHelper) All(siteULID string) (Jobs, error) {
	var records Jobs
	err := h.DB.Select("*").
		From("job").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *jobHelper) Where(siteULID string, sql string, args ...interface{}) (Jobs, error) {
	var records Jobs
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("job").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *jobHelper) SQL(siteULID string, sql string, args ...interface{}) (Jobs, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Jobs
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *jobHelper) One(siteULID string, sql string, args ...interface{}) (*Job, error) {
	var record Job
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("job").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *jobHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *jobHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
	if orderByFieldName == "" || orderByFieldName == "default" {
		// we only want the first field name
		orderByFieldName = strings.Split(h.orderBy, ",")[0]
		orderByFieldName = strings.Trim(orderByFieldName, " ")
	}
	i := sort.SearchStrings(h.fieldNames, orderByFieldName)
	// check the orderby exists within the fields as this could be an easy sql injection hole.
	if !(i < len(h.fieldNames) && h.fieldNames[i] == orderByFieldName) { // NOT
		return nil, errors.New("field name [" + orderByFieldName + "]  isn't a valid field name")
	}

	if !(direction == "asc" || direction == "desc" || direction == "") {
		return nil, errors.New("direction isn't valid")
	}

	var records Jobs
	err := h.DB.Select("*").
		From("job").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(job_ulid) from job where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *jobHelper) Save(siteULID string, record *Job) error {
	return h.save(siteULID, record)
}

func (h *jobHelper) SaveMany(siteULID string, records Jobs) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *jobHelper) save(siteULID string, record *Job) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Job update failed. SiteID Mismatch")
	}
	cols := []string{ "job_ulid", "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "name", "quote_notes", "notes", "status" }
	vals := []interface{}{ record.JobULID, record.ClientULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.Name, record.QuoteNotes, record.Notes, record.Status }
	err = h.DB.Upsert("job").
		Columns(cols...).
		Values(vals...).
		Where("job_ulid = $1", record.JobULID).
		Returning("job_ulid").
		QueryStruct(record)

	//	if record.JobULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("job")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("job_ulid = $1", record.JobULID)
	//		b.Returning("job_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("job").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("job_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *jobHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("job").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and job_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *jobHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("job").
		Where("site_ulid=$1 and job_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *jobHelper) validate(record *Job) (err error) {
	return nil
//	validationErrors := h.validator.Struct(record)
//	if validationErrors != nil {
//		errMessage := ""
//		for _, err := range err.(validator.ValidationErrors) {
//			errMessage += err.Kind().String() + " validation Error on field "+err.Field()
//		}
//		if errMessage != "" {
//			err = errors.New(errMessage)
//		}
//	}
//	return err
}



func (job *Job) SaveMilestones(siteULID string) error {
	return MilestoneHelper().SaveMany(siteULID, job.Milestones)
}

func (job *Job) LoadMilestones(siteULID string) error {
	return job.LoadMilestonesWhere(siteULID, "job_ulid = $1 $SITEULID", job.JobULID)
}

func (job *Job) LoadMilestonesWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := MilestoneHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	job.Milestones = children
	return nil
}
