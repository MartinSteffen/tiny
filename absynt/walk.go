package absynt  // ast2 might not work


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
		WalkExpr (v, ts.E)   
		WalkStmt (v, ts.SL1)
		WalkStmt (v, ts.SL2)
	case *READ:
		WalkIdent (v, ts.I)
	case *WRITE:
		WalkExpr (v, ts.E)
	case *REPEAT:
		WalkStmt(v,ts.SL)
		WalkExpr(v,ts.C)
	case *ASSIGN:
		WalkIdent(v,ts.I)
		WalkExpr(v,ts.E)
	}
}
func WalkExpr (v Visitor, e Expr) {
	v.VisitExpr(e)          // action 
	switch te := e.(type) { // type assertion
	case *SIMPLEEXPR:
		WalkSimpleExpr(v,te.S)
	case *COMPAREEXPR:
		WalkCompareOp(v, te.CO)
		WalkSimpleExpr(v,te.SE1)
		WalkSimpleExpr(v,te.SE2)
	}
}

func WalkSimpleExpr(v Visitor, se SimpleExpr) {
	v.VisitSimpleExpr(se)
	switch tse :=se.(type) {
	case *TERM:
		WalkTerm(v,tse.T)
        case *ADDEXPR:
		WalkAddOp(v,tse.O)
		WalkSimpleExpr(v,tse.SE)
		WalkTerm(v,tse.T)
	}
}



func WalkAddOp(v Visitor, ao AddOp) {
	v.VisitAddOp(ao)
}


func WalkMulOp(v Visitor, mo MulOp) {
	v.VisitMulOp(mo)
}


func WalkCompareOp(v Visitor, co CompareOp) {
	v.VisitCompareOp (co)
}

func WalkIdent (v Visitor, i Ident) {
	v.VisitIdent (i)
}


func WalkNumber (v Visitor, n Number) {
	v.VisitNumber (n)
}


func WalkFactor(v Visitor, f Factor) { // fill
	v.VisitFactor(f)
	switch tf := f.(type) {
	case *ID:
		WalkIdent(v,tf.I)
	case *EXPR:
		WalkExpr(v,tf.E)
	case *NUMBER:
		WalkNumber(v,tf.N)
	}
}



func WalkTerm(v Visitor, t Term) {
	v.VisitTerm(t)                // action
	switch tt := t.(type) {
	case *FACTOR:
		WalkFactor(v,tt.F)
	case *MULEXPR:
		WalkMulOp(v,tt.MO)
		WalkTerm(v,tt.T)
		WalkFactor(v,tt.F)
	}
}


func WalkSymbol(v Visitor, s Symbol) {
	v.VisitSymbol(s)
}



