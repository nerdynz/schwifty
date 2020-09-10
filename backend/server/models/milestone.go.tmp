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

var milestoneHelperGlobal *milestoneHelper

// Milestone Record
type Milestone struct {
  MilestoneULID string `db:"milestone_ulid" json:"MilestoneULID"`
  JobULID string `db:"job_ulid" json:"JobULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  Title string `db:"title" json:"Title"`
  Description string `db:"description" json:"Description"`
  Hours int `db:"hours" json:"Hours"`
  Depth int `db:"depth" json:"Depth"`
  
	TimeEntries  TimeEntries  `json:"TimeEntries"`
	
}

type Milestones []*Milestone

func (h *milestoneHelper) beforeSave(record *Milestone) (err error) {
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

func (h *milestoneHelper) afterSave(record *Milestone) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type milestoneHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func MilestoneHelper() *milestoneHelper {
	if milestoneHelperGlobal == nil {
		milestoneHelperGlobal = newMilestoneHelper(modelDB, modelCache)
	}
	return milestoneHelperGlobal
}

func newMilestoneHelper(db *runner.DB, cache Cache) *milestoneHelper {
	helper := &milestoneHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"milestone_ulid", "milestone_ulid", "job_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "title", "description", "hours", "depth"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *milestoneHelper) New(siteULID string) *Milestone {
	record := &Milestone{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *milestoneHelper) FromRequest(siteULID string, req *http.Request) (*Milestone, error) {
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
		return nil, errors.New("*Milestone update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *milestoneHelper) Load(siteULID string, ulid string) (*Milestone, error) {
	record, err := h.One(siteULID, "milestone_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *milestoneHelper) All(siteULID string) (Milestones, error) {
	var records Milestones
	err := h.DB.Select("*").
		From("milestone").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *milestoneHelper) Where(siteULID string, sql string, args ...interface{}) (Milestones, error) {
	var records Milestones
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("milestone").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *milestoneHelper) SQL(siteULID string, sql string, args ...interface{}) (Milestones, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Milestones
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *milestoneHelper) One(siteULID string, sql string, args ...interface{}) (*Milestone, error) {
	var record Milestone
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("milestone").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *milestoneHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *milestoneHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Milestones
	err := h.DB.Select("*").
		From("milestone").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(milestone_ulid) from milestone where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *milestoneHelper) Save(siteULID string, record *Milestone) error {
	return h.save(siteULID, record)
}

func (h *milestoneHelper) SaveMany(siteULID string, records Milestones) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *milestoneHelper) save(siteULID string, record *Milestone) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Milestone update failed. SiteID Mismatch")
	}
	cols := []string{ "milestone_ulid", "job_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "title", "description", "hours", "depth" }
	vals := []interface{}{ record.MilestoneULID, record.JobULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.Title, record.Description, record.Hours, record.Depth }
	err = h.DB.Upsert("milestone").
		Columns(cols...).
		Values(vals...).
		Where("milestone_ulid = $1", record.MilestoneULID).
		Returning("milestone_ulid").
		QueryStruct(record)

	//	if record.MilestoneULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("milestone")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("milestone_ulid = $1", record.MilestoneULID)
	//		b.Returning("milestone_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("milestone").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("milestone_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *milestoneHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("milestone").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and milestone_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *milestoneHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("milestone").
		Where("site_ulid=$1 and milestone_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *milestoneHelper) validate(record *Milestone) (err error) {
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



func (milestone *Milestone) SaveTimeEntries(siteULID string) error {
	return TimeEntryHelper().SaveMany(siteULID, milestone.TimeEntries)
}

func (milestone *Milestone) LoadTimeEntries(siteULID string) error {
	return milestone.LoadTimeEntriesWhere(siteULID, "milestone_ulid = $1 $SITEULID", milestone.MilestoneULID)
}

func (milestone *Milestone) LoadTimeEntriesWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := TimeEntryHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	milestone.TimeEntries = children
	return nil
}
