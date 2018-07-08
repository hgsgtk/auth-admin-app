package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Staff struct {
	gorm.Model
	ID         int64
	Name, Role string
}

func staffIndexHandler(w http.ResponseWriter, r *http.Request) *appError {
	// FIXME: 共通化
	// sqlite3だろうがmysqlだろうがinterfaceでの抽象化
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	staffs := []Staff{}
	db.Debug().Find(&staffs)
	return staffIndexTmpl.Execute(w, r, staffs)
}

func staffAddHandler(w http.ResponseWriter, r *http.Request) *appError {
	// FIXME: 共通化
	// sqlite3だろうがmysqlだろうがinterfaceでの抽象化
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return staffEditTmpl.Execute(w, r, nil)
}

func staffDetailHandler(w http.ResponseWriter, r *http.Request) *appError {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "bad book id: %v", err)
	}
	// FIXME: 共通化
	// sqlite3だろうがmysqlだろうがinterfaceでの抽象化
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	staffs := []Staff{}
	db.Where("id = ?", id).First(&staffs)
	return staffDetailTmpl.Execute(w, r, staffs)
}

func staffCreateHandler(w http.ResponseWriter, r *http.Request) *appError {
	staff := &Staff{
		Name: r.FormValue("name"),
		Role: r.FormValue("role"),
	}

	// FIXME: 共通化
	// sqlite3だろうがmysqlだろうがinterfaceでの抽象化
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Create(&staff)

	http.Redirect(w, r, "/staffs/", http.StatusFound)
	return nil
}

func staffUpdateHandler(w http.ResponseWriter, r *http.Request) *appError {
	// Form validation
	name := r.FormValue("name")
	role := r.FormValue("role")

	// FIXME: 共通化
	// sqlite3だろうがmysqlだろうがinterfaceでの抽象化
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Create(&Staff{Name: name, Role: role})

	http.Redirect(w, r, "/staffs/", http.StatusFound)
	return nil
}
