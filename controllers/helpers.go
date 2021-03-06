package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	err = schema.NewDecoder().Decode(dst, r.PostForm)
	if err != nil {
		return err
	}
	return nil
}
