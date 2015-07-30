package main
import ("fmt")


type record struct{x int}

func inc (s  *record) {
	s.x = 7
}

func main () {
	r:= record{4}
	fmt.Println("1",r)
	r.x = 5
	fmt.Println("2",r)
	inc(&r)
	fmt.Println("2",&r)
	
}







