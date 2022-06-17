package controllers

import "lenslocked.com/views"

type Static struct {
	Home    *views.View
	Contact *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewViews("bootstrap", "static/home"),
		Contact: views.NewViews("bootstrap", "static/contact"),
	}
}
