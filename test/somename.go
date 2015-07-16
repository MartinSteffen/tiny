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



// just to see that it compiles

func main () {
	f :=    &ast.SIMPLEEXPR{&ast.TERM{&ast.FACTOR{&ast.NUMBER{1}}}}
	fmt.Println(f)
	
	
}



// The following is not possible. One cannot define new methods on
// non-local type ast.FACTOR:

// func ( *ast.FACTOR ) print_factor () {}



// the following is intended to make a ``visitor''. It seems that the
// ``external methods'' of go should be able to do that.


// func main () {
// 	e :=    &Ident{}  // possible, but not nice
// 	e =     &Ident{"a"}        // possible
// 	e1 :=     &Ident{Name:"a"}   // also possible
// 	s1 :=     &SkipStmt{}
// 	s2 :=     &IfStmt{COND:e1,IF:s1,ELSE:s1}   // also possible
// 	fmt.Println(e)
// 	fmt.Println(e1)
// 	fmt.Println(s1)
// 	fmt.Println(s2)
//	}



