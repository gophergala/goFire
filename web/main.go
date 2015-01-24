package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"idstc.com/model"
)

type ctx struct {
	Database *sqlx.DB
	r        *render.Render
	User     *model.User
}

func main() {

	store := cookiestore.New([]byte(os.Getenv("cookiestore")))

	//init db
	db, err := sqlx.Connect("mysql", os.Getenv("connectionstring"))

	//if err != nil {
	//log.Print(err)
	//log.Fatal("Error Initializing Database...")
	//}

	//setup render
	ren := render.New(render.Options{
		Layout:        "shared/layout",
		IndentJSON:    true,
		IsDevelopment: true,
	})

	session := ctx{nil, ren, nil}

	//create routers
	router := mux.NewRouter()

	n := negroni.Classic()

	//register session middleware
	n.Use(sessions.Sessions("idstc", store))
	n.UseHandler(router)

	//not found handler
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	//open routes
	router.HandleFunc("/", session.Home)
	router.HandleFunc("/people", session.People)
	router.HandleFunc("/projects", session.Projects)
	router.HandleFunc("/account", session.Account)
	router.HandleFunc("/analytics", session.Analytics)

	n.Run(
		fmt.Sprint(":", os.Getenv("PORT")),
	)
}
