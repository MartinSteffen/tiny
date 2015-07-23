

package absynt  // ast2 might not work 


// This one is modelled after the visitor in the go-ast. That one has
// already the property I wanted, that there are ``walkers'' specific for
// various sub-categories, not a flat, generic case-switch.  Later, one may
// make the various functions internal, for the time being, I export  all
// functions.

type Visitor interface {
	Visit(Program) (Visitor)
}


type StmtVisitor interface {
	
}
// Note: the Walk-function is _not_ supposed to adhere to the Visitor
// interface, therefore it does not return a Visitor, it's a
// side-effect-only function. If one wanted a return type, one would have
// to write for instance as func Walk (v int, p Program) (int) {....

func Walk (v Visitor, p Program) {
}



func WalkStmt (sv StmtVisitor, s Stmt) {
	// we might want to check of the visitor is empty.  or perhaps it
	// should be one at the call site. We leave it out for the moment.
	switch ts := s.(type) { // type assertion
	case *IF:
		WalkStmt (sv, ts.SL1)
		WalkStmt (sv, ts.SL2)
	case *READ:
	case *WRITE:
	case *REPEAT:
	case *ASSIGN:
	}
}
	// IF struct {E Expr
	// READ struct {I Ident}
	// WRITE struct {E Expr}
	// REPEAT struct {SL Stmt; C Expr}  // slice
	// ASSIGN struct {I Ident; E Expr}

        // }



