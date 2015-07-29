package main
import ("github.com/MartinSteffen/tiny/absynt"
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

func (v visitor) VisitStmt(s absynt.Stmt) (w absynt.Visitor) {
	fmt.Println("send 2")  // never printed
	v <- 2
	return v
}


func (v visitor) VisitExpr(e absynt.Expr) (w absynt.Visitor) {
	fmt.Println("send 1")  // never printed
	v <- 1
	return v     // why can I just write "return"?
}

//-----------------------------------------------------------------
// 

var f = &absynt.ID{I:"s"}   // factor
var t = &absynt.FACTOR{f}   // term
var se = &absynt.TERM{t}    // simple expr
var e = &absynt.SIMPLEEXPR{se}  // expr 
var sr = &absynt.READ{I:"x"}    // read stmt
var sl1 = sr             // stmt (list) 
var sl2 = sl1            // stmt (list) 
var s = &absynt.IF{e,sl1,sl2}   // stmt 




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
		absynt.WalkStmt(v, s)
		close(v)
	}()
	return v
}


func main () {
	fmt.Println("here")
	n := 0 
	for range idents() {  // read from the channel (but forget the value), stop when closed
		fmt.Println("====")
		n++

	}
	fmt.Println(n)
}



// func main () {
// 	v := make (visitor)   // creating a visitor (which is a synchronous chan)
// 	go func () {
// 		absynt.WalkStmt (v,s)
// 		close(v)
// 	}()
// 	fmt.Println("here")
// 	n := 0 
// 	for range v {  // read from the channel (but forget the value), stop when closed
// 		n++
// 		fmt.Println("====")
// 	}
// 	fmt.Println(n)
// 	fmt.Println(v)
// }



