package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver

	"github.com/dkantikorn/go-fullstack/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) error {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		"postgres")

	// connect to the postgres db just to be able to run the create db statement
	db, err := gorm.Open(Dbdriver, connStr)
	if err != nil {
		return err
	}

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", DbName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		return rs.Error
	}

	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s WITH OWNER = %s ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8' TABLESPACE = pg_default CONNECTION LIMIT = -1;", DbName, DbUser)
		if rs := db.Exec(stmt); rs.Error != nil {
			return rs.Error
		}

		// close db connection
		defer db.DB().Close()
	}
	return nil
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	if Dbdriver == "mysql" {
		if err := server.CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName); err != nil {
			fmt.Printf("Cannot create database %s with %s database server", DbName, Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			var err error
			DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
			server.DB, err = gorm.Open(Dbdriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database", Dbdriver)
				log.Fatal("This is the error:", err)
			} else {
				fmt.Printf("We are connected to the %s database", Dbdriver)
			}
		}
	}
	if Dbdriver == "postgres" {
		if err := server.CreateDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName); err != nil {
			fmt.Printf("Cannot create database %s with %s database server", DbName, Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			var err error
			DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
			server.DB, err = gorm.Open(Dbdriver, DBURL)
			if err != nil {
				fmt.Printf("Cannot connect to %s database", Dbdriver)
				log.Fatal("This is the error:", err)
			} else {
				fmt.Printf("We are connected to the %s database", Dbdriver)
			}
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) //database migration

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port: " + addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
