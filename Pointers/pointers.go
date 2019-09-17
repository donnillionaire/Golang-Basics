package main

import "fmt"

func addOne(num *int){
  *num = *num + 1 // dereferencing the pointers
}



func main(){

  x:= 5
  fmt.Println(x)

  xPtr := &x
  fmt.Println(xPtr)

  addOne(xPtr)
  fmt.Println(x)




}
/*
xPtr is the address where 5 is stored into, so we pass in the address
to addOne function which adds 1 to the value stored in that address then
later finally, we out put the value of x

Normally when we are passing variables in functions, we are simply creating a copy of it

in conclusion, pointers can be useful when you want to change the function to actually change the variable
you are passing into it rather than just getting a copy and changing it

*/
