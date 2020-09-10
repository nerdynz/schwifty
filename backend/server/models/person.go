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

var personHelperGlobal *personHelper

// Person Record
type Person struct {
  PersonULID string `db:"person_ulid" json:"PersonULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  Name string `db:"name" json:"Name"`
  Email string `db:"email" json:"Email"`
  Phone string `db:"phone" json:"Phone"`
  Role string `db:"role" json:"Role"`
  Initials string `db:"initials" json:"Initials"`
  Password string `db:"password" json:"Password"`
  
	
}

type People []*Person

func (h *personHelper) beforeSave(record *Person) (err error) {
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

func (h *personHelper) afterSave(record *Person) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type personHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func PersonHelper() *personHelper {
	if personHelperGlobal == nil {
		personHelperGlobal = newPersonHelper(modelDB, modelCache)
	}
	return personHelperGlobal
}

func newPersonHelper(db *runner.DB, cache Cache) *personHelper {
	helper := &personHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"person_ulid", "person_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "name", "email", "phone", "role", "initials", "password"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *personHelper) New(siteULID string) *Person {
	record := &Person{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *personHelper) FromRequest(siteULID string, req *http.Request) (*Person, error) {
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
		return nil, errors.New("*Person update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *personHelper) Load(siteULID string, ulid string) (*Person, error) {
	record, err := h.One(siteULID, "person_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *personHelper) All(siteULID string) (People, error) {
	var records People
	err := h.DB.Select("*").
		From("person").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *personHelper) Where(siteULID string, sql string, args ...interface{}) (People, error) {
	var records People
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("person").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *personHelper) SQL(siteULID string, sql string, args ...interface{}) (People, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records People
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *personHelper) One(siteULID string, sql string, args ...interface{}) (*Person, error) {
	var record Person
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("person").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *personHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *personHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records People
	err := h.DB.Select("*").
		From("person").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(person_ulid) from person where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *personHelper) Save(siteULID string, record *Person) error {
	return h.save(siteULID, record)
}

func (h *personHelper) SaveMany(siteULID string, records People) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *personHelper) save(siteULID string, record *Person) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Person update failed. SiteID Mismatch")
	}
	cols := []string{ "person_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "name", "email", "phone", "role", "initials", "password" }
	vals := []interface{}{ record.PersonULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.Name, record.Email, record.Phone, record.Role, record.Initials, record.Password }
	err = h.DB.Upsert("person").
		Columns(cols...).
		Values(vals...).
		Where("person_ulid = $1", record.PersonULID).
		Returning("person_ulid").
		QueryStruct(record)

	//	if record.PersonULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("person")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("person_ulid = $1", record.PersonULID)
	//		b.Returning("person_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("person").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("person_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *personHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("person").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and person_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *personHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("person").
		Where("site_ulid=$1 and person_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *personHelper) validate(record *Person) (err error) {
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


