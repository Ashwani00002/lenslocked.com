package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// We need a View to Render the signup page. Accessing User's objet here.
// NewView fileds help us avoid global variables scattered across application.
type Users struct {
	NewViews *views.View
}

// NewUsers is used to create a new Users controller. This function will panic if templates are not
// parsed correclty and should be used during intial setup.
// Setting up new User's object with NewView.
func NewUsers() *Users {
	return &Users{
		NewViews: views.NewViews("bootstrap", "views/users/new.gohtml"),
	}
}

// New is used to Render signup form for creating new user.
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewViews.Render(w, nil)
}

type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create is used to process the signup form when submits it. This is used to create a NewUser account.
// POST /signup

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
