

package absynt  // ast2 might not work 


// This one is modelled after the visitor in the go-ast. That one has
// already the property I wanted, that there are ``walkers'' specific for
// various sub-categories, not a flat, generic case-switch.  Later, one may
// make the various functions internal, for the time being, I export  all
// functions.

type Visitor interface {
	Visit(p Program) (w Visitor)
}

func Walk (v Visitor, p Program) {


}



