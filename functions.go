package main
import "fmt"


func sayHello(name string){
  fmt.Println("Hello", name)
}

func sayGoodbye(name string){
  fmt.Println("Goodbye", name)
}

// Return Fuctions
func addOne(x int) int {
  return x + 1
}

func beSocial(name string){
  sayHello(name)
  fmt.Println("Hows the weather")
  sayGoodbye(name)
}

//Recursion:-Functions Can call themselves (making infinite loop)
func sayAbunchofHellos(){
  fmt.Println("hello")
  sayAbunchofHellos()

}


// main fuction
func main(){
  beSocial("Don")
  beSocial ("Collins")

  //calling the return function (addOne)
  x:= 5
  x = addOne(x)
  fmt.Println(x)

  //Composing Fuctions
  x = addOne(addOne(x))
    fmt.Println(x)





}
