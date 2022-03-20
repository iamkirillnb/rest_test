package main

import (
	"flag"
	"fmt"
	"github.com/iamkirillnb/rus_law/internal"
	"github.com/iamkirillnb/rus_law/internal/handlers"
	"github.com/iamkirillnb/rus_law/internal/repos"
	"github.com/iamkirillnb/rus_law/pkg/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var cfgPath string

func init(){
	flag.StringVar(&cfgPath, "config", "dev.yaml", "config file path")
}

func main() {
	fmt.Println("Law APP start")

	flag.Parse()

	log := logging.NewLogger()
	cfg := internal.GetConfig(cfgPath)

	connStr := cfg.DbConfig.DbUrl()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("connection to DB fail")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("ping to DB fail")
	}

	repoLaw := repos.NewDbLaw(db, &log)
	data, err := repoLaw.GetAll()
	if err != nil {
		log.Fatal("get data from DB fail")
	}

	fmt.Println(data)

	handler := handlers.NewHandler(log)

	log.Fatal(handler.ListenAndServe())






}
