// Let's try the same with a visitor pattern. We should afterwards test, if
// one can have client code in a different package. 

package ast

// The first thing to have is the interface. This is modelled after the one
// from the go language. 


// Next they had 3 functions (not methods) which they called "helper
// functions. It seems they are iterators for "lists", i.e., slices.  Let
// leave them out for a moment, resp. see what we will need in our setting
// later.

// The main function in that file is "Walk", which is the visitor
// functionality. The visitor is invoked on a node of type "Node" which is
// the catch-all interface for all nodes in the ast there. Inside the
// visitor, there's a flat switch over all.

// The ``visitor'' pattern is basically a higher-order function. It applies
// the function argument to all nodes of the structure. In the go
// implementation, the way it's done is flat. This is specified via the
// Visitor interface, which requires the Visit-function. In a way, it's not
// so clear if it's a function or method, because the Visit is used in the
// code in the form v.Visit(n), where n is a node.  One problem we will get
// is the non-flat structure. Below, there's the visitor, which visits a
// statement. That's fine as it is, but we also want to visit other
// stuff. Perhaps we need to introduce super-interfaces or stuff.





type Visitor_Stmt interface {
	Visit (stmt Stmt) (w Visitor_Stmt)
}



func Walk_Stmt (v Visitor_Stmt, s Stmt) {
	// we check for non-nil-ness of the visitor (not the statement). If
	// the statement is nil, probably not good either. 
	if v =v.Visit(s); v == nil {
		return
	}
}
