// This is to test what ``local'' and ``non-local'' means wrt. methods.

package ast

import ("fmt") // can't hurt



// Unlike in the package "test" which is in a different directory, here we
// can leave out to import "ast" (obviously) and also we don't need to write
// ast.SIMPLEEXPR. etc

func main () {
	f :=    &SIMPLEEXPR{&TERM{&FACTOR{&NUMBER{1}}}}
	fmt.Println(f)
	
	
}

// It seems that inside the same package, one can ``attach'' methods to
// types defined in a different file. Across packages that does not work.

//func (Stmt) print_stmt () {
//	fmt.Println("stmt")
//}

func ( *FACTOR ) print_factor () {fmt.Println("factor")}


