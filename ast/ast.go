// The abstract syntax is on the one hand inspired by the ML version, but
// also of course the C version, where the C version is rather
// unstructured. The final source of inspiration is the go-ast, to learn
// how they did it.


// The first statement here must be the declaration of the package.  We
// cannot name it package "main", because then it needs a function main, it
// seems.

package ast 

// import ("fmt")

// In the ML code, I use information about the position. That's still more
// or less dummy information. I leave it out here. An elegant solution here
// would be to ``embed'' it. It should be relatively easy.

type symbol string      // might be replaced by something more efficient
type ident symbol
type number int



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


type Program  []  Stmt   // this is a slice type

//--------------------------------------------------------


type Stmt interface {
}

type (
	IF struct {e Expr
		sl1 [] Stmt
		sl2 [] Stmt
	}
	READ struct {i ident}
	WRITE struct {e Expr}
	REPEAT struct {sl [] Stmt; c Expr}
	ASSIGN struct {i ident; e Expr}
)

type Expr interface {}

type (
	SIMPLEEXPR   struct { s SimpleExpr}
	COMPAREEXPR  struct {
		co Compare_Op
		se1 SimpleExpr
		se2 SimpleExpr
	}
)

//-----------------------------------------------------------------

type SimpleExpr interface {}

type (
	TERM struct {t Term}
	ADDEXPR struct {o Add_Op ; se SimpleExpr; t TERM }
)

//-----------------------------------------------------------------

type Term interface {
	termNode()
}


type (
	FACTOR struct {f Factor}
)

func (*FACTOR)   termNode() {}

//-----------------------------------------------------------------
type Factor interface {
	factorNode () 
}

type (
	ID   struct {i ident}
	EXPR struct {e Expr}
	NUMBER struct {n number}
)

func (*ID)   factorNode() {}
func (*EXPR) factorNode() {}
func (*NUMBER) factorNode() {}
//------------------------------------------------------------






func main () {}
