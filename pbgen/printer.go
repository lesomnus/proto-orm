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
)

type Printer struct {
	t *template.Template
}

func NewPrinter(pkg string) *Printer {
	t := newTemplate(pkg)

	return &Printer{t}
}

type templateBond interface {
	TemplateName() string
}

func PrintFile(o io.Writer, f *ProtoFile) error {
	p := NewPrinter(string(f.Package))
	return p.Print(o, f)
}

func (p *Printer) Print(w io.Writer, def templateBond) error {
	return p.t.ExecuteTemplate(w, def.TemplateName()+".go.tpl", def)
}

func newTemplate(pkg string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"split": func(sep, v string) []string {
			return strings.Split(v, sep)
		},
		"indent": func(n int, v string) string {
			pad := strings.Repeat("\t", n)
			return strings.Replace(v, "\n", "\n"+pad, -1)
		},
		"include": func(name string, data interface{}) (string, error) {
			b := bytes.Buffer{}
			if err := t.ExecuteTemplate(&b, fmt.Sprintf("%s.go.tpl", name), data); err != nil {
				return "", err
			}
			return b.String(), nil
		},
		"type": func(v Type) string {
			r, ok := strings.CutPrefix(string(v), pkg)
			if !ok {
				return string(v)
			}
			if r[0] != '.' || len(r) == 0 {
				return string(v)
			}
			return r[1:]
		},
	})
	if _, err := t.ParseFS(template_files, "*"); err != nil {
		panic(err)
	}

	return t
}
