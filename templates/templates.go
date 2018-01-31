package loltemplates

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

import "html/template"

var Templates *template.Template

func init() {
	Templates = template.New("base")

	for _, n := range AssetNames() {
		Templates = template.Must(Templates.Parse(string(MustAsset(n))))
	}
}
