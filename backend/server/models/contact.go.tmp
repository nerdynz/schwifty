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

var contactHelperGlobal *contactHelper

// Contact Record
type Contact struct {
  ContactULID string `db:"contact_ulid" json:"ContactULID"`
  ClientULID string `db:"client_ulid" json:"ClientULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  
	
}

type Contacts []*Contact

func (h *contactHelper) beforeSave(record *Contact) (err error) {
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

func (h *contactHelper) afterSave(record *Contact) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type contactHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func ContactHelper() *contactHelper {
	if contactHelperGlobal == nil {
		contactHelperGlobal = newContactHelper(modelDB, modelCache)
	}
	return contactHelperGlobal
}

func newContactHelper(db *runner.DB, cache Cache) *contactHelper {
	helper := &contactHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"contact_ulid", "contact_ulid", "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *contactHelper) New(siteULID string) *Contact {
	record := &Contact{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *contactHelper) FromRequest(siteULID string, req *http.Request) (*Contact, error) {
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
		return nil, errors.New("*Contact update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *contactHelper) Load(siteULID string, ulid string) (*Contact, error) {
	record, err := h.One(siteULID, "contact_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *contactHelper) All(siteULID string) (Contacts, error) {
	var records Contacts
	err := h.DB.Select("*").
		From("contact").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *contactHelper) Where(siteULID string, sql string, args ...interface{}) (Contacts, error) {
	var records Contacts
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("contact").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *contactHelper) SQL(siteULID string, sql string, args ...interface{}) (Contacts, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Contacts
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *contactHelper) One(siteULID string, sql string, args ...interface{}) (*Contact, error) {
	var record Contact
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("contact").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *contactHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *contactHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Contacts
	err := h.DB.Select("*").
		From("contact").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(contact_ulid) from contact where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *contactHelper) Save(siteULID string, record *Contact) error {
	return h.save(siteULID, record)
}

func (h *contactHelper) SaveMany(siteULID string, records Contacts) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *contactHelper) save(siteULID string, record *Contact) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Contact update failed. SiteID Mismatch")
	}
	cols := []string{ "contact_ulid", "client_ulid", "site_ulid", "date_created", "date_modified", "date_deleted" }
	vals := []interface{}{ record.ContactULID, record.ClientULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted }
	err = h.DB.Upsert("contact").
		Columns(cols...).
		Values(vals...).
		Where("contact_ulid = $1", record.ContactULID).
		Returning("contact_ulid").
		QueryStruct(record)

	//	if record.ContactULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("contact")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("contact_ulid = $1", record.ContactULID)
	//		b.Returning("contact_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("contact").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("contact_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *contactHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("contact").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and contact_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *contactHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("contact").
		Where("site_ulid=$1 and contact_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *contactHelper) validate(record *Contact) (err error) {
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


