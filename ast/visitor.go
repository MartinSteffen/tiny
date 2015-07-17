// Let's try the same with a visitor pattern. We should afterwards test, if
// one can have client code in a different package. 

package ast

// The first thing to have is the interface. This is modelled after the one
// from the go language. 
type Visitor interface {
	Visit (node Node) (w Visitor)
}


// Next they had 3 functions (not methods) which they called "helper
// functions. It seems they are iterators for "lists", i.e., slices.  Let
// leave them out for a moment, resp. see what we will need in our setting
// later.

// The main function in that file is "Walk", which is the visitor
// functionality. The visitor is invoked on a node of type "Node" which is
// the catch-all interface for all nodes in the ast there. Inside the
// visitor, there's a flat switch over all 

