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
// function. 


func Print_Symbol(s Symbol) {fmt.Println(s)}

func Print_Ident(i Ident) {fmt.Println(i)}  // unclear why I can't use Print_Symbol

func Print_Program (p Program) {
	fmt.Println("Program( ");
	Print_Stmt (p)            // no constructor here Program  = Stmt (currently)
	fmt.Println(" )");
}

func  Print_Stmt (s Stmt) {  // *Stmt won't work, the * must be at the constructors
 	 switch sw := s.(type) {
	 case *IF:
		 fmt.Println("IF" )
		 Print_Expr(sw.E)
		 fmt.Println(" THEN " )
		 Print_Stmt(sw.SL1)
		 fmt.Println(" ELSE  " )
		 Print_Stmt(sw.SL2)
	 case *READ:
		 fmt.Println("READ" )
		 Print_Ident(sw.I)
		 fmt.Println(" " )		 
	 case *WRITE:
		 fmt.Println("WRITE" )
		 Print_Expr(sw.E)
		 fmt.Println(" " )
	 case *REPEAT:
		 fmt.Println("REPEAT" )
		 Print_Stmt(sw.SL)
		 Print_Expr(sw.C)
		 fmt.Println(" " )
	 case *ASSIGN:
		 Print_Ident(sw.I)
		 fmt.Println(" := " )		 				 
		 Print_Expr(sw.E)
		 fmt.Println(" " )		 				 
	 default:  	 fmt.Println("stmt")
	 }
 }

func Print_Expr (e Expr) {
	switch expr := e.(type) {
	case *SIMPLEEXPR:
		Print_SimpleExpr(expr.S)
	case *COMPAREEXPR:
		Print_SimpleExpr(expr.SE1)		
		Print_Compare_Op(expr.CO)
		Print_SimpleExpr(expr.SE2)				
	}
}

func Print_SimpleExpr(s SimpleExpr) {
	switch  se := s.(type) {
	case *TERM:
		Print_Term(se.T)  // error
	case *ADDEXPR:
		Print_Add_Op(se.O)
		Print_SimpleExpr(se.SE)
		Print_Term(se.T)  // error
	}
	
}



func Print_Compare_Op(o Compare_Op) {
	switch o.(type) {
	case *LT:
		fmt.Println("<")
		
	case *EQ:
		fmt.Println("=")
	}
	
}



func Print_Term (t Term) {
	switch te := t.(type) {
	case *FACTOR:
		Print_Factor(te.F)
	}
}


func Print_Add_Op (op Add_Op) {
}

func Print_Factor(f Factor) {
	switch fc := f.(type) {
	case *ID:
		Print_Ident(fc.I)
	case *EXPR:
		Print_Expr(fc.E)
	case *NUMBER:
		Print_Number(fc.N)
	}
}


func Print_Number (n Number) {
	fmt.Println(n)
}



// It seems that inside the same package, one can ``attach'' methods to
// types defined in a different file. Across packages that does not work.


// func ( *FACTOR ) print_factor () {fmt.Println("factor")}


