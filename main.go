package main

import (
	"flag"
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Parse()

	cfg := &packages.Config{
		Mode:  packages.LoadSyntax,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, flag.Arg(0))
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		// obj := pkg.Types.Scope().Lookup(flag.Arg(1))
		// fmt.Fprintf(os.Stderr, "%#v\n", obj)

		// if obj != nil {
		// 	pos := obj.Pos()
		// 	file := pkg.Fset.File(pos)
		// 	position := pkg.Fset.PositionFor(pos, false)

		// 	fmt.Fprintf(os.Stderr, "pos: %#v\n", pos)
		// 	fmt.Fprintf(os.Stderr, "position: %#v\n", position)
		// 	fmt.Fprintf(os.Stderr, "file.Base(): %#v\n", file.Base())
		// 	// fmt.Fprintf(os.Stderr, "file: %#v\n", file)

		// 	fmt.Printf("%s:#%d", position.Filename, int(pos)-file.Base())

		// 	break
		// }
		// fmt.Println(pkg.Name, pkg.PkgPath)
		for _, file := range pkg.Syntax {
			// fmt.Println("\tfile:", file.Name.String())
			obj := file.Scope.Lookup(flag.Arg(1))
			if obj != nil {
				position := pkg.Fset.Position(obj.Pos())
				fmt.Printf("%s:#%d\n", position.Filename, position.Offset)
				break
			}
		}
	}
}
