module github.com/nerdynz/schwifty/backend

go 1.14

replace github.com/nerdynz/schwifty/backend/server/models => ../backend/server/models

require (
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/jackc/pgx/v4 v4.8.1
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/lib/pq v1.8.0
	github.com/nerdynz/dat v1.3.0 // indirect
	github.com/nerdynz/datastore v0.0.0-20200402045006-0f63cc077d94
	github.com/nerdynz/rcache v0.0.0-20200404024229-09aee2ea3078
	github.com/nerdynz/schwifty/backend/server/models v0.0.0-00010101000000-000000000000
	github.com/nerdynz/trove v0.0.0-20200425063959-61f6ab2f6311
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/negroni v1.0.0
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/mattes/migrate.v1 v1.3.2 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
