package genutil

import (
	"flag"
	"go/token"

	"golang.org/x/tools/go/packages"
)

//Pkgs provides a copy of the configuration used to parse the packages
//bound together with the results of loading the specified package search
//patterns.
type Pkgs struct {
	cfg  packages.Config
	pats []string
	pkgs []*packages.Package
}

func (pkgs Pkgs) Config() packages.Config {
	return pkgs.cfg
}

//FileSet provides a short-cut method to the oft-used FileSet object.
func (pkgs Pkgs) FileSet() *token.FileSet {
	return pkgs.cfg.Fset
}

//Packages provides a list of packages that were discovered using the
//specified patterns.
func (pkgs Pkgs) Packages() []*packages.Package {
	return pkgs.pkgs
}

//Patterns returns the list of patterns that were originally provided
//to the package loading process.
func (pkgs Pkgs) Patterns() []string {
	return pkgs.pats
}

//AddPackagesToFileSet updates an existing fileset with the packages found
//by searching the provided package patterns.
func AddPackagesToFileSet(fset *token.FileSet, patterns ...string) (pkgs Pkgs, err error) {
	cfg := packages.Config{
		Mode: packages.NeedSyntax,
		Fset: fset,
	}

	pkgs.pkgs, err = packages.Load(&cfg, patterns...)
	if err != nil {
		return pkgs, err
	}

	pkgs.cfg = cfg
	pkgs.pats = patterns
	return pkgs, nil
}

func PackagesFromPatterns(patterns ...string) (Pkgs, error) {
	fset := token.NewFileSet()
	return AddPackagesToFileSet(fset, patterns...)
}

func PackagesFromArgs() (Pkgs, error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	return PackagesFromPatterns(flag.Args()...)
}
