// One can compile this file, despite the fact that the name of the file,
// the name of the directory and the name of the package are all different.
// One can do it with go build somename.go (also install and run), but one
// cannot do it just with go install (without saying what).  There, the
// package is called main, but the path is part of the package name.  Note:
// if one compiles this file with go build, it gives the file ./text (not
// somename)





package main


import ("fmt"
	"github.com/MartinSteffen/tiny/absynt")


// If we want to make use if the visitor, which is basically a higher-order
// function, we need to implement the visitor interface.


//func idents(f *ast.File) <-chan *ast.Ident {
//	v := make(visitor)
//	go func() {
//		ast.Walk(v, f)
//		close(v)
//	}()
//	return v





// Originally chan *ast.Ident In the example, the visitor type is a channel
// (here a channel containing references to statements). Later below, we
// have dot


//type visitor chan *ast.Stmt  // it's not so clear of that's a confusing
			     // name



// The following actually defines the function being iterated, I assume.
// Visit: visitor .-> Stmt -> Visitor_Stmt


//func (v visitor) Visit(n ast.Stmt) (w ast.Visitor_Stmt) {
//	return v
//}

func main () {
	f := &absynt.ID{I:"s"}
	var t = &absynt.FACTOR{f}   // term
	var se = &absynt.TERM{t}    // simple expr
	var e = &absynt.SIMPLEEXPR{se}  // expr 
	var sr = &absynt.READ{I:"x"}    // read stmt
	var sl1 = sr             // stmt (list) 
	var sl2 = sl1            // stmt (list) 
	var s = &absynt.IF{e,sl1,sl2}   // stmt 
	fmt.Println ()
	v := make(absynt.visitor)

}


//func main () {
//	e1 :=    &ast.SIMPLEEXPR{&ast.TERM{&ast.FACTOR{&ast.NUMBER{1}}}}
//	s1 :=    &ast.ASSIGN{"x",e1}
//	ast.Print_Expr(e1)
//	ast.Print_Stmt(s1)
//	fmt.Println("---")
//	v:=make(visitor)
//	ast.Walk_Stmt(v,f)
//}







