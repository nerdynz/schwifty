package main

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/nerdynz/trove"
	"github.com/sirupsen/logrus"
)

// func main() {
// 	// settings := trove.Load()
// 	// db, err := sql.Open("postgres", settings.Get("DATABASE_URL"))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
// 	// log := logrus.New()
// 	// cache := rcache.New(settings.Get("REDIS_URL"), log)
// 	// store := datastore.New(log, settings, cache, nil)

// 	// models.Init(store)

// 	ctx := context.Background()
// 	con, err := pgx.Connect(ctx, settings.Get("DATABASE_URL"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = con.Ping(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for {
// 		logrus.Info("hi")
// 		ctx := context.Background()
// 		n, err := con.WaitForNotification(ctx)
// 		if err != nil {
// 			logrus.Error("xx", err)
// 			continue
// 		}
// 		logrus.Info("xxx")
// 		logrus.Info(n.Payload)
// 	}
// }

func main() {
	settings := trove.Load()
	conn := settings.Get("DATABASE_URL")

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	minReconn := 10 * time.Second
	maxReconn := time.Minute
	l := pq.NewListener(conn, minReconn, maxReconn, reportProblem)
	err := l.Listen("getwork")
	if err != nil {
		panic(err)
	}
	fmt.Println("entering main loop")
	for {
		select {
		case <-l.Notify:
			logrus.Info("asdf")
		case <-time.After(2 * time.Second):
			go l.Ping()
			// Check if there's more work available, just in case it takes
			// a while for the Listener to notice connection loss and
			// reconnect.
			fmt.Println("received no work for 90 seconds, checking for new work")
		}
	}
}
