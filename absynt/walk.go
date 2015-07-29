

package absynt  // ast2 might not work

import ("fmt")


// This one is modelled after the visitor in the go-ast. That one has
// already the property I wanted, that there are ``walkers'' specific for
// various sub-categories, not a flat, generic case-switch.  Later, one may
// make the various functions internal, for the time being, I export  all
// functions.

// There current design has the following problem. The definition of
// statements and expressions are mutually recursive. That implies that
// also the walk functions are mutually recursive (where by the different
// walk functions I mean the walker for statements and the walker for
// expression, and perhaps more). Currently, with the walkers as functions,
// the different walkers are distinguished by name (if that's necessary
// also if the walkers are methods is unclear, also if that would be
// better). Anyhow that's not the current problem. The problem is the
// recursive call in a walk function to other walk functions, where a/the
// visitor has to be passed as argument.  Now, there are _different_
// visitors of course, and they are an argument of the walk function. It
// would be great if one could avoid handing over all visitors are
// arguments, but instead make use of polymorphism or something. To do
// that, I make a Visitor interfaces for Node (which I make a supertype).

type Visitor interface {
	VisitStmt (Stmt)           (Visitor)
	VisitExpr (Expr)           (Visitor)	
	VisitIdent (Ident)         (Visitor)	
	VisitCompareOp (CompareOp) (Visitor)
	VisitTerm (Term)           (Visitor)
	VisitFactor (Factor)       (Visitor)
	VisitNumber (Number)      (Visitor)
	VisitAddOp (AddOp)        (Visitor)
}


// it seems that the following are just synonyms then.
//type StmtVisitor interface {
//	Visitor
//}
//
//
//type ExprVisitor interface {
//	Visitor
//}	
// type StmtVisitor interface {
// 	VisitStmt(Stmt) (StmtVisitor)   // can I ``overload'' Visit?
// }
// type ExprVisitor interface {
// 	Visit(Expr) (ExprVisitor)
// 	// dummy
// }


//type SimpleExprVisitor interface {
//	// dummy
//	
//}
// Note: the Walk-function is _not_ supposed to adhere to the Visitor
// interface, therefore it does not return a Visitor, it's a
// side-effect-only function. If one wanted a return type, one would have
// to write for instance as func Walk (v int, p Program) (int) {....


//func Walk (v Visitor, p Program) {
//}



func WalkStmt (v Visitor, s Stmt) {
	// we might want to check of the visitor is empty.  or perhaps it
	// should be one at the call site. We leave it out for the moment.
	v.VisitStmt(s)          // action 
	switch ts := s.(type) { // type assertion
	case *IF:
		fmt.Println ("IF (")
		WalkExpr (v, ts.E)   
		WalkStmt (v, ts.SL1)
		WalkStmt (v, ts.SL2)
		fmt.Println (")")
	case *READ:
		fmt.Println ("Read (")
		WalkIdent (v, ts.I)
		fmt.Println (")")
	case *WRITE:
		fmt.Println ("Write(")
		WalkExpr (v, ts.E)
		fmt.Println (")")
	case *REPEAT:
		fmt.Println ("Repeat(")
		WalkStmt(v,ts.SL)
		WalkExpr(v,ts.C)
		fmt.Println (")")
	case *ASSIGN:
		fmt.Println("Assign(")
		WalkIdent(v,ts.I)
		WalkExpr(v,ts.E)
		fmt.Println(")")
	}
}
func WalkExpr (v Visitor, e Expr) {
	v.VisitExpr(e)          // action 
	fmt.Println("Expr(")	
	switch te := e.(type) { // type assertion
	case *SIMPLEEXPR:
		WalkSimpleExpr(v,te.S)
	case *COMPAREEXPR:
		WalkCompareOp(v, te.CO)
		WalkSimpleExpr(v,te.SE1)
		WalkSimpleExpr(v,te.SE2)
	}
	fmt.Println(")")		
}

func WalkSimpleExpr(v Visitor, se SimpleExpr) {
	fmt.Println("SimpleExpr(")
	switch tse :=se.(type) {
	case *TERM:
		WalkTerm(v,tse.T)
        case *ADDEXPR:
		WalkAddOp(v,tse.O)
		WalkSimpleExpr(v,tse.SE)
		WalkTerm(v,tse.T)
	}
	fmt.Println(")")
}



func WalkAddOp(v Visitor, ao AddOp) {
	fmt.Println("AddOp(")
	v.VisitAddOp(ao)
	fmt.Println(")")	
}


func WalkCompareOp(v Visitor, co CompareOp) {
	fmt.Println("CompareOp(")
	v.VisitCompareOp (co)
	fmt.Println(")")	
	
}

func WalkIdent (v Visitor, i Ident) {
	fmt.Println("Ident(")
	v.VisitIdent (i)
	fmt.Println(")")	
	
}


func WalkTerm(v Visitor, t Term) {
	fmt.Println("Term(")
	v.VisitTerm(t)
	fmt.Println(")")
	
}



