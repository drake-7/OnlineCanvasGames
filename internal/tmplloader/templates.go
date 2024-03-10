package tmplloader

import "text/template"

var Templates = template.Must(template.ParseGlob("web/templates/**/*.html"))
