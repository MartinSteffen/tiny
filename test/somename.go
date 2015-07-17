// One can compile this file, despite the fact that the name of the file,
// the name of the directory and the name of the package are all different.
// One can do it with go build somename.go (also install and run), but one
// cannot do it just with go install (without saying what).  There, the
// package is called main, but the path is part of the package name.  Note:
// if one compiles this file with go build, it gives the file ./text (not
// somename)



package main


import ("fmt"
	"github.com/MartinSteffen/tiny/ast")


// If we want to make use if the visitor, which is basically a higher-order
// function, we need to implement the visitor interface.



func main () {
	e1 :=    &ast.SIMPLEEXPR{&ast.TERM{&ast.FACTOR{&ast.NUMBER{1}}}}
	s1 :=    &ast.ASSIGN{"x",e1}
	ast.Print_Expr(e1)
	ast.Print_Stmt(s1)
	fmt.Println(";")
	
	
}







