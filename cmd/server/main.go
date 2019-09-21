package main

import (
	"../../internal/hooks"
	"../../internal/psuserver"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/speakerkfm/web-crud/pkg/service"
	"github.com/speakerkfm/web-crud/pkg/store"
	"github.com/speakerkfm/web-crud/rpc/psu"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//mysql
	mysqlConf := mysql.NewConfig()
	mysqlConf.Net = "tcp"
	mysqlConf.Addr = os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT")
	mysqlConf.User = os.Getenv("DATABASE_USER")
	mysqlConf.Passwd = os.Getenv("DATABASE_PASSWORD")
	mysqlConf.DBName = os.Getenv("DATABASE_NAME")
	mysqlConf.MultiStatements = true
	mysqlConf.ParseTime = true
	mysqlConf.Loc = time.Local

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)

	st := store.NewStore(db)
	StudentService := service.NewStudentService(st)

	server := &psuserver.Server{
		StudentService: StudentService,
	}

	hook := hooks.LoggingHooks(os.Stderr)
	twirpHandler := psu.NewPsuServer(server, hook)

	http.ListenAndServe(":8080", twirpHandler)
}
