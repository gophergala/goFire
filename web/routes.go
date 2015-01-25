package main

import (
	"net/http"

	"idstc/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func (c *ctx) Home(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("", w, r, "")

	p := model.Page{
		"dashboard",
		"dashboard",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/index", p)
}

func (c *ctx) People(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("people", w, r, "")

	p := model.Page{
		"people",
		"people",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/people", p)
}

func (c *ctx) Projects(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("projects", w, r, "")

	p := model.Page{
		"projects",
		"projects",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/projects", p)
}

func (c *ctx) ProjectDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	ValidateRoute("projects/detail", w, r, id)

	p := model.Page{
		"project detail",
		"projects",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/projectdetail", p)
}

func (c *ctx) Analytics(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("analytics", w, r, "")

	p := model.Page{
		"analytics",
		"analytics",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/analytics", p)
}

func (c *ctx) Account(w http.ResponseWriter, r *http.Request) {
	ValidateRoute("account", w, r, "")

	p := model.Page{
		"account",
		"account",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/account", p)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	ren := render.New(render.Options{
		Layout: "shared/layout",
	})

	p := model.Page{
		"Page not found",
		"404",
		nil,
	}

	ren.HTML(w, http.StatusOK, "404", p)
}

func (c *ctx) AddTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	ValidateRoute("projects/tasks/add", w, r, id)

	p := model.Page{
		"Add Task",
		"projects",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/addtask", p)
}

func (c *ctx) AddProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	ValidateRoute("projects/add", w, r, id)

	p := model.Page{
		"Add Project",
		"projects",
		nil,
	}

	c.r.HTML(w, http.StatusOK, "home/addproject", p)
}
