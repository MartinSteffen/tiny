package main
import ("github.com/MartinSteffen/tiny/absynt"
	"github.com/MartinSteffen/tiny/walkimperative"
	"github.com/MartinSteffen/tiny/example"
	"fmt"
)



type visitor struct {count int}   // we cannot make it as pointer to a struct




// The following contains all the visit-methods. It is instructive to
// compare their definition with the functional setting.  The type
// ~visitor~ in both cases is the same (the struct). One obvious difference
// is, of course, that in the functional setting the visit functions give
// back an argument of type ~Visitor~ whereas here, there is no return
// value. Less obvious is the type of the receiver. Here it is crucial,
// that it receiver is of type ~*visitor~. Note that with both this type
// and with plain type ~visitor~, the record ~visitor~ is turned into a
// ~Visitor~. Note that in the below definitons, one may even /mix/ the
// receiver types and the compiler won't complain. That means that the
// interface for ``methods'' does not distinguish between call-by-reference
// and call-by-value. Note that one can recursively call in its body for
// instance ~VisitStmt~ as ~v.VisitStmt(s)~ as well as ~(*v).VisitStmt(s)~
// (which of course results in meaningless behavior here). Alternatively,
// if one had the receiver type ~*visitor~, one could make a recursive call
// ~v.VisitStmt(s)~ or ~(&v).VisitStmt(s)~. That's probably because the
// dot-notation is overloaded, and implicitly dereferences member
// access. It's important here that the ~v~ parameter is called by
// reference. Otherwise, the original value does not change.

func (v *visitor) VisitStmt(s absynt.Stmt) {
	fmt.Println("send Stmt")
	v.count = v.count + 1
	fmt.Println("{",v.count,"}")
}


func (v *visitor) VisitExpr(e absynt.Expr) {
	fmt.Println("send Expr")

	v.count = v.count + 1
}


func (v *visitor) VisitSimpleExpr(e absynt.SimpleExpr) {
	fmt.Println("send SimpleExpr")
}


func (v *visitor) VisitIdent(i absynt.Ident)  {
	fmt.Println("I(", i, ")")
	v.count = v.count + 1	
}


func (v *visitor) VisitCompareOp(co absynt.CompareOp) {
	fmt.Println("CompareOp(...)")
	v.count = v.count + 1
}

func (v *visitor) VisitTerm(t absynt.Term)  {
	fmt.Println("Term(...)")
	v.count = v.count + 1	
}

func (v *visitor) VisitFactor(t absynt.Factor)  {
	fmt.Println("Factor(...)")
	v.count = v.count + 1	
}

func (v *visitor) VisitAddOp(ao absynt.AddOp)  {
	fmt.Println("AddOp(...)")
	v.count = v.count + 1	

}

func (v *visitor) VisitMulOp(s absynt.MulOp) {
	fmt.Println("MulOp")
	v.count = v.count + 1	

}



func (v *visitor) VisitNumber(n absynt.Number)  {
	fmt.Println("Number(...)")
	v.count = v.count + 1	

}

func (v *visitor) VisitSymbol(s absynt.Symbol)  {
	fmt.Println("Symbol(...)")
	v.count = v.count + 1	
}


var s = example.S2     // stmt 



// This time we don't make use of concurrency. Therefore we don't need to
// spawn a new process-
func main () {
	v := &visitor{0}   // address of
	walkimperative.WalkStmt(v, s)
	fmt.Println(v)
}
