package unitemplate

import "github.com/robfig/revel"

type Controller struct {
	*revel.Controller
	currentLayout string
	noLayout      bool
}

func (c Controller) Render(extraRenderArgs ...interface{}) revel.Result {

}

func (c Controller) RenderTemplate(name string) revel.Result {

}

func (c Controller) RenderTemplateWithLayout(name, layout string) revel.Result {

}

func (c Controller) ContentFor(name, templateName string) error {

}

func (c Controller) Layout(layout string) error {

}
