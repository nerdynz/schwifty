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

var boardHelperGlobal *boardHelper

// Board Record
type Board struct {
  BoardULID string `db:"board_ulid" json:"BoardULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  ClientULID string `db:"client_ulid" json:"ClientULID"`
  Title string `db:"title" json:"Title"`
  Color string `db:"color" json:"Color"`
  SortPosition int `db:"sort_position" json:"SortPosition"`
  
	Tasks  Tasks  `json:"Tasks"`
	
}

type Boards []*Board

func (h *boardHelper) beforeSave(record *Board) (err error) {
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

func (h *boardHelper) afterSave(record *Board) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type boardHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func BoardHelper() *boardHelper {
	if boardHelperGlobal == nil {
		boardHelperGlobal = newBoardHelper(modelDB, modelCache)
	}
	return boardHelperGlobal
}

func newBoardHelper(db *runner.DB, cache Cache) *boardHelper {
	helper := &boardHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"board_ulid", "board_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "client_ulid", "title", "color", "sort_position"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *boardHelper) New(siteULID string) *Board {
	record := &Board{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *boardHelper) FromRequest(siteULID string, req *http.Request) (*Board, error) {
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
		return nil, errors.New("*Board update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *boardHelper) Load(siteULID string, ulid string) (*Board, error) {
	record, err := h.One(siteULID, "board_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *boardHelper) All(siteULID string) (Boards, error) {
	var records Boards
	err := h.DB.Select("*").
		From("board").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *boardHelper) Where(siteULID string, sql string, args ...interface{}) (Boards, error) {
	var records Boards
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("board").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *boardHelper) SQL(siteULID string, sql string, args ...interface{}) (Boards, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Boards
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *boardHelper) One(siteULID string, sql string, args ...interface{}) (*Board, error) {
	var record Board
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("board").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *boardHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *boardHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Boards
	err := h.DB.Select("*").
		From("board").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(board_ulid) from board where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *boardHelper) Save(siteULID string, record *Board) error {
	return h.save(siteULID, record)
}

func (h *boardHelper) SaveMany(siteULID string, records Boards) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *boardHelper) save(siteULID string, record *Board) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Board update failed. SiteID Mismatch")
	}
	cols := []string{ "board_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "client_ulid", "title", "color", "sort_position" }
	vals := []interface{}{ record.BoardULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.ClientULID, record.Title, record.Color, record.SortPosition }
	err = h.DB.Upsert("board").
		Columns(cols...).
		Values(vals...).
		Where("board_ulid = $1", record.BoardULID).
		Returning("board_ulid").
		QueryStruct(record)

	//	if record.BoardULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("board")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("board_ulid = $1", record.BoardULID)
	//		b.Returning("board_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("board").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("board_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *boardHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("board").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and board_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *boardHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("board").
		Where("site_ulid=$1 and board_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *boardHelper) validate(record *Board) (err error) {
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



func (board *Board) SaveTasks(siteULID string) error {
	return TaskHelper().SaveMany(siteULID, board.Tasks)
}

func (board *Board) LoadTasks(siteULID string) error {
	return board.LoadTasksWhere(siteULID, "board_ulid = $1 $SITEULID", board.BoardULID)
}

func (board *Board) LoadTasksWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := TaskHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	board.Tasks = children
	return nil
}
