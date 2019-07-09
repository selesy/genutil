package genutil

import (
	"go/ast"
	"strings"
)

type Directives map[string][]string

type Config struct {
	KeywordSeparator    string
	KeyValueSeparator   string
	TrimValueWhitespace bool
	ValueSeparator      string
}

var DefaultConfig = Config{
	KeywordSeparator:    "-",
	KeyValueSeparator:   ":",
	TrimValueWhitespace: true,
	ValueSeparator:      ",",
}

func CommentGroup(decl *ast.GenDecl, prefix string, cfg Config) (Directives, error) {
	d := Directives(make(map[string][]string))
	if decl.Doc == nil {
		return d, nil
	}
	for _, c := range decl.Doc.List {
		t := c.Text[2:] // Skip over // or /*
		if !strings.HasPrefix(t, prefix) {
			continue
		}
		idx := strings.Index(t, cfg.KeyValueSeparator)
		key, val := t[len(prefix):idx], t[idx:]
		key = strings.ReplaceAll(key, cfg.KeywordSeparator, "")
		if cfg.TrimValueWhitespace {
			val = strings.Trim(val, " ")
		}
		vals, ok := d[key]
		if !ok {
			vals = *new([]string)
		}
		vals = append(vals, strings.Split(val, cfg.ValueSeparator)...)
		d[key] = vals
	}
	return d, nil
}

func CommentGroupWithDefaultConfig(node *ast.GenDecl, prefix string) (Directives, error) {
	return CommentGroup(node, prefix, DefaultConfig)
}

func StructFieldTags(node *ast.StructType, prefix string) (Directives, error) {
	return nil, nil
}

// func blah(s *ast.StructType, prefix string) []string {
// 	return nil
// }

func (d Directives) UniqueDirective(node *ast.Node, suffix string) (string, error) {
	var v string
	return v, nil
}

func (d Directives) MultipleDirective(node *ast.Node, suffix string) ([]string, error) {
	var v []string
	return v, nil
}
