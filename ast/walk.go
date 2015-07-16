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


// Originally I thought perhaps one could have an ``easy'' visitor, at
// least inside one package. However, the following is not possible:
// func (s *Stmt) print_stmt () {
//	fmt.Println("stmt") .....
// }


// The problem is: Stmt is an interface, and that's not allowed as receiver
// for a method. Thefore, one cannod hand a statement over to that function
// and then make a big case distinction.  What does work instead is a
// function:


func  Print_Stmt (s Stmt) {  // *Stmt won't work, the * must be at the constructors
 	 switch sw := s.(type) {
	 case *IF:
		 Print_Expr(sw.E)
		 fmt.Println(sw)
	 default:  	 fmt.Println("stmt")
	 }
 }

func Print_Expr (e Expr) {}




// It seems that inside the same package, one can ``attach'' methods to
// types defined in a different file. Across packages that does not work.


func ( *FACTOR ) print_factor () {fmt.Println("factor")}


