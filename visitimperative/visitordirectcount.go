package main
import ("github.com/MartinSteffen/tiny/absynt"
	"github.com/MartinSteffen/tiny/walkimperative"
	"github.com/MartinSteffen/tiny/example"
	"fmt"
)



type visitor struct {count int}   // we cannot make it as pointer to a struct



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
	absynt.WalkStmt(v, s)
	fmt.Println(v)
}
