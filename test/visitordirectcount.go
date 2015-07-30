package main
import ("github.com/MartinSteffen/tiny/absynt"
	"fmt"
)


type visitor struct {count int}   // we cannot make it as pointer to a struct



func (v *visitor) VisitStmt(s absynt.Stmt) {
	fmt.Println("send Stmt") 
	v.count = v.count + 1
	fmt.Println("{",v.count,"}")
}


func (v visitor) VisitExpr(e absynt.Expr) {
	fmt.Println("send Expr")
	v.count = v.count + 1
}


func (v visitor) VisitSimpleExpr(e absynt.SimpleExpr) {
	fmt.Println("send SimpleExpr")
}


func (v visitor) VisitIdent(i absynt.Ident)  {
	fmt.Println("I(", i, ")")
	v.count = v.count + 1	
}


func (v visitor) VisitCompareOp(co absynt.CompareOp) {
	fmt.Println("CompareOp(...)")
	v.count = v.count + 1
}

func (v visitor) VisitTerm(t absynt.Term)  {
	fmt.Println("Term(...)")
	v.count = v.count + 1	
}

func (v visitor) VisitFactor(t absynt.Factor)  {
	fmt.Println("Factor(...)")
	v.count = v.count + 1	
}

func (v visitor) VisitAddOp(ao absynt.AddOp)  {
	fmt.Println("AddOp(...)")
	v.count = v.count + 1	

}

func (v visitor) VisitMulOp(s absynt.MulOp) {
	fmt.Println("MulOp")
	v.count = v.count + 1	

}



func (v visitor) VisitNumber(n absynt.Number)  {
	fmt.Println("Number(...)")
	v.count = v.count + 1	

}

func (v visitor) VisitSymbol(s absynt.Symbol)  {
	fmt.Println("Symbol(...)")
	v.count = v.count + 1	
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

//repeat (if "s" (read "x") (read "x"))
// until  ("f5" - 42)

// -----------------------------------------------------

// The visitor is a channel. To use it we need some concurrency, i.e., the
// visiting function needs to run in parallel to send into the channel, and
// another thread to read from it. Let's make the walking function the one
// being spawned. The visitor sends on the synchronous channel, and the
// main function receives the values. This is done via the range-command
// (where currently the value is actually not remembered, since the main
// function simply counts.  



//func idents() <-chan int {
//	v := visitor{}      // create a struct
//	go func() {
//		
//	}()
//	return v
//}

// This time we don't make use of concurrency. Therefore we don't need to
// spawn a new process-
func main () {
	v := &visitor{-1}   // address of
	absynt.WalkStmt(v, s2)
	fmt.Println(v)
}
