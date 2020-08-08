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

var taskHelperGlobal *taskHelper

// Task Record
type Task struct {
  TaskULID string `db:"task_ulid" json:"TaskULID"`
  BoardULID string `db:"board_ulid" json:"BoardULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  Title string `db:"title" json:"Title"`
  Notes string `db:"notes" json:"Notes"`
  Status string `db:"status" json:"Status"`
  SortPosition int `db:"sort_position" json:"SortPosition"`
  
	Actionables  Actionables  `json:"Actionables"`
	TimeEntries  TimeEntries  `json:"TimeEntries"`
	
}

type Tasks []*Task

func (h *taskHelper) beforeSave(record *Task) (err error) {
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

func (h *taskHelper) afterSave(record *Task) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type taskHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func TaskHelper() *taskHelper {
	if taskHelperGlobal == nil {
		taskHelperGlobal = newTaskHelper(modelDB, modelCache)
	}
	return taskHelperGlobal
}

func newTaskHelper(db *runner.DB, cache Cache) *taskHelper {
	helper := &taskHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"task_ulid", "task_ulid", "board_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "title", "notes", "status", "sort_position"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *taskHelper) New(siteULID string) *Task {
	record := &Task{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *taskHelper) FromRequest(siteULID string, req *http.Request) (*Task, error) {
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
		return nil, errors.New("*Task update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *taskHelper) Load(siteULID string, ulid string) (*Task, error) {
	record, err := h.One(siteULID, "task_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *taskHelper) All(siteULID string) (Tasks, error) {
	var records Tasks
	err := h.DB.Select("*").
		From("task").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *taskHelper) Where(siteULID string, sql string, args ...interface{}) (Tasks, error) {
	var records Tasks
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("task").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *taskHelper) SQL(siteULID string, sql string, args ...interface{}) (Tasks, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Tasks
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *taskHelper) One(siteULID string, sql string, args ...interface{}) (*Task, error) {
	var record Task
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("task").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *taskHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *taskHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Tasks
	err := h.DB.Select("*").
		From("task").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(task_ulid) from task where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *taskHelper) Save(siteULID string, record *Task) error {
	return h.save(siteULID, record)
}

func (h *taskHelper) SaveMany(siteULID string, records Tasks) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *taskHelper) save(siteULID string, record *Task) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Task update failed. SiteID Mismatch")
	}
	cols := []string{ "task_ulid", "board_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "title", "notes", "status", "sort_position" }
	vals := []interface{}{ record.TaskULID, record.BoardULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.Title, record.Notes, record.Status, record.SortPosition }
	err = h.DB.Upsert("task").
		Columns(cols...).
		Values(vals...).
		Where("task_ulid = $1", record.TaskULID).
		Returning("task_ulid").
		QueryStruct(record)

	//	if record.TaskULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("task")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("task_ulid = $1", record.TaskULID)
	//		b.Returning("task_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("task").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("task_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *taskHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("task").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and task_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *taskHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("task").
		Where("site_ulid=$1 and task_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *taskHelper) validate(record *Task) (err error) {
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



func (task *Task) SaveActionables(siteULID string) error {
	return ActionableHelper().SaveMany(siteULID, task.Actionables)
}

func (task *Task) LoadActionables(siteULID string) error {
	return task.LoadActionablesWhere(siteULID, "task_ulid = $1 $SITEULID", task.TaskULID)
}

func (task *Task) LoadActionablesWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := ActionableHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	task.Actionables = children
	return nil
}

func (task *Task) SaveTimeEntries(siteULID string) error {
	return TimeEntryHelper().SaveMany(siteULID, task.TimeEntries)
}

func (task *Task) LoadTimeEntries(siteULID string) error {
	return task.LoadTimeEntriesWhere(siteULID, "task_ulid = $1 $SITEULID", task.TaskULID)
}

func (task *Task) LoadTimeEntriesWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := TimeEntryHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	task.TimeEntries = children
	return nil
}
