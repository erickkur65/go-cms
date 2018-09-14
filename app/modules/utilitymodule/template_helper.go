package utilitymodule

import (
	"fmt"
	"html/template"
	"os"
)

// TemplateHelper helper to handle template
type TemplateHelper struct{}

// PopulateTemplates for populate all html templates
func (templateHelper *TemplateHelper) PopulateTemplates() *template.Template {
	templates, err := template.ParseGlob("../templates/*.html")

	if err != nil {
		fmt.Println(err)
		fmt.Println("error when parse templates")
		os.Exit(1)
	}

	return templates
}
