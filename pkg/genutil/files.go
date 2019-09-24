package genutil

import (
	"errors"
	"strings"
)

const (
	goSuffix    = ".go"
	goGenSuffix = "_gen.go"
)

//GoFilesFromPkgs returns the list of files that were parsed when loading
//the packages and are therefore included in the FileSet.
func GoFilesFromPkgs(pkgs Pkgs) (files []string) {
	for _, pkg := range pkgs.pkgs {
		files = append(files, pkg.GoFiles...)
	}
	return files
}

func GoFilesFromPatterns(patterns ...string) ([]string, error) {
	pkgs, err := PackagesFromPatterns(patterns...)
	if err != nil {
		return nil, err
	}
	return GoFilesFromPkgs(pkgs), nil
}

func GoFilesFromArgs() ([]string, error) {
	pkgs, err := PackagesFromArgs()
	if err != nil {
		return nil, err
	}
	return GoFilesFromPkgs(pkgs), nil
}

func GeneratedGoFileName(goFile string) (string, error) {
	if !strings.HasSuffix(goFile, goSuffix) {
		return "", errors.New("this method is expecting the name of a .go file")
	}
	return goFile[0:strings.LastIndex(goFile, goSuffix)] + goGenSuffix, nil
}
