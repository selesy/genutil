package genutil

import (
	"fmt"
	"go/ast"
	"strings"
)

type Directives map[string][]string

func (d Directives) Add(key string, vals ...string) {
	stored, ok := d[key]
	if !ok {
		d[key] = vals
		return
	}
	d[key] = append(stored, vals...)
}

func (d Directives) merge(directives Directives) {
	for key, vals := range directives {
		d.Add(key, vals...)
	}
}

func (d Directives) Merge(directives ...Directives) {
	for _, e := range directives {
		if e == nil {
			continue
		}
		d.merge(e)
	}
}

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

func CommentGroup(node *ast.GenDecl, prefix string, cfg Config) (Directives, error) {
	d := Directives(make(map[string][]string))
	if node.Doc == nil {
		return d, nil
	}
	for _, c := range node.Doc.List {
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
		vals := strings.Split(val, cfg.ValueSeparator)
		d.Add(key, vals...)
	}
	return d, nil
}

func CommentGroupWithDefaultConfig(node *ast.GenDecl, prefix string) (Directives, error) {
	return CommentGroup(node, prefix, DefaultConfig)
}

func structFieldTagDirectives(field ast.Field, prefix string) Directives {
	d := Directives(make(map[string][]string))
	for _, tag := range strings.Split(field.Tag.Value, "") {
		idx := strings.Index(tag, ":")
		key, val := tag[:idx], tag[idx:]
		if key == prefix {
			val = strings.Trim(val, "\"")
			for _, v := range strings.Split(val, ",") {
				fmt.Println("v: ", v)
			}
			break
		}
	}
	return d
}

func StructFieldTags(node *ast.StructType, prefix string, cfg Config) (Directives, error) {
	d := Directives(make(map[string][]string))
	if node.Fields == nil {
		return d, nil
	}
	// for _, f := range node.FieldList.List {
	// 	t := f.V
	// }
	return d, nil
}

func StructFieldTagsWithDefaultConfig(node *ast.StructType, prefix string) (Directives, error) {
	return StructFieldTags(node, prefix, DefaultConfig)
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
