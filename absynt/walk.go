package absynt  // ast2 might not work
import ("fmt")

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



