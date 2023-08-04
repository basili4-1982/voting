package main

import (
	"net/http"
	"os"
	"strconv"
	"time"
	"voting/pkg/handler"
	"voting/pkg/poll"
	"voting/pkg/storage"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	listen := os.Getenv("LISTEN")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 8)
	if err != nil {
		dbPort = 5432
	}
	dbUserName := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Info)
	db, err := storage.OpenDB(dbHost, int(dbPort), dbUserName, dbPassword, dbName)
	if err != nil {
		panic(err)
	}
	dbRep := storage.NewDbRep(db)
	r.HandleFunc("/createPoll", handler.CreatePoll(poll.NewPoll(dbRep))).Methods("POST")
	r.HandleFunc("/getResult", handler.GetResult(poll.NewPoll(dbRep))).Methods("GET")
	r.HandleFunc("/poll", handler.Poll(poll.NewPoll(dbRep))).Methods("POST")
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    listen,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Запуск слушаю порт ", listen)
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
