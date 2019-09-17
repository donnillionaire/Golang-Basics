// what basically happens is that you package different types together
//type -> a new type called position and struct
/*
-you can call structs inside other structs
- you can parse them to other functions besides main
*/
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
 
func whereIsBadGuy(init badGuy){
  //init is the name you are going to use to refer to what you are parsing
  x := init.pos.x
  y := init.pos.y
  fmt.Println("(", x,",",y,")")
}




func main(){
//  var p position
// Struct initialization Syntax
p := position{4,2}
fmt.Println(p.x) //accessing the fields in the struct

b := badGuy{"Megamind", 100, p}
fmt.Println(b)

whereIsBadGuy(b)
}


/* in Go, if you have created a variable and you haven't set it to anything,
 it will set everything to zero
 */
