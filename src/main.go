package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/akrylysov/algnhsa"
	"github.com/namsral/flag"

	_ "github.com/go-sql-driver/mysql"
)

// Employee struct to define an employee
type Employee struct {
	ID   int
	Name string
	City string
}

var dbHost, dbName, dbUser, dbPass string

func initVariables() {
	flag.StringVar(&dbHost, "dbHost", "localhost", "database host")
	flag.StringVar(&dbName, "dbName", "isen", "database isen")
	flag.StringVar(&dbUser, "dbUser", "", "database user")
	flag.StringVar(&dbPass, "dbPass", "", "database password")
	flag.Parse()
}
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func new(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.ExecuteTemplate(w, "New", nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nID := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM employee WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.Name = name
		emp.City = city
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO employee(name, city) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func contextHandler(w http.ResponseWriter, r *http.Request) {
	proxyReq, ok := algnhsa.ProxyRequestFromContext(r.Context())
	if ok {
		fmt.Fprint(w, proxyReq.RequestContext.AccountID)
	}
}

func main() {

	initVariables()

	_ = dbConn()

	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/new", new)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/context", contextHandler)

	if os.Getenv("_HANDLER") != "" {
		log.Println("Serving requests in Lambda")
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		log.Println("Server started on: http://localhost:8080")
		http.ListenAndServe(":8080", nil)
	}

}
