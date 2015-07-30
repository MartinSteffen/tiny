package main
import ("github.com/MartinSteffen/tiny/absynt"
	"github.com/MartinSteffen/tiny/walkimperative"
	"github.com/MartinSteffen/tiny/example"
	"fmt"
)

// Note that the visitor type is not the same of the same as Visitor.  It
// will be ``turned'' into a ``Visitor'' only when attaching the
// appropriate Visit-method to it. The code is intended for some initial
// tests only. Let's take as accumulate measure the sum of notes.  We just
// also make it to a channel, but perhaps a better way would be to all sum
// it up on the fly, not just send 1s over a channel. We try that later.

type visitor chan int 

// What is unclear is the role of w: the w is not used (for me it could be
// left out), but it's unclear anyhow: one could return w as well, or call
// in the body something like w.Visit (n), but it's unclear what this is
// supposed to mean. 

// func (v visitor) Visit(n absynt.Program) (w absynt.Visitor) {
//	 w.Visit (n)  // what would that mean
//	 return 
// }

func (v visitor) VisitStmt(s absynt.Stmt) {
	fmt.Println("send Stmt") 
	v <- 2
}


func (v visitor) VisitExpr(e absynt.Expr) {
	fmt.Println("send Expr") 
	v <- 1
}


func (v visitor) VisitSimpleExpr(e absynt.SimpleExpr)  {
	fmt.Println("send SimpleExpr") 
	v <- 9
}


func (v visitor) VisitIdent(i absynt.Ident) {
	fmt.Println("I(", i, ")") 
	v <- 3
}


func (v visitor) VisitCompareOp(co absynt.CompareOp)  {
	fmt.Println("CompareOp(...)")
	v <- 4
}

func (v visitor) VisitTerm(t absynt.Term) {
	fmt.Println("Term(...)")
	v <- 5
}

func (v visitor) VisitFactor(t absynt.Factor) {
	fmt.Println("Factor(...)")
	v <- 6
}

func (v visitor) VisitAddOp(ao absynt.AddOp) {
	fmt.Println("AddOp(...)")
	v <- 7

}

func (v visitor) VisitMulOp(s absynt.MulOp)  {
	fmt.Println("MulOp")
	v <- 11

}



func (v visitor) VisitNumber(n absynt.Number)  {
	fmt.Println("Number(...)")
	v <- 8

}

func (v visitor) VisitSymbol(s absynt.Symbol) {
	fmt.Println("Symbol(...)")
	v <- 10

}

var s = example.S2     // stmt


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
		walkimperative.WalkStmt(v, s)
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
