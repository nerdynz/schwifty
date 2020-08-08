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

var actionableHelperGlobal *actionableHelper

// Actionable Record
type Actionable struct {
  ActionableULID string `db:"actionable_ulid" json:"ActionableULID"`
  TaskULID string `db:"task_ulid" json:"TaskULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  
	
}

type Actionables []*Actionable

func (h *actionableHelper) beforeSave(record *Actionable) (err error) {
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

func (h *actionableHelper) afterSave(record *Actionable) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type actionableHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func ActionableHelper() *actionableHelper {
	if actionableHelperGlobal == nil {
		actionableHelperGlobal = newActionableHelper(modelDB, modelCache)
	}
	return actionableHelperGlobal
}

func newActionableHelper(db *runner.DB, cache Cache) *actionableHelper {
	helper := &actionableHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"actionable_ulid", "actionable_ulid", "task_ulid", "site_ulid", "date_created", "date_modified", "date_deleted"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *actionableHelper) New(siteULID string) *Actionable {
	record := &Actionable{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *actionableHelper) FromRequest(siteULID string, req *http.Request) (*Actionable, error) {
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
		return nil, errors.New("*Actionable update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *actionableHelper) Load(siteULID string, ulid string) (*Actionable, error) {
	record, err := h.One(siteULID, "actionable_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *actionableHelper) All(siteULID string) (Actionables, error) {
	var records Actionables
	err := h.DB.Select("*").
		From("actionable").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *actionableHelper) Where(siteULID string, sql string, args ...interface{}) (Actionables, error) {
	var records Actionables
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("actionable").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *actionableHelper) SQL(siteULID string, sql string, args ...interface{}) (Actionables, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Actionables
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *actionableHelper) One(siteULID string, sql string, args ...interface{}) (*Actionable, error) {
	var record Actionable
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("actionable").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *actionableHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *actionableHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Actionables
	err := h.DB.Select("*").
		From("actionable").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(actionable_ulid) from actionable where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *actionableHelper) Save(siteULID string, record *Actionable) error {
	return h.save(siteULID, record)
}

func (h *actionableHelper) SaveMany(siteULID string, records Actionables) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *actionableHelper) save(siteULID string, record *Actionable) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Actionable update failed. SiteID Mismatch")
	}
	cols := []string{ "actionable_ulid", "task_ulid", "site_ulid", "date_created", "date_modified", "date_deleted" }
	vals := []interface{}{ record.ActionableULID, record.TaskULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted }
	err = h.DB.Upsert("actionable").
		Columns(cols...).
		Values(vals...).
		Where("actionable_ulid = $1", record.ActionableULID).
		Returning("actionable_ulid").
		QueryStruct(record)

	//	if record.ActionableULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("actionable")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("actionable_ulid = $1", record.ActionableULID)
	//		b.Returning("actionable_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("actionable").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("actionable_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *actionableHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("actionable").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and actionable_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *actionableHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("actionable").
		Where("site_ulid=$1 and actionable_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *actionableHelper) validate(record *Actionable) (err error) {
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


