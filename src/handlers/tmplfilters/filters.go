package tmplfilters

import (
	"github.com/flosch/pongo2"
	"github.com/uwine4850/foozy/pkg/router/tmlengine"
)

var filters = []tmlengine.Filter{
	{
		Name: "imgOrDef",
		Fn: func(in, param *pongo2.Value) (out *pongo2.Value, err *pongo2.Error) {
			var imagePath string
			inputPath := in.String()
			if inputPath != "" {
				imagePath = inputPath
			} else {
				imagePath = "/static/img/default/default.jpg"
			}
			return pongo2.AsValue(imagePath), nil
		},
	},
}

func RegisterFilters() {
	tmlengine.RegisterMultipleGlobalFilter(filters)
}
