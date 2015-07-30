package main
import ("github.com/MartinSteffen/tiny/absynt"
	"fmt"
)


type visitor chan int 


func (v visitor) VisitStmt(s absynt.Stmt) (w absynt.Visitor) {
	fmt.Println("send Stmt") 
	v <- 2
	return v
}


func (v visitor) VisitExpr(e absynt.Expr) (w absynt.Visitor) {
	fmt.Println("send Expr") 
	v <- 1
	return v     // why can I just write "return"?
}


func (v visitor) VisitSimpleExpr(e absynt.SimpleExpr) (w absynt.Visitor) {
	fmt.Println("send SimpleExpr") 
	v <- 9
	return v     // why can I just write "return"?
}


func (v visitor) VisitIdent(i absynt.Ident) (w absynt.Visitor) {
	fmt.Println("I(", i, ")") 
	v <- 3
	return v     // why can I just write "return"?
}


func (v visitor) VisitCompareOp(co absynt.CompareOp) (w absynt.Visitor) {
	fmt.Println("CompareOp(...)")
	v <- 4
	return v     
}

func (v visitor) VisitTerm(t absynt.Term) (w absynt.Visitor) {
	fmt.Println("Term(...)")
	v <- 5
	return v
}

func (v visitor) VisitFactor(t absynt.Factor) (w absynt.Visitor) {
	fmt.Println("Factor(...)")
	v <- 6
	return v
}

func (v visitor) VisitAddOp(ao absynt.AddOp) (w absynt.Visitor) {
	fmt.Println("AddOp(...)")
	v <- 7
	return v

}

func (v visitor) VisitMulOp(s absynt.MulOp) (w absynt.Visitor) {
	fmt.Println("MulOp")
	v <- 11
	return v

}



func (v visitor) VisitNumber(n absynt.Number) (w absynt.Visitor) {
	fmt.Println("Number(...)")
	v <- 8
	return v

}

func (v visitor) VisitSymbol(s absynt.Symbol) (w absynt.Visitor) {
	fmt.Println("Symbol(...)")
	v <- 10
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



func idents() <-chan int {
	v := make(visitor)
	go func() {
		absynt.WalkStmt(v, s2)
		close(v)
	}()
	return v
}


func main () {
 	fmt.Println("here")
 	n := 0 
 	for range idents() {  // read from the channel (but forget the value), stop when closed
 		fmt.Println("loop")
 		n++
		
 	}
 	fmt.Println(n)
}
