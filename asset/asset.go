package asset

import (
	"github.com/flosch/pongo2"

	"github.com/gobuffalo/packr/v2"
	"github.com/wbsifan/pongo2-packr"
)

var (
	Box *packr.Box
	Tpl *pongo2.TemplateSet
)

func init() {
	Box = packr.New("template", "../template")
	loader := pongo2packr.NewLoader(Box)
	Tpl = pongo2.NewSet("", loader)
}

func FileList(prefix string) ([]string, error) {
	list := make([]string, 0)
	err := Box.WalkPrefix(prefix, func(path string, file packr.File) error {
		list = append(list, path)
		return nil
	})
	return list, err
}
