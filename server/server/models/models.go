package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	runner "github.com/nerdynz/dat/sqlx-runner"
	"github.com/nerdynz/datastore"
	"github.com/pinzolo/casee"
)

const NoRows = "sql: no rows in result set"

// var modelValidator *validator.Validate
var modelDB *runner.DB
var modelStore *datastore.Datastore

var modelCache Cache
var modelDecoder *schema.Decoder

type Cache interface {
	// GetSectretKey() string
	// GetApiKey() string
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
}

// var modelDecoder *schema.Decoder
// type ULID string

func Init(store *datastore.Datastore) {
	// modelValidator = validator.New()
	// modelDecoder = schema.NewDecoder()
	// modelDecoder.IgnoreUnknownKeys(true)
	modelStore = store
	modelDB = store.DB
	modelCache = store.Cache
}

type PagedData struct {
	Sort      string      `json:"sort"`
	Direction string      `json:"direction"`
	Records   interface{} `json:"records"`
	Total     int         `json:"total"`
	PageNum   int         `json:"pageNum"`
	Limit     int         `json:"limit"`
}

func NewPagedData(records interface{}, orderBy string, direction string, itemsPerPage int, pageNum int, total int) *PagedData {
	return &PagedData{
		Records:   records,
		Direction: direction,
		Sort:      casee.ToPascalCase(orderBy),
		Limit:     itemsPerPage,
		PageNum:   pageNum,
		Total:     total,
	}
}

func appendSiteID(siteID int, whereSQLOrMap string, args ...interface{}) (string, []interface{}) {
	args = append(args, siteID)
	position := len(args)
	if strings.Contains(whereSQLOrMap, ".$SITEID") {
		newSQL := strings.Split(whereSQLOrMap, "$SITEID")[0]
		replaceSQLParts := strings.Split(newSQL, " ")
		replaceSQLTablePrefix := replaceSQLParts[len(replaceSQLParts)-1]

		whereSQLOrMap = strings.Replace(whereSQLOrMap, replaceSQLTablePrefix+"$SITEID", " and "+replaceSQLTablePrefix+"site_id = $"+strconv.Itoa(position), -1)
	} else if strings.Contains(whereSQLOrMap, "$SITEID") {
		whereSQLOrMap = strings.Replace(whereSQLOrMap, "$SITEID", " and site_id = $"+strconv.Itoa(position), -1)
	} else {
		whereSQLOrMap += " and site_id = $" + strconv.Itoa(position)
	}
	return whereSQLOrMap, args
}

func appendSiteULID(siteULID string, whereSQLOrMap string, args ...interface{}) (string, []interface{}) {
	args = append(args, siteULID)
	position := len(args)
	if strings.Contains(whereSQLOrMap, ".$SITEULID") {
		newSQL := strings.Split(whereSQLOrMap, "$SITEULID")[0]
		replaceSQLParts := strings.Split(newSQL, " ")
		replaceSQLTablePrefix := replaceSQLParts[len(replaceSQLParts)-1]

		whereSQLOrMap = strings.Replace(whereSQLOrMap, replaceSQLTablePrefix+"$SITEULID", " and "+replaceSQLTablePrefix+"site_ulid = $"+strconv.Itoa(position), -1)
	} else if strings.Contains(whereSQLOrMap, "$SITEULID") {
		whereSQLOrMap = strings.Replace(whereSQLOrMap, "$SITEULID", " and site_ulid = $"+strconv.Itoa(position), -1)
	} else {
		whereSQLOrMap += " and site_ulid = $" + strconv.Itoa(position)
	}
	return whereSQLOrMap, args
}
