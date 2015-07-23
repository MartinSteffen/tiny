
package absynt


type visitor chan int 


// ----------------------------------
// the following is taken from the go testprinter etc.
// We take it as first inspiration
// 
//
// type visitor chan *ast.Ident

// func (v visitor) Visit(n ast.Node) (w ast.Visitor) {
// 	if ident, ok := n.(*ast.Ident); ok {
// 		v <- ident
// 	}
// 	return v
// }

// // idents is an iterator that returns all idents in f via the result channel.
// func idents(f *ast.File) <-chan *ast.Ident {
// 	v := make(visitor)
// 	go func() {
// 		ast.Walk(v, f)
// 		close(v)
// 	}()
// 	return v
// }

// // identCount returns the number of identifiers found in f.
// func identCount(f *ast.File) int {
// 	n := 0
// 	for range idents(f) {
// 		n++
// 	}
// 	return n
// }


// import ("fmt")

// This stuff here is old code, so it probably should be removed. I have
// commented it out to avoid double declarations

// //----------------------------------------------------

// type  Stmt_Visitor interface {
// 	visit_IF     (is *IF) 
// 	visit_READ   (rs *READ)
// 	visit_REPEAT (rs *REPEAT)
// 	visit_ASSIGN (ss *REPEAT)
// }

// //--------------------------------------------------------


// ///////////////////////////////////////////////////////////////
// // Client code ///



// type Visitor struct{
//     Messages []string
// }

// func (this *Visitor) visit_Expr (e Expr) {
// 	switch e.(type) {
// 	case *SIMPLEEXPR:
// 		fmt.Println("")
// 	}
	
	
// } 



// func (this *Visitor) visit_Stmt (s Stmt) {
// 	switch x := s.(type){
// 	case *IF:
// 		fmt.Println("")
// 		this.visit_Expr (x.E)
// 		this.visit_Stmt (x.SL1)
// 		this.visit_Stmt (x.SL2)
// 	}
// }

 
// func (this *Visitor) visit_IF(i *IF) {
// 	this.Messages = append(this.Messages,   // collect the stuff
// 		fmt.Sprintf("Visiting IF\n"))
// 	fmt.Println(i.E)
//  	e  := i.E
//  	s1 := i.SL1
//  	s2 := i.SL2
//  	this.visit_Expr (e)  // dispatch
//  	this.visit_Stmt (s1)  // dispatch
//  	this.visit_Stmt (s2)  // dispatch
	
// }



