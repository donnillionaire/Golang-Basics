package main
import "fmt"


func sayHello(name string){
  fmt.Println("Hello", name)
}

func sayGoodbye(name string){
  fmt.Println("Goodbye", name)
}

func beSocial(name string){
  sayHello(name)
  fmt.Println("Hows the weather")
  sayGoodbye(name)
}
// main fuction
func main(){
  beSocial("Don")
  beSocial ("Collins")
}
