package tmplloader

import "html/template"

var Templates = template.Must(template.ParseGlob("web/templates/**/*.html"))
