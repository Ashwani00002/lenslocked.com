package views

import (
	"html/template"
	"path/filepath"
)

var (
	LayoutDir   = "views/layout/"
	TemplateExt = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		return nil
	}
	return files
}

func NewViews(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
