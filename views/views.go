package views

import "html/template"

type View struct {
	Template *template.Template
	Layout   string
}

func NewViews(layout string, files ...string) *View {
	files = append(files, "views/layout/bootstrap.gohtml",
		"views/layout/footer.gohtml",
		"views/layout/navbar.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}
