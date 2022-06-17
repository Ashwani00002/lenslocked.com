package controllers

import "lenslocked.com/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewViews("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewViews("boootstrap", "views/static/contact.gohtml"),
	}
}
