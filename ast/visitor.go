// Let's try the same with a visitor pattern. We should afterwards test, if
// one can have client code in a different package. 

package ast

// The first thing to have is the interface
type Visitor interface {
	Visit (node Node) (w Visitor)
}


