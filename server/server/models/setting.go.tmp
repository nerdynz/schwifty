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

var settingHelperGlobal *settingHelper

// Setting Record
type Setting struct {
  SettingULID string `db:"setting_ulid" json:"SettingULID"`
  SiteULID string `db:"site_ulid" json:"SiteULID"`
  DateCreated time.Time `db:"date_created" json:"DateCreated"`
  DateModified time.Time `db:"date_modified" json:"DateModified"`
  DateDeleted time.Time `db:"date_deleted" json:"DateDeleted"`
  LogoPicture string `db:"logo_picture" json:"LogoPicture"`
  PrimaryColor string `db:"primary_color" json:"PrimaryColor"`
  SecondaryColor string `db:"secondary_color" json:"SecondaryColor"`
  
	
}

type Settings []*Setting

func (h *settingHelper) beforeSave(record *Setting) (err error) {
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

func (h *settingHelper) afterSave(record *Setting) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type settingHelper struct {
	DB            *runner.DB
	Cache         Cache
	fieldNames    []string
	orderBy       string
}

func SettingHelper() *settingHelper {
	if settingHelperGlobal == nil {
		settingHelperGlobal = newSettingHelper(modelDB, modelCache)
	}
	return settingHelperGlobal
}

func newSettingHelper(db *runner.DB, cache Cache) *settingHelper {
	helper := &settingHelper{}
	helper.DB = db
	helper.Cache = cache

	// Fields
	fieldnames := []string{"setting_ulid", "setting_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "logo_picture", "primary_color", "secondary_color"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames
	
	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *settingHelper) New(siteULID string) *Setting {
	record := &Setting{}
	// check DateCreated
	record.DateCreated = time.Now()
	record.SiteULID = siteULID
	return record
}

func (h *settingHelper) FromRequest(siteULID string, req *http.Request) (*Setting, error) {
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
		return nil, errors.New("*Setting update failed. Site ULID Mismatch")
	}
	record.SiteULID = siteULID
	return record, nil
}


func (h *settingHelper) Load(siteULID string, ulid string) (*Setting, error) {
	record, err := h.One(siteULID, "setting_ulid = $1", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *settingHelper) All(siteULID string) (Settings, error) {
	var records Settings
	err := h.DB.Select("*").
		From("setting").
		Where("site_ulid = $1", siteULID).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *settingHelper) Where(siteULID string, sql string, args ...interface{}) (Settings, error) {
	var records Settings
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.Select("*").
		From("setting").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *settingHelper) SQL(siteULID string, sql string, args ...interface{}) (Settings, error) {
	if !strings.Contains(sql, "$SITEID") {
		return nil, errors.New("No $SITEID placeholder defined")
	}
	var records Settings
	sql, args = appendSiteULID(siteULID, sql, args...)
	err := h.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *settingHelper) One(siteULID string, sql string, args ...interface{}) (*Setting, error) {
	var record Setting
	sql, args = appendSiteULID(siteULID, sql, args...)

	err := h.DB.Select("*").
		From("setting").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *settingHelper) Paged(siteULID string, pageNum int, itemsPerPage int) (*PagedData, error) {
	pd, err := h.PagedBy(siteULID, pageNum, itemsPerPage, "date_created", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *settingHelper) PagedBy(siteULID string, pageNum int, itemsPerPage int, orderByFieldName string, direction string) (*PagedData, error) {
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

	var records Settings
	err := h.DB.Select("*").
		From("setting").
		Where("site_ulid = $1", siteULID).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	count := 0
	h.DB.SQL(`select count(setting_ulid) from setting where site_ulid = $1`, siteULID).QueryStruct(&count)
	return NewPagedData(records, orderByFieldName, direction, itemsPerPage, pageNum, count), nil
}

func (h *settingHelper) Save(siteULID string, record *Setting) error {
	return h.save(siteULID, record)
}

func (h *settingHelper) SaveMany(siteULID string, records Settings) error {
	for _, record := range records {
		err := h.save(siteULID, record)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *settingHelper) save(siteULID string, record *Setting) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteULID != siteULID {
		return errors.New("*Setting update failed. SiteID Mismatch")
	}
	cols := []string{ "setting_ulid", "site_ulid", "date_created", "date_modified", "date_deleted", "logo_picture", "primary_color", "secondary_color" }
	vals := []interface{}{ record.SettingULID, record.SiteULID, record.DateCreated, record.DateModified, record.DateDeleted, record.LogoPicture, record.PrimaryColor, record.SecondaryColor }
	err = h.DB.Upsert("setting").
		Columns(cols...).
		Values(vals...).
		Where("setting_ulid = $1", record.SettingULID).
		Returning("setting_ulid").
		QueryStruct(record)

	//	if record.SettingULID != "" {
	//		// UPDATE
	//		b := h.DB.Update("setting")
	//		for i := range cols {
	//			b.Set(cols[i], vals[i])
	//		}
	//		b.Where("setting_ulid = $1", record.SettingULID)
	//		b.Returning("setting_ulid")
	//		err = b.QueryStruct(record)
	//	} else {
	//		// INSERT
	//		err = h.DB.
	//			InsertInto("setting").
	//			Columns(cols...).
	//			Values(vals...).
	//			Returning("setting_ulid").
	//			QueryStruct(record)
	//	}
	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *settingHelper) Delete(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		Update("setting").
		Set("date_deleted", time.Now()).
		Where("site_ulid=$1 and setting_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *settingHelper) Purge(siteULID string, recordULID string) (bool, error) {
	result, err := h.DB.
		DeleteFrom("setting").
		Where("site_ulid=$1 and setting_ulid=$2", siteULID, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *settingHelper) validate(record *Setting) (err error) {
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


