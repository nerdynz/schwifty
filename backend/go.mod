module github.com/schwifty/backend

go 1.14

replace github.com/nerdynz/flow => ../../flow

replace github.com/nerdynz/security => ../../security

replace github.com/nerdynz/schwifty/backend/server/models => ./server/models

replace github.com/nerdynz/schwifty/backend/server => ./server

require (
	github.com/lib/pq v1.7.1 // indirect
	github.com/nats-io/nats-server/v2 v2.1.8 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nerdynz/datastore v0.0.0-20200402045006-0f63cc077d94
	github.com/nerdynz/rcache v0.0.0-20200404024229-09aee2ea3078
	github.com/nerdynz/schwifty/backend/server v0.0.0-20200808111049-d4a553d250b1
	github.com/nerdynz/schwifty/backend/server/models v0.0.0-20200808111049-d4a553d250b1
	github.com/nerdynz/trove v0.0.0-20200425063959-61f6ab2f6311
	github.com/rs/cors v1.7.0
	github.com/shomali11/xredis v0.0.0-20190608143638-0b54a6bbf40b // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/negroni v1.0.0
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	golang.org/x/sys v0.0.0-20200720211630-cb9d2d5c5666 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
