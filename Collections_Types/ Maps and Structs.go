/*
MAPS
example of a map
state Populations := map[string] int{
    "California": 39250017,
    "Texas":      27862596,
    "Florida":    20612439,
    "New York":   19745289,
    "Pennsylvania": 12802503,
    "Illinois":    12801539,
    "Ohio":       11614373,

}
Constraints. The key type (strings) and int have to be consistent throughout the map
Another constraint is, all the key types have to come from a
data type that can be tested for equality like Bool, Arrays string, int, float and etc
Data types that cannot be used as keys (those that cannot be tested for equality) include;
, slices and other functions
Note that an Array is a valid key type but a slice is not


Alternatively, you can use a make() function  to make a map
-this syntax can be used when you are populating values in a loop

statePopulations := make(map[string]int)
state Populations = map[string] int{
    "California": 39250017,
    "Texas":      27862596,
    "Florida":    20612439,
    "New York":   19745289,
    "Pennsylvania": 12802503,
    "Illinois":    12801539,
    "Ohio":       11614373,

}
fmt.Println(statePopulations)

=> Maniputlating the values in the map
 fmt.Println(statePopulations["Ohio"])- this will pull up the valu of ohio entry

 =>Adding values to map
 statePopulations["New State"] = 103103171
 fmt.Println(statePopulations["New State"])

 -- note that the return order of the map is not guaranteed. This is
 not a slice or an array which will provide the same order we provided them

==>Delete Entries in a map
delete(statePopulations, "New State" = 103103171 )

==>interrogating the map
after deleting, if we call the entry, only 0 will return. Which can be very confusing
The same thing also applies if we misspelled the name of the key
-so if you are in a situation where you are not sure whether a key is in the map or not,
you can use (, ok) syntax to check its availability. For instance
population, ok :=statePopulations["Ohio"]
fmt.Println(population, ok)
output-> 1161373 true

but if you want to check for presence, you can just use the righ-only operator
to through in the actual value. example
_, ok := statePopulations["Ohio"]
fmt.Println(ok)
output -> true


==>Cheking the number of elements in the map  len(s)
fmt.Println(len(statePopulations))

==> if you start passing maps around and you start  Maniputlating the data within the map
keep in mind that you can have side effects in any ohter place the map is referred to example
sp := statePopulations
delete(sp, "Ohio")
fmt.Println(sp)
fmt.Println(statePopulations)
Output -> Ohio will be deleted in both sp and statePopulations variables




STRUCT
in struct, you can store any type of data. int, string, structs and many more
example
-they are the most valueble
//
type Doctor struct {
//Think of this section as the field name
    number int
    actorName string
    companions []string
}
func main(){

      aDoctor := Doctor{
                number:3,
                actorName: "Jon Pertwee",
                companions: []string {
                "Liz Shaw",
                "Jo Grant",
              },
    }
    fmt.Println(aDoctor)
}

output: ->Jon Pertwee [Liz Shaw Jo Grant Sarah Jane Smith]

==>interrogating the struct, we use the dot syntax
fmt.Println(aDoctor.actorName) ->joe Grant
fmt.Println(companions[1]) -> Liz Shaw

=>Naming convention
You should use Pascal and Camel Notation but you should not use underscore in
your struct name

==>Anonymous STRUCT
these are structs with very short lifespans
no need to decalare a type. Check the example below;
aDoctor:= struct{name string} {name: "John Pertwee"}
fmt.Println(aDoctor)

==> Structs are of value type. Which means that if you pass
structs around, you are actually passing the copies of the struct
Therefore we use pointers

==> Embedding
GO language doesn't support the natural object oriented principles
Which means it doesn't have inheritence. But it uses a similar concept called
compositions

type Animal struct {
  Name string
  Origin string
}

type Bird struct {
Animal //defining animal properties in Bird, But this means that the animal doesn't have bird properties. it is not a two way relationship
SpeedKPH  float 32
CanFly bool

}
func main (){
b: = Bird{}
b.Name = "Emu"
b.Origin ="Australia"
b.SpeedKPH = 48
b.CanFly = false
fmt.Println(b.Name)
}
-this is definitely is not a two way relationship. But if want to include that concept
we are going to have to use interfaces

TAGS





*/
