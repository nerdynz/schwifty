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

var clientHelperGlobal *clientHelper

// Client Record
type Client struct {
  ClientULID string `db:"client_ulid" json:"ClientULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  Address string `db:"address" json:"Address"`
  Name string `db:"name" json:"Name"`
  Rate int `db:"rate" json:"Rate"`
  
	Boards  Boards  `json:"Boards"`
	Contacts  Contacts  `json:"Contacts"`
	Jobs  Jobs  `json:"Jobs"`
	TimeEntries  TimeEntries  `json:"TimeEntries"`
	
}

type Clients []*Client

func (h *clientHelper) beforeSave(record *Client) (err error) {
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

func (h *clientHelper) afterSave(record *Client) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type clientHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func ClientHelper() *clientHelper {
	if clientHelperGlobal == nil {
		clientHelperGlobal = newClientHelper(modelDB, modelCache)
	}
	return clientHelperGlobal
}

func newClientHelper(db *runner.DB, cache Cache) *clientHelper {
	helper := &clientHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"client_ulid", "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "address", "name", "rate"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *clientHelper) New(siteULID string) *Client {
	record := &Client{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *clientHelper) FromRequest(siteULID string, req *http.Request) (*Client, error) {
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
		return nil, errors.New("*Client update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *clientHelper) Load(siteULID string, ulid string) (*Client, error) {
	record, err := h.One(siteULID, "client_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *clientHelper) All(siteULID string) (Clients, error) {
	var records Clients
	err := h.DB.Select("*").
		From("client").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *clientHelper) Where(siteULID string, sql string, args ...interface{}) (Clients, error) {
	var records Clients
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("client").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *clientHelper) SQL(siteULID string, sql string, args ...interface{}) (Clients, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Clients
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *clientHelper) One(siteULID string, sql string, args ...interface{}) (*Client, error) {
	var record Client
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("client").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *clientHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *clientHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Clients
	err := h.DB.Select("*").
		From("client").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(client_ulid) from client where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *clientHelper) Save(siteULID string, record *Client) error {
	return h.save(siteULID, record)
}

func (h *clientHelper) SaveMany(siteULID string, records Clients) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *clientHelper) save(siteULID string, record *Client) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Client update failed. SiteID Mismatch")
	}
	cols := []string{ "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "address", "name", "rate" }
	vals := []interface{}{ record.ClientULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.Address, record.Name, record.Rate }
	err = h.DB.Upsert("client").
		Columns(cols...).
		Values(vals...).
		Where("client_ulid = $1", record.ClientULID).
		Returning("client_ulid").
		QueryStruct(record)

	//	if record.ClientULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("client")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("client_ulid = $1", record.ClientULID)
	//		b.Returning("client_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("client").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("client_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *clientHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("client").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and client_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *clientHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("client").
		Where("site_ulid=$1 and client_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *clientHelper) validate(record *Client) (err error) {
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



func (client *Client) SaveBoards(siteULID string) error {
	return BoardHelper().SaveMany(siteULID, client.Boards)
}

func (client *Client) LoadBoards(siteULID string) error {
	return client.LoadBoardsWhere(siteULID, "client_ulid = $1 $SITEULID", client.ClientULID)
}

func (client *Client) LoadBoardsWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := BoardHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	client.Boards = children
	return nil
}

func (client *Client) SaveContacts(siteULID string) error {
	return ContactHelper().SaveMany(siteULID, client.Contacts)
}

func (client *Client) LoadContacts(siteULID string) error {
	return client.LoadContactsWhere(siteULID, "client_ulid = $1 $SITEULID", client.ClientULID)
}

func (client *Client) LoadContactsWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := ContactHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	client.Contacts = children
	return nil
}

func (client *Client) SaveJobs(siteULID string) error {
	return JobHelper().SaveMany(siteULID, client.Jobs)
}

func (client *Client) LoadJobs(siteULID string) error {
	return client.LoadJobsWhere(siteULID, "client_ulid = $1 $SITEULID", client.ClientULID)
}

func (client *Client) LoadJobsWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := JobHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	client.Jobs = children
	return nil
}

func (client *Client) SaveTimeEntries(siteULID string) error {
	return TimeEntryHelper().SaveMany(siteULID, client.TimeEntries)
}

func (client *Client) LoadTimeEntries(siteULID string) error {
	return client.LoadTimeEntriesWhere(siteULID, "client_ulid = $1 $SITEULID", client.ClientULID)
}

func (client *Client) LoadTimeEntriesWhere(siteULID string, sql string, args ...interface{}) error {
	children, err := TimeEntryHelper().Where(siteULID, sql, args...)
	if err != nil {
		return err
	}
	client.TimeEntries = children
	return nil
}
