

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
}



