package main

import ("fmt")

func main(){

	goku :=saiyan{"Goku", 9000}
	super(goku)
	fmt.Println(goku.power)
}
func super(s saiyan){

	s.power +=10000
}
type saiyan struct{
	name  string
	power int
}
