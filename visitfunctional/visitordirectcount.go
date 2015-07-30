package main
import ("github.com/MartinSteffen/tiny/absynt"
	"github.com/MartinSteffen/tiny/example"
	"github.com/MartinSteffen/tiny/walkfunctional"
	"fmt"
)


// even if functional, the visitor cannot be just an ~int~. Therefore we
// make a struct again. We don't to change it, so the argument is not
// passed by reference.

type visitor struct {count int}



func (v visitor) VisitStmt(s absynt.Stmt) (w walkfunctional.Visitor) {
	fmt.Println("send Stmt") 
	v.count = v.count + 1
	fmt.Println("{",v.count,"}")  // w.cound does not work (lack of public field
	w = visitor{v.count+1+1}
	return v
}


func (v visitor) VisitExpr(e absynt.Expr) (w walkfunctional.Visitor) {
	fmt.Println("send Expr")
	w  = visitor{v.count + 1}
	return w     // why can I just write "return"?
}


func (v visitor) VisitSimpleExpr(e absynt.SimpleExpr) (w walkfunctional.Visitor) {
	fmt.Println("send SimpleExpr")
	return v     // why can I just write "return"?
}


func (v visitor) VisitIdent(i absynt.Ident) (w walkfunctional.Visitor) {
	fmt.Println("I(", i, ")")
	v.count = v.count + 1	
	return v     // why can I just write "return"?
}


func (v visitor) VisitCompareOp(co absynt.CompareOp) (w walkfunctional.Visitor) {
	fmt.Println("CompareOp(...)")
	v.count = v.count + 1
	return v     
}

func (v visitor) VisitTerm(t absynt.Term) (w walkfunctional.Visitor) {
	fmt.Println("Term(...)")
	v.count = v.count + 1	
	return v
}

func (v visitor) VisitFactor(t absynt.Factor) (w walkfunctional.Visitor) {
	fmt.Println("Factor(...)")
	v.count = v.count + 1	
	return v
}

func (v visitor) VisitAddOp(ao absynt.AddOp) (w walkfunctional.Visitor) {
	fmt.Println("AddOp(...)")
	v.count = v.count + 1	
	return v

}

func (v visitor) VisitMulOp(s absynt.MulOp) (w walkfunctional.Visitor) {
	fmt.Println("MulOp")
	v.count = v.count + 1	
	return v

}



func (v visitor) VisitNumber(n absynt.Number) (w walkfunctional.Visitor) {
	fmt.Println("Number(...)")
	v.count = v.count + 1	
	return v

}

func (v visitor) VisitSymbol(s absynt.Symbol) (w walkfunctional.Visitor) {
	fmt.Println("Symbol(...)")
	v.count = v.count + 1	
	return v
}


//---------------------------------
var s = example.S2     // stmt 

// This time we don't make use of concurrency. Therefore we don't need to
// spawn a new process.
func main () {
	v := visitor{0}
	walkfunctional.WalkStmt(v, s)
	fmt.Println(v)
}
