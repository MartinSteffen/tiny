// This is to test what ``local'' and ``non-local'' means wrt. methods.

package ast

import ("fmt") // can't hurt



// Unlike in the package "test" which is in a different directory, here we
// can leave out to import "ast" (obvourly) and also we don't need to write
// ast.SIMPLEEXPR. etc

func main () {
	f :=    &SIMPLEEXPR{&TERM{&FACTOR{&NUMBER{1}}}}
	fmt.Println(f)
	
	
}


func ( *FACTOR ) print_factor () {}


