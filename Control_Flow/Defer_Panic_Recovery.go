=> DEFER 
/*
delays the functions execution until the program finishes, but before it actually returns
*/
 func main (){
 	fmt.Println("start")
 	 defer fmt.Println("start")
 	fmt.Println("start")

 }
 /*
Output --> Start, End, Middle
 */