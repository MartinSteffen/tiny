package walkfunctional  // ast2 might not work

//import ("github.com/MartinSteffen/tiny/absynt")
import ("absynt")


type Visitor interface {
	VisitStmt (absynt.Stmt)             (Visitor)
	VisitExpr (absynt.Expr)             (Visitor)	
	VisitSimpleExpr (absynt.SimpleExpr) (Visitor)	
	VisitIdent (absynt.Ident)           (Visitor)	
	VisitCompareOp (absynt.CompareOp)   (Visitor)
	VisitTerm (absynt.Term)             (Visitor)
	VisitFactor (absynt.Factor)         (Visitor)
	VisitNumber (absynt.Number)         (Visitor)
	VisitSymbol (absynt.Symbol)         (Visitor)
	VisitAddOp  (absynt.AddOp)          (Visitor)
	VisitMulOp  (absynt.MulOp)          (Visitor)
}


func WalkStmt (v Visitor, s absynt.Stmt)  {
	// we might want to check of the visitor is empty.  or perhaps it
	// should be one at the call site. We leave it out for the moment.
	v=v.VisitStmt(s)          // action 
	switch ts := s.(type) { // type assertion
	case *absynt.IF:
		WalkExpr (v, ts.E)   
		WalkStmt (v, ts.SL1)
		WalkStmt (v, ts.SL2)
	case *absynt.READ:
		WalkIdent (v, ts.I)
	case *absynt.WRITE:
		WalkExpr (v, ts.E)
	case *absynt.REPEAT:
		WalkStmt(v,ts.SL)
		WalkExpr(v,ts.C)
	case *absynt.ASSIGN:
		WalkIdent(v,ts.I)
		WalkExpr(v,ts.E)
	}
}
func WalkExpr (v Visitor, e absynt.Expr) {
	v=v.VisitExpr(e)          // action 
	switch te := e.(type) { // type assertion
	case *absynt.SIMPLEEXPR:
		WalkSimpleExpr(v,te.S)
	case *absynt.COMPAREEXPR:
		WalkCompareOp(v, te.CO)
		WalkSimpleExpr(v,te.SE1)
		WalkSimpleExpr(v,te.SE2)
	}
}

func WalkSimpleExpr(v Visitor, se absynt.SimpleExpr) {
	v.VisitSimpleExpr(se)
	switch tse :=se.(type) {
	case *absynt.TERM:
		WalkTerm(v,tse.T)
        case *absynt.ADDEXPR:
		WalkAddOp(v,tse.O)
		WalkSimpleExpr(v,tse.SE)
		WalkTerm(v,tse.T)
	}
}



func WalkAddOp(v Visitor, ao absynt.AddOp) {
	v.VisitAddOp(ao)
}


func WalkMulOp(v Visitor, mo absynt.MulOp) {
	v.VisitMulOp(mo)
}


func WalkCompareOp(v Visitor, co absynt.CompareOp) {
	v.VisitCompareOp (co)
}

func WalkIdent (v Visitor, i absynt.Ident) {
	v.VisitIdent (i)
}


func WalkNumber (v Visitor, n absynt.Number) {
	v.VisitNumber (n)
}


func WalkFactor(v Visitor, f absynt.Factor) {
	v.VisitFactor(f)
	switch tf := f.(type) {
	case *absynt.ID:
		WalkIdent(v,tf.I)
	case *absynt.EXPR:
		WalkExpr(v,tf.E)
	case *absynt.NUMBER:
		WalkNumber(v,tf.N)
	}
}



func WalkTerm(v Visitor, t absynt.Term) {
	v.VisitTerm(t)                // action
	switch tt := t.(type) {
	case *absynt.FACTOR:
		WalkFactor(v,tt.F)
	case *absynt.MULEXPR:
		WalkMulOp(v,tt.MO)
		WalkTerm(v,tt.T)
		WalkFactor(v,tt.F)
	}
}


func WalkSymbol(v Visitor, s absynt.Symbol) {
	v.VisitSymbol(s)
}



