package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   = "views/layout/"
	TemplateDir = "views/"
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

	addTemplatePath(files)
	addTemplateExt(files)

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

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

// Render is used to render the view with pre-defined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// addTemplatePath takes slice of strings representing file paths for templates, and it
// prepends the TemplateDir directory to each string in the slice.
//
// Eg the input {"home"} would result in the output {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes slice of strings representing file paths for templates, and it
// appends the TemplateExt to each string in the slice.

// Eg the input {"home"} would result in the output {"home.gohtml"} if TemplateExt == ".gohtml"

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
