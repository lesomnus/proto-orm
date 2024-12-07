package pbgen

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"strings"
	"text/template"
)

var (
	//go:embed *.tpl
	template_files embed.FS
	proto_template *template.Template
)

type templateBond interface {
	TemplateName() string
}

func Execute(w io.Writer, def templateBond) error {
	return proto_template.ExecuteTemplate(w, def.TemplateName()+".go.tpl", def)
}

func init() {
	proto_template = template.New("")
	proto_template.Funcs(template.FuncMap{
		"split": func(sep, v string) []string {
			return strings.Split(v, sep)
		},
		"indent": func(n int, v string) string {
			pad := strings.Repeat("\t", n)
			return strings.Replace(v, "\n", "\n"+pad, -1)
		},
		"include": func(name string, data interface{}) (string, error) {
			b := bytes.Buffer{}
			if err := proto_template.ExecuteTemplate(&b, fmt.Sprintf("%s.go.tpl", name), data); err != nil {
				return "", err
			}
			return b.String(), nil
		},
	})

	t, err := proto_template.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	proto_template = t
}
