// The abstract syntax is on the one hand inspired by the ML version, but
// also of course the C version, where the C version is rather
// unstructured. The final source of inspiration is the go-ast, to learn
// how they did it.


// The first statement here must be the declaration of the package.  We
// cannot name it package "main", because then it needs a function main, it
// seems. If it contains a main and is called package main, then it's
// installed under bin

package ast

//import ("fmt")

// In the ML code, I use information about the position. That's still more
// or less dummy information. I leave it out here. An elegant solution here
// would be to ``embed'' it via a top-level interace. That should be
// relatively easy.



// The following node is currently added to make the visitor compile. That
// ast in the go compiler has a "Node" interface as well (with positioning
// info) which is embedded in the rest of the structures.


type Node interface {

}




// The following types need to be capitalized as they are needed
// externally. Non-capitalized type declarations are ``private''. The same





// holds for fields.

type Symbol string      // might be replaced by something more efficient.
type Ident Symbol
type Number int         
                        
//--------------------------------------------------------
type Compare_Op interface {    // abstract
	compare_opNode ()
}
type (                         // that's just grouping
	LT struct {}
	EQ struct {}
)
func (*LT) compare_opNode() {}
func (*EQ) compare_opNode() {}

//--------------------------------------------------------

type Add_Op interface{
	add_opNode()
}
type (
	PLUS struct {}
	MINUS struct {}
)

func (*PLUS) add_opNode() {}
func (*MINUS) add_opNode() {}


//--------------------------------------------------------

type Mul_Op interface{
	mul_opNode()
}
type (
	TIMES struct {}
	 OVER struct {}
)

func (*TIMES) mul_opNode() {}
func (*OVER)  mul_opNode() {}


type Program    [] Stmt   // this is a slice type.

//--------------------------------------------------------


type Stmt interface {
	stmt_Node ()
}

type (
	IF struct {E Expr
		SL1  [] Stmt  // slice
		SL2  [] Stmt  // slice
	}
	READ struct {I Ident}
	WRITE struct {E Expr}
	REPEAT struct {SL [] Stmt; C Expr}  // slice
	ASSIGN struct {I Ident; E Expr}
)


func (*IF) stmt_Node() {}
func (*READ) stmt_Node() {}
func (*WRITE) stmt_Node() {}
func (*REPEAT) stmt_Node() {}
func (*ASSIGN) stmt_Node() {}

//--------------------------------------------------------

type Expr interface {
	expr_Node ()
}

type (
	SIMPLEEXPR   struct {S SimpleExpr}
	COMPAREEXPR  struct {
		CO Compare_Op
		SE1 SimpleExpr
		SE2 SimpleExpr
	}
)

func (*SIMPLEEXPR) expr_Node() {}
func (*COMPAREEXPR) expr_Node() {}

//-----------------------------------------------------------------
type SimpleExpr interface {
	simpleexpr_Node ()
}

type (
	TERM struct {T Term}
	ADDEXPR struct {O Add_Op ; SE SimpleExpr; T Term }
)

func (*TERM) simpleexpr_Node() {}
func (*ADDEXPR) simpleexpr_Node() {}
//-----------------------------------------------------------------

type Term interface {
	term_Node ()
}

type (
	FACTOR struct {F Factor}
)

func (*FACTOR)   term_Node() {}
//-----------------------------------------------------------------

type Factor interface {
	factor_Node () 
}

type (
	ID   struct {I Ident}
	EXPR struct {E Expr}
	NUMBER struct {N Number}
)

func (*ID)   factor_Node() {}
func (*EXPR) factor_Node() {}
func (*NUMBER) factor_Node() {}
//------------------------------------------------------------








