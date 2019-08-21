package genutil

import "text/template"

func LoadTemplates(template *template.Template, templates ...string) error {
	for _, t := range templates {
		_, err := template.Parse(t)
		if err != nil {
			return err
		}
	}
	return nil
}
