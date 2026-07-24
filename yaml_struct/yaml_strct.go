package yaml_struct

import "html/template"

type PageMeta struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
}

type PageData struct {
	Title string
	Date  string
	Body  template.HTML
}
