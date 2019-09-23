/*NOTES
-> with the design of GO language,  arrays are laid out contigous in memory which means
it is very easy to work with them and very fast accessing the contents of the arrays

 -> Arrays can  only hold one data type at a time
 var students [...] - telling an array to hold as many as they will be specified

 -> if you are wondering why the array index starts from zero is because, GO will have a pointers
 at the beginning of that array. And the index that we pass, in this case 0 is going to tell it
 how many strings to walk forward. So it knows that when it has a string and it has a certain lenght,
 so it knows who many steps it's going to take forward

-DETERMINING THE SIZE OF THE ARRAY : len(students)

-DECLARING ARRAYS IN ARRAYS (REPRESENTING A 3by3 MATRIX)
var identityMatrix [3][3] int
identityMatrix[0] = [3] int {1,0,0}
identityMatrix[1] = [3] int {0,1,0}
identityMatrix[2] = [3] int {,0,1}

 output -> [1 0 0] [0 1 0] [1 0 1]

-in GO arrays are actually considered as values. in alot of languages when you create an
array, it is actually pointing to the values in that array, so if you pass things arounc,
you are actually passing the same data. But in Go, when you pass an array, you are
actually creating a literal copy so it is not pointing to the same underline data
it's pointing to a different set of data. Which means that it has to reassign the entire length
of that array. Check the example below

a:= [...] int {1,2,3}
b:= a
b[1] =5
fmt.Println(a)
fmt.Println(b)

output -> [1 2 3]
          [1 5 3]


 To avoid this problem, we use pointers. Check the example below;
 a:= [...] int {1,2,3}
 b:= &a
 b[1] =5
 fmt.Println(a)
 fmt.Println(b)

 output -> [1 5 3]
          &[1 5 3]

=> SLICES
Arrays are indeed powerful so, but the fact that you have to know the size of the arrar
at compile time, limits their usage. Because of this, SLICES come in

Everything you can do with arrays, you can do with slices and more.
For example, besides the length function, there is a capacity function as well. This is
because, the number of elements in a slice doesn't necessarily match the size
of the backing array because a slice is just a projection of that underlying arrays

- unlike arrays,slices are reference types which means they point to the same
undelying data

a:= [] int {1,2,3}
fmt.Println(a)
output -> [1 2 3]

=> WAYS OF CREATING SLICES
a:= []int{1,2,3,4,5,6,7,8,9,10}
b:= a[:] //slice of all the elements of what it is referring to which are all the elements of a
c:= a[3:]// Starts with the 3rd element of the parent and copy all the values after that
d:= a[:6]// this will copy everything up to the six Syntax
e:=a[3:6] //slice the 4th, 5th and 6th element
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)

output
[1 2 3 4 5 6 7 8 9 10]
[1 2 3 4 5 6 7 8 9 10]
[4 5 6 7 8 9 10]
[1 2 3 4 5 6 ]
[4 5 6]

=> unlike arrays,slices are reference types which means they point to the same
undelying data. Check the example below;
a:= []int{1,2,3,4,5,6,7,8,9,10}
b:= a[:] //slice of all the elements of what it is referring to which are all the elements of a
c:= a[3:]// Starts with the 3rd element of the parent and copy all the values after that
d:= a[:6]// this will copy everything up to the six Syntax
e:=a[3:6] //slice the 4th, 5th and 6th element

//changing the value of index 5
a[5] = 42

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)


output
[1 2 3 4 5 42 6 7 8 9 10]
[1 2 3 4 5 42 7 8 9 10]
[4 5 42 7 8 9 10]
[1 2 3 4 5 42 ]
[4 5 42]

=> Slices can also work with ARRAYS
And if you do so, all the properties of slices will be adopted in the array
a:= [...]int{1,2,3,4,5,6,7,8,9,10}-> declaring a slice as an array
b:= a[:] //slice of all the elements of what it is referring to which are all the elements of a
c:= a[3:]// Starts with the 3rd element of the parent and copy all the values after that
d:= a[:6]// this will copy everything up to the six Syntax
e:=a[3:6] //slice the 4th, 5th and 6th element

//changing the value of index 5
a[5] = 42

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
fmt.Println(e)


output
[1 2 3 4 5 42 6 7 8 9 10]
[1 2 3 4 5 42 7 8 9 10]
[4 5 42 7 8 9 10]
[1 2 3 4 5 42 ]
[4 5 42]

=> ALTERNATIVELY YOU CAN USE make() function to create a slice and it takes in two
or three arguments,
 one is the type of object we are going to create, then  the length of the slices and capacity as the third argument

 a:= make([]int, 3)
 fmt.Println(a)
 fmt.Println(len(a))
 fmt. Println(capa(a))

=>ADDING ELEMENTS TO THE SLICE using :-> append(array, newelement, new element)
a:= make([]int, 3)
fmt.Println(a)
fmt.Println(len(a))
fmt. Println(cap(a))
 a = append(a, 1)

CONCATNATION OF SLICES
 a = append(a, []int{2,3,4,5}...)


STACK OPERATIONS WITH SLICES
for example, lets say we are treating our elements as a stack. So we want
to be able to add some element or push them off the stack

append() will allow us to push elements on the stack
but we can use:
-> shift operations which means we want to be able to remove the first element from the stack
a:=[]int{1,2,3,4,5}
//b:=a[1:] =>Shift operation-> this will create a new index value that starts at index 1
b:a[:len(a)-1]- removing an element at the end of the slice
it is pretty easy to remove an element from the beginning of the slice or at the end of the slice
But what if we want to remove elements from the middle? here is how we do its
- we are going to have to concatenate two slices together. The first slice is going
to be up to where we want to remove the element and then the other slice is from the
next element after the one want  to remove

b: = append(a[:2], a[3:]...)


SUMMARY
Arrays
--> Collection of items with the same types
--> Fixed Size
--> Declaration Styles
   ->a := [3]int {1,2,3}
   ->b :=[...] int{1,2,3}
   ->var a [3]identityMatrix
-->Access via zero-based index
--> len function returns the size of the array
--> copies refer to different underlying data


SLICES
--> Backed by arrays
-->Creation Styles
  ->Slice exisiting array or slice
  ->Literal Style
  ->using make function
    a:= make([]int, 10)// create a slice with a capacity and length of 10
    a:= make([]int, 10, 100)//create a slice with a  length of 10 and capacity of 100
--> len function returns length of the slices
-->cap function returns the lenght of the underlying array
--> append function adds elements to the slice
    ->may cause expensive copy operations if the underlying array is too small
-->> Copies refer to same underlying array 










*/
