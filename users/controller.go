package users

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Controller struct {
	Repository Repository
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	users := c.Repository.GetUsers()
	tmpl.ExecuteTemplate(w, "Index", users)

	return
}

func (c *Controller) New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)

	return
}

func (c *Controller) InsertUser(w http.ResponseWriter, r *http.Request) {
	var user User
	age, _ := strconv.Atoi(r.FormValue("age"))
	phone, _ := strconv.ParseInt(r.FormValue("phone"), 10, 32)

	user.FullName = r.FormValue("name")
	user.Age = age
	user.Phone = int32(phone)
	user.Address = r.FormValue("address")

	err := c.Repository.InsertUser(user)
	checkErr(err)

	users := c.Repository.GetUsers()
	tmpl.ExecuteTemplate(w, "Index", users)

	return
}

func (c *Controller) SearchUser(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	users := c.Repository.GetUsersByString(query)
	tmpl.ExecuteTemplate(w, "Index", users)

	return
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *User
	id, _ := strconv.Atoi(r.FormValue("id"))
	age, _ := strconv.Atoi(r.FormValue("age"))
	phone, _ := strconv.ParseInt(r.FormValue("phone"), 10, 32)

	user.ID = id
	user.FullName = r.FormValue("name")
	user.Age = age
	user.Phone = int32(phone)
	user.Address = r.FormValue("address")

	err := c.Repository.UpdateUser(user)
	checkErr(err)
	tmpl.ExecuteTemplate(w, "Show", user)

	return
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userid, err := strconv.Atoi(id)
	checkErr(err)

	user := c.Repository.GetUserById(userid)
	tmpl.ExecuteTemplate(w, "Show", user)

	return
}

func (c *Controller) GetUserForEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userid, err := strconv.Atoi(id)
	checkErr(err)

	user := c.Repository.GetUserById(userid)
	tmpl.ExecuteTemplate(w, "Edit", user)

	return
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId, err := strconv.Atoi(id)
	checkErr(err)

	err = c.Repository.DeleteUser(userId)
	checkErr(err)

	users := c.Repository.GetUsers()
	tmpl.ExecuteTemplate(w, "Index", users)

	return
}
