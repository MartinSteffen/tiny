package main
import ("github.com/MartinSteffen/tiny/absynt"
	"github.com/MartinSteffen/tiny/walkfunctional"
	"fmt"
)


// even if functional, the visitor cannot be just an ~int~. Therefore we
// make a struct again. We don't to change it, so the argument is not
// passed by reference.

type visitor struct {count int}



func (v visitor) VisitStmt(s absynt.Stmt) (walkfunctional.Visitor) {
	fmt.Println("send Stmt") 
	v.count = v.count + 1
	fmt.Println("{",v.count,"}")
	return visitor{v.count +1}
}


func (v visitor) VisitExpr(e absynt.Expr) (w walkfunctional.Visitor) {
	fmt.Println("send Expr")
	v.count = v.count + 1
	return v     // why can I just write "return"?
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


//-----------------------------------------------------------------
// Some concrete tree

var f = &absynt.ID{I:"s"}   // factor
var t = &absynt.FACTOR{f}   // term
var se = &absynt.TERM{t}    // simple expr
var e = &absynt.SIMPLEEXPR{se}  // expr 
var sr = &absynt.READ{I:"x"}    // read stmt
var sl1 = sr             // stmt (list) 
var sl2 = sl1            // stmt (list) 
var s = &absynt.IF{e,sl1,sl2}     // stmt 
var ao =  &absynt.PLUS{}           // add op "+"
var f2 = &absynt.NUMBER{N:42}      // factor
var t2 = &absynt.FACTOR{f2}      // term
var se2 =  &absynt.TERM{t2}       // simple expression
var ao4 = &absynt.MINUS{}         // add op "-"
var f5  = &absynt.ID{"f5"}
var t5 = &absynt.FACTOR{f5}
var se5 = &absynt.TERM{t5}      
var se4 =  &absynt.ADDEXPR{ao4,se5,t2}    // simple expression
var e3  = &absynt.SIMPLEEXPR{se4} // expression
var f3  = &absynt.EXPR{e3}      // factor
var t3 = &absynt.FACTOR{f3}           // term
var se3 = &absynt.ADDEXPR{ao4,se2,t3}  // simple expression
var s2 = &absynt.REPEAT{s,e3}     // stmt 


// This time we don't make use of concurrency. Therefore we don't need to
// spawn a new process.
func main () {
	v := visitor{-1}
	walkfunctional.WalkStmt(v, s2)
	fmt.Println(v)
}
