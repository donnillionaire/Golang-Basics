package main

import "fmt"

type position struct {
  x float32
  y float32
}
// you can put structs inside of other structs
type badGuy struct{
  name string
  health int
  pos position
}

func whereIsBadGuy(init *badGuy){
  //init is the name you are going to use to refer to what you are parsing

// you don't have to deference these variables, Go does that for you
  x := init.pos.x
  y := init.pos.y
  fmt.Println("(", x,",",y,")")
}

/*func addOne(num *int){
  *num = *num + 1
}*/



func main(){

  x:= 5
  fmt.Println(x)

  xPtr := &x
  fmt.Println(xPtr)

//  addOne(xPtr)
//  fmt.Println(x)
  p := position{4,2}
//  fmt.Println(p.x) //accessing the fields in the struct

  badGuy := badGuy{"Megamind", 100, p}
  //fmt.Println(b)

  whereIsBadGuy(&badGuy)

}
/*Like we have said earlier, when you pass a variable  to a function
 you are simply making a copy of it which can pretty eat alot of memory when
 you are dealing with variables that are not of type int, therefore wasting memory.

 


*/
