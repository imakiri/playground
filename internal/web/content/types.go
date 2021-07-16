package content

import "html/template"

type Content struct {
	Index  *template.Template
	Static struct {
		Ico []byte
		Css []byte
	}
	Gorum struct {
		Index *template.Template
	}
}
