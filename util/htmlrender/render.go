package htmlrender

import (
	"html/template"
	"path"

	. "videoporter/templates"

	"github.com/gin-gonic/gin/render"
)

type Render struct {
	Layout       string
	TemplatesDir string
}

func (r *Render) Instance(name string, data interface{}) render.Render {
	tpl := template.New(name)
	content, err := Asset(path.Join(path.Join(r.TemplatesDir, name)))
	if err != nil {
		panic(err)
	}
	tpl = template.Must(tpl.Parse(string(content)))
	return render.HTML{
		Template: tpl,
		Data:     data,
	}
}

func New() *Render {
	return &Render{
		Layout:       "",
		TemplatesDir: "templates",
	}
}
