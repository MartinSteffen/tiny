package absynt  // ast2 might not work
import ("fmt")

type Visitor interface {
	VisitStmt (Stmt)             (Visitor)
	VisitExpr (Expr)             (Visitor)	
	VisitSimpleExpr (SimpleExpr) (Visitor)	
	VisitIdent (Ident)           (Visitor)	
	VisitCompareOp (CompareOp)   (Visitor)
	VisitTerm (Term)             (Visitor)
	VisitFactor (Factor)         (Visitor)
	VisitNumber (Number)         (Visitor)
	VisitSymbol (Symbol)         (Visitor)
	VisitAddOp  (AddOp)          (Visitor)
	VisitMulOp  (MulOp)          (Visitor)
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
	v.VisitSimpleExpr(se)
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


func WalkMulOp(v Visitor, mo MulOp) {
	fmt.Println("MulOp(")
	v.VisitMulOp(mo)
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


func WalkNumber (v Visitor, n Number) {
	fmt.Println("Number(")
	v.VisitNumber (n)
	fmt.Println(")")		
}


func WalkFactor(v Visitor, f Factor) { // fill
	fmt.Println("Term(")           // action
	v.VisitFactor(f)
	switch tf := f.(type) {
	case *ID:
		WalkIdent(v,tf.I)
	case *EXPR:
		WalkExpr(v,tf.E)
	case *NUMBER:
		WalkNumber(v,tf.N)
	}
	fmt.Println(")")
}



func WalkTerm(v Visitor, t Term) {
	fmt.Println("Term(")
	v.VisitTerm(t)                // action
	switch tt := t.(type) {
	case *FACTOR:
		WalkFactor(v,tt.F)
	case *MULEXPR:
		WalkMulOp(v,tt.MO)
		WalkTerm(v,tt.T)
		WalkFactor(v,tt.F)
	}
	fmt.Println(")")
}


func WalkSymbol(v Visitor, s Symbol) {
	fmt.Println("Symbol(")
	v.VisitSymbol(s)
	fmt.Println(")")
}



