package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"

	runner "github.com/nerdynz/dat/sqlx-runner"
)

var messageHelperGlobal *messageHelper

// Message Record
type Message struct {
	MessageULID  string    `db:"message_ulid" json:"MessageULID"`
	DateCreated  time.Time `db:"date_created" json:"DateCreated"`
	DateModified time.Time `db:"date_modified" json:"DateModified"`
	DateDeleted  time.Time `db:"date_deleted" json:"DateDeleted"`
	SiteULID     string    `db:"site_ulid" json:"SiteULID"`
	BoardULID    string    `db:"board_ulid" json:"BoardULID"`
	Message      string    `db:"message" json:"Message"`
	PersonULID   string    `db:"person_ulid" json:"PersonULID"`
	TaskULID     string    `db:"task_ulid" json:"TaskULID"`
}

type Messages []*Message

func (h *messageHelper) beforeSave(record *Message) (err error) {
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

func (h *messageHelper) afterSave(record *Message) (err error) {
	saveActivity(record.MessageULID, "Message", record.SiteULID)
	return err
}

// GENERATED CODE - Leave the below code alone
type messageHelper struct {
	DB         *runner.DB
	Cache      Cache
	fieldNames []string
	orderBy    string
}

func MessageHelper() *messageHelper {
	if messageHelperGlobal == nil {
		messageHelperGlobal = newMessageHelper(modelDB, modelCache)
	}
	return messageHelperGlobal
}

func newMessageHelper(db *runner.DB, cache Cache) *messageHelper {
	helper := &messageHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"message_ulid", "message_ulid", "date_created", "date_modified", "date_deleted", "site_ulid", "board_ulid", "message", "person_ulid", "task_ulid"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames

	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *messageHelper) New(siteULID string) *Message {
	record := &Message{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *messageHelper) FromRequest(siteULID string, req *http.Request) (*Message, error) {
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
		return nil, errors.New("*Message update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}

func (h *messageHelper) Load(siteULID string, ulid string) (*Message, error) {
	record, err := h.One(siteULID, "message_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *messageHelper) All(siteULID string) (Messages, error) {
	var records Messages
	err := h.DB.Select("*").
		From("message").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *messageHelper) Where(siteULID string, sql string, args ...interface{}) (Messages, error) {
	var records Messages
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("message").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *messageHelper) SQL(siteULID string, sql string, args ...interface{}) (Messages, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Messages
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *messageHelper) One(siteULID string, sql string, args ...interface{}) (*Message, error) {
	var record Message
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("message").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *messageHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *messageHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Messages
	err := h.DB.Select("*").
		From("message").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(message_ulid) from message where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *messageHelper) Save(siteULID string, record *Message) error {
	return h.save(siteULID, record)
}

func (h *messageHelper) SaveMany(siteULID string, records Messages) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *messageHelper) save(siteULID string, record *Message) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Message update failed. SiteID Mismatch")
	}
	cols := []string{"message_ulid", "date_created", "date_modified", "date_deleted", "site_ulid", "board_ulid", "message", "person_ulid", "task_ulid"}
	vals := []interface{}{record.MessageULID, record.DateCreated, record.DateModified, record.DateDeleted, record.SiteULID, record.BoardULID, record.Message, record.PersonULID, record.TaskULID}
	err = h.DB.Upsert("message").
		Columns(cols...).
		Values(vals...).
		Where("message_ulid = $1", record.MessageULID).
		Returning("message_ulid").
		QueryStruct(record)

	//	if record.MessageULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("message")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("message_ulid = $1", record.MessageULID)
	//		b.Returning("message_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("message").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("message_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *messageHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("message").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and message_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *messageHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("message").
		Where("site_ulid=$1 and message_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *messageHelper) validate(record *Message) (err error) {
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
