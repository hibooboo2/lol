package ddragon

import (
	"net/http"

	loltemplates "github.com/hibooboo2/lol/templates"
)

func (i *Item) Html(w http.ResponseWriter) {
	loltemplates.Templates.ExecuteTemplate(w, "item", i)
}
