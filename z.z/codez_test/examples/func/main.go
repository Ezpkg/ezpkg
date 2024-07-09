package main

import (
	"fmt"

	"ezpkg.io/codez"
	"ezpkg.io/colorz"
	"ezpkg.io/errorz"
)

func main() {
	pkgs := errorz.Must(codez.LoadPackages("ezpkg.io/-/codez_test/testdata/logging/main"))
	fmt.Println(colorz.Blue.Wrap("👉 loaded packages:"))
	for _, pkg := range pkgs.Packages() {
		fmt.Printf("\t%v\n", pkg.PkgPath)
	}
	fmt.Println()
	fmt.Println(colorz.Blue.Wrap("👉 all packages:"))
	for _, pkg := range pkgs.AllPackages() {
		fmt.Printf("\t%v\n", pkg.PkgPath)
	}
	fmt.Println()
	fmt.Println(colorz.Blue.Wrap("👉 filter ezpkg.io/... , golang.org/..."))
	for _, pkg := range pkgs.AllPackages("ezpkg.io/...", "golang.org/...") {
		fmt.Printf("\t%v\n", pkg.PkgPath)
	}
}
