
package ast2

//import ("fmt")

type Node interface {

}


type Symbol string      
type Ident Symbol
type Number int         
                        
//--------------------------------------------------------
type Compare_Op interface { 
	compare_opNode ()
}
type (                      
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


type Program    [] Stmt   

//--------------------------------------------------------



type  Stmt_Visitor interface {}
type Stmt interface {
	stmt_Node ()
	Accept (Stmt_Visitor)
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







