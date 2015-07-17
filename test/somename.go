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


//func idents(f *ast.File) <-chan *ast.Ident {
//	v := make(visitor)
//	go func() {
//		ast.Walk(v, f)
//		close(v)
//	}()
//	return v





//type visitor int // orginally chan *ast.Ident In the example, the visitor
// type is a channel (here a channel containing references to statements)


type visitor chan *ast.Stmt



// The following actually defines the function being iterated, I assume.


func (v visitor) Visit(n ast.Stmt) (w ast.Visitor_Stmt) {
	return v
}


func main () {
	e1 :=    &ast.SIMPLEEXPR{&ast.TERM{&ast.FACTOR{&ast.NUMBER{1}}}}
	s1 :=    &ast.ASSIGN{"x",e1}
	ast.Print_Expr(e1)
	ast.Print_Stmt(s1)
	fmt.Println("---")
	v:=make(visitor)

	
	
}







